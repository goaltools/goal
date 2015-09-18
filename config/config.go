// Package config is a wrapper around Thomasdezeeuw/ini
// parser.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Thomasdezeeuw/ini"
)

const (
	systemSection = ini.Global

	// ReadFromFile passed to ParseFile as a second argument means
	// active section name may be found inside ini config's
	// system section.
	ReadFromFile = systemSection

	// Below are reserved key names.
	keyActiveSection = "active.section"
	keyExtend        = "extend"
)

// Variables in a ${NAME} form inside configuration file are
// expected to be treated as ENV vars.
var envVar = regexp.MustCompile(`\${([A-Za-z0-9._\-]+)}`)

// Getter is used to get values from configuration file.
type Getter interface {
	// StringDefault returns a value associated with the key.
	// If such key does not exist, default value is returned.
	StringDefault(key, defaultValue string) string

	// Section returns keys and values of the specified section
	// as a map[string]string.
	Section(name string) map[string]string
}

// Config implements Getter interface. It provides methods
// for work with ini style configuration files.
type Config struct {
	// activeSection will be used as a default one
	// when getting values from config.
	activeSection string

	// data is a content extracted from config file.
	// It contains section names -> keys -> values.
	data map[string]map[string]string

	// paths is a list of all files included by this config.
	// It is used to make sure there are no cycle dependencies.
	paths map[string]bool
}

// New allocates and returns a new Config.
func New() *Config {
	return &Config{}
}

// ParseFile opens and parses the specified ini file
// and writes the result to the Config if there are no errors.
// activeSession will be used by methods such as StringDefault
// when using extracted data.
func (c *Config) ParseFile(path, activeSection string) error {
	// Openning the configuration file.
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Parsing the openned file.
	c.data, err = ini.Parse(f)
	if err != nil {
		return err
	}

	// Trying to extend this config from a parent,
	// if this is required.
	if c.paths == nil {
		c.paths = map[string]bool{}
	} else if c.paths[path] {
		return fmt.Errorf(`cannot parse "%s", cicle inheritance of config files is not allowed`, path)
	}
	c.paths[path] = true
	if p, ok := c.getString(systemSection, keyExtend); ok {
		// Parse the parent file.
		nc := New()
		nc.paths = c.paths
		err := nc.ParseFile(relFilepath(path, p), "")
		if err != nil {
			return err
		}

		// Joining current config with the parent.
		for s := range nc.data {
			for k := range nc.data[s] {
				// Check whether child has already has such a section - key pair.
				if _, ok := c.getString(s, k); ok {
					continue
				}

				// Otherwise, copy the associated value from parent.
				c.setString(s, k, nc.data[s][k])
			}
		}
	}

	// Selecting an active section.
	switch activeSection {
	case ReadFromFile:
		c.activeSection, _ = c.String(activeSection, keyActiveSection)
	default:
		c.activeSection = activeSection
	}
	return nil
}

// relFilepath gets a path of base ini config and its parent's
// path and returns the latter as a relative one, e.g.:
//	baseConf = ./config/app.ini
//	parent = test.ini
// The function will return:
//	config/test.ini
func relFilepath(baseConf, parent string) string {
	if filepath.IsAbs(parent) {
		return parent
	}
	return filepath.Join(filepath.Dir(baseConf), parent)
}

// String returns a value associated with the key from the section.
// If no such kind of key exists in the specified section, a value
// from default one is returned.
// If there is no such key there, too empty string and false as a second
// argument are returned.
func (c *Config) String(section, key string) (string, bool) {
	if v, ok := c.getString(section, key); ok {
		return v, ok
	}

	// If not, make sure section is not a global one.
	// If so, try to extract value from the global section.
	if g := ini.Global; g != section {
		return c.String(g, key)
	}

	// If still no result, return empty string and false.
	return "", false
}

// getString returns a value associated with the specified key
// from the requested section or false if no such key is found.
func (c *Config) getString(section, key string) (string, bool) {
	// Make sure such section exists.
	if _, ok := c.data[section]; !ok {
		return "", false
	}

	// Check whether such key exists withing the section.
	val, ok := c.data[section][key]
	if !ok {
		return "", false
	}

	// Try to replace environment variables.
	val = envVar.ReplaceAllStringFunc(val, func(k string) string {
		return os.Getenv(envVar.ReplaceAllString(k, "$1"))
	})

	// Return the value.
	return val, true
}

// setString sets a value of the specified key from a section.
// It allocates a new section map if necessary.
func (c *Config) setString(section, key, value string) {
	if _, ok := c.data[section]; !ok {
		c.data[section] = map[string]string{}
	}
	c.data[section][key] = value
}

// StringDefault gets a value associated with the requested key
// from an active section and returns it.
// If no value is found, default one is returned.
func (c *Config) StringDefault(key, defaultValue string) string {
	if v, ok := c.String(c.activeSection, key); ok {
		return v
	}
	return defaultValue
}

// Section returns the whole section of a configuration file.
func (c *Config) Section(name string) map[string]string {
	return c.data[name]
}
