// Package ini provides a type that implements Interface of the
// "github.com/conveyer/config" for the INI configuration format.
package ini

import (
	"strings"

	"github.com/conveyer/config"

	"github.com/conveyer/ini"
)

const separator = "."

// INI is an implementation of config.Interface for ini
// configuration files.
type INI struct {
	data       map[string]map[string]interface{}
	prefixPath string
}

// New allocates and returns a new INI type.
func New(data map[string]map[string]interface{}) *INI {
	return &INI{data: data}
}

// New allocates a new configuration by parsing the
// requested file and returns it.
// Elements of prefixPath will be joined as follows:
//	prefixPath[0] + "." + ... + "." + prefixPath[N]
func (c *INI) New(file string, prefixPath ...string) (config.Interface, error) {
	m, err := ini.OpenFile(file)
	if err != nil {
		return nil, err
	}
	t := New(m)
	t.prefixPath = strings.Join(prefixPath, separator)
	return t, nil
}

// Join merges a requested file with the current configuration file.
// Values of a new file must have priority over the values of the
// current configuration. E.g. if the config we have looks as follows:
//	obj:
//		key1 = value1
//		key2 = value2
// and the new input configuration is:
//	obj:
//		key2 = another_value
//		key3 = value3
// The original config must be turned into:
//	obj:
//		key1 = value1
//		key2 = another_value
//		key3 = value3
func (c *INI) Join(file string) error {
	// Make sure current configuration is initialized.
	if c.data == nil {
		c.data = map[string]map[string]interface{}{}
	}

	// Open the requested configuration file and parse it.
	m, err := ini.OpenFile(file)
	if err != nil {
		return err
	}

	// Iterate over all available sections of the input config.
	for section := range m {
		// Make sure such section exists in the current config's map.
		if _, ok := c.data[section]; !ok {
			c.data[section] = map[string]interface{}{}
		}

		// Iterate over all available keys of the section and join them.
		for key := range m[section] {
			c.data[section][key] = m[section][key]
		}
	}
	return nil
}

// Value is an equivalent of ValuePrefixless but returns a value associated
// with the pathPrefix specified in constructor + input path. As an example:
//	c, _ := (&Config{}).New("/path/to/file.ini", "my", "default", "section")
//	d := c.Value("some", "key")
// The code above will try to return a value associated with:
//	[my.default.section]
//	some.key = value that will be returned
func (c *INI) Value(path ...string) config.ValueInterface {
	return c.ValuePrefixless(append([]string{c.prefixPath}, path...)...)
}

// ValuePrefixless returns a value associated with the requested path.
// The first element of the path is used as a section name, all others
// are joined as path[1] + "." + ... + "." + path[N] to be used as a key.
// If only one parameter is provided, default section name is used.
// As an example:
//	c, _ := (&INI{}).New("path/to/file.ini")
//	d := c.ValuePrefixless("mySection", "user", "name", "length")
// The code above will search for the value associated with:
//	[mySection]
//	user.name.length = 179
func (c *INI) ValuePrefixless(path ...string) config.ValueInterface {
	// Prepare a section name:
	// If there are more than 1 element in the path, assume
	// that the first one is the section name.
	// Otherwise, use a default section name.
	sect := ""
	switch l := len(path); true {
	case l > 1:
		sect = path[0]
		path = path[1:]
	default:
		sect = c.ValuePrefixless("@config", "default", "section", "name").StringDefault("")
	}

	// Check whether requested section does exist.
	if _, ok := c.data[sect]; !ok {
		return config.NewValue(nil)
	}

	// Prepare key name and make sure it is presented
	// in the requested section.
	k := strings.Join(path, ".")
	if v, ok := c.data[sect][k]; ok {
		return config.NewValue(v)
	}
	return config.NewValue(nil)
}
