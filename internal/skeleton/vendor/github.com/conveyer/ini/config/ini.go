package config

// INI is an implementation if Interface for ini
// configuration files.
type INI struct {
	data           map[string]map[string]interface{}
	defaultSection string
}

// String returns a string value associated with the requested key.
func (c *INI) String(key string) (string, bool) {
	return c.StringFromSection(c.defaultSection, key)
}

// StringDefault returns a value associated with the requested key
// or default value if nothing else can be returned.
func (c *INI) StringDefault(key, defaultValue string) string {
	if v, ok := c.String(key); ok {
		return v
	}
	return defaultValue
}

// StringFromSection returns a string value associated with the requested key
// in the requested section.
func (c *INI) StringFromSection(section, key string) (string, bool) {
	// Make sure value does exist.
	v, ok := c.get(section, key)
	if !ok {
		return "", false
	}

	// Check whether it can be type asserted.
	switch v.(type) {
	case string:
		return v.(string), true
	}

	// If it can't assume no associated value exists.
	return "", false
}

// get retrieves an object from the configuration by its section and
// key names.
func (c *INI) get(section, key string) (interface{}, bool) {
	// Make sure requested section does exist.
	if _, ok := c.data[section]; !ok {
		return nil, false
	}

	// Make sure requested value does exist.
	if v, ok := c.data[section][key]; ok {
		return v, true
	}

	// Otherwise, return nothing.
	return nil, false
}
