// Package ini provides functions for parsing INI configuration
// files with extended syntax. E.g. arrays and references are supported.
// For more details, see the README.md.
package ini

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/conveyer/ini/parser"
)

const (
	refKey   = "$"
	refPref  = "&"
	arrayLit = "[]"
)

// config represents a parsed and processed configuration
// file. It has the following structure:
//	section_name:
//		key:   string_value
//		key[]: []string_values
type config map[string]map[string]interface{}

// context implements methods for processing
// of INI sections object and its transformation
// into a configuration map.
type context struct {
	obj, refs config
}

// allocate makes sure a map with the requested key in the config
// is allocated.
func (c config) allocate(n string) {
	if _, ok := c[n]; ok {
		return
	}
	c[n] = map[string]interface{}{}
}

// OpenFile gets a path to INI file, opens, parses, and returns it.
// A non-nil error is returned as a second argument in
// case the requested file cannot be parsed.
func OpenFile(path string) (map[string]map[string]interface{}, error) {
	// Try to open the requested file.
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Scan and parse it.
	sections, err := parser.Parse(bufio.NewScanner(f))
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %s", err)
	}

	// Transform into the final object and return
	// if there are no errors.
	c := &context{}
	if err = c.process(sections); err != nil {
		return nil, fmt.Errorf("failed to process: %s", err)
	}
	return c.obj, nil
}

// process gets a number of INI sections returned by
// a parser and transforms them into a configuration.
func (c *context) process(ss []parser.Section) error {
	// Process reference sections.
	err := c.processRefs(ss)
	if err != nil {
		return err
	}

	// Process other sections.
	return c.processSections(ss)
}

// processRefs processes special reference sections.
// It makes sure that there are no links to other sections
// inside them as only regular sections can use
// "$ = &section_name" syntax.
func (c *context) processRefs(ss []parser.Section) error {
	// Allocate the reference config object.
	c.refs = config{}

	// Iterate over all available sections to find
	// the reference ones.
	for i := range ss {
		// Ignore non-reference sections.
		n := string(ss[i].Name)
		if !strings.HasPrefix(n, refPref) {
			continue
		}

		// As soon as a reference section has been found,
		// add its key-value pairs to the config.
		// Make sure there are no link keys inside ("false" argument).
		c.refs.allocate(n)
		err := c.appendKVs(c.refs[n], ss[i].Keys, ss[i].Values, false)
		if err != nil {
			return fmt.Errorf(
				`reference section "%s": no references allowed, %s`, n, err,
			)
		}
	}
	return nil
}

// processSections processes regular sections and replaces
// "$ = &section_name" key-value pairs by the key-values of the
// respective sections. E.g. there is a configuration:
//	section1:
//		key1 = value1
//		& = section2
//	section2:
//		key2 = value2
// It should be transformed into:
//	section1:
//		key1 = value1
//		key2 = value2
//	section2:
//		key2 = value2
func (c *context) processSections(ss []parser.Section) error {
	// Allocate the config object.
	c.obj = config{}

	// Iterate over all available sections to find
	// the regular ones.
	for i := range ss {
		// Ignore reference sections.
		n := c.processSectionName(ss[i].Name)
		if strings.HasPrefix(n, refPref) {
			continue
		}

		// As soon as a regular section has been found,
		// add its values to the config.
		c.obj.allocate(n)
		err := c.appendKVs(c.obj[n], ss[i].Keys, ss[i].Values, true)
		if err != nil {
			return fmt.Errorf(
				`section "%s": %s`, n, err,
			)
		}
	}
	return nil
}

// appendKVs gets a map and pairs of keys & values.
// It inserts the key-value pairs into the map.
func (c *context) appendKVs(m map[string]interface{}, ks, vs [][]byte, allowRefs bool) error {
	for i := range ks {
		// Process all of the possible errors associated with the references.
		k := string(ks[i])
		v := replaceEnvVars(string(vs[i])) // Replace ${NAME} by respective environment variables.
		ok, err := c.processRef(k, v, allowRefs)
		if err != nil {
			return err
		}
		if ok {
			// Current key-value pair is a reference and there are no
			// any errors so far, so join the maps.
			c.join(m, c.refs[v])
			continue
		}

		// If no array literals are presented, just add
		// the key-value pair to the map.
		if !strings.HasSuffix(k, arrayLit) {
			m[k] = v
			continue
		}
		// Otherwise, check whether the array has already been
		// declared earlier. If it isn't, do it now by adding
		// the first element.
		k = strings.TrimSuffix(k, arrayLit) // Array literal is not a part of key's name.
		if _, ok := m[k]; !ok {
			m[k] = []string{v}
			continue
		}

		// If the array element with current key has already
		// exist, append the value.
		m[k] = append(m[k].([]string), v)
	}
	return nil
}

// processRef checks correctness of a reference.
func (c *context) processRef(k, v string, allowRefs bool) (bool, error) {
	// If this is not a reference, do nothing.
	if k != refKey {
		return false, nil
	}

	// Otherwise, make sure references are allowed.
	if !allowRefs {
		return true, fmt.Errorf(`"%s = %s" was not expected here`, k, v)
	}

	// Make sure a referenced section name starts with a "&".
	if !strings.HasPrefix(v, refPref) {
		return true, fmt.Errorf(`"%s = %s": a reference section was expected instead of "%s"`, k, v, v)
	}

	// Make sure a referenced section does exist.
	if _, ok := c.refs[v]; !ok {
		return true, fmt.Errorf(`"%s = %s": reference section "%s" does not exist`, k, v, v)
	}
	return true, nil
}

// join adds values of the child map to the parent one.
// Slice objects are appended instead of being overridden. E.g.:
//	[section1]:
//		arr[] = a
//		$ = &smth
//	[&smth]
//		arr[] = b
//		arr[] = c
// In the configuration above arr[] is equal to [a, b, c]
func (c *context) join(parent, child map[string]interface{}) {
	for k, v := range child {
		switch v.(type) {
		case []string:
			parent[k] = append(parent[k].([]string), v.([]string)...)
		default:
			parent[k] = v
		}
	}
}

// processSectionName takes care unification of different variations
// of default section.
func (c *context) processSectionName(n []byte) string {
	s := string(n)
	if strings.ToLower(s) == "default" {
		return ""
	}
	return s
}
