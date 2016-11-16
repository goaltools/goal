// Package config is a high level package for work with
// the github.com/conveyer/ini parser.
package config

// Interface describes the methods that must be implemented
// by a configuration parser to be compatible with this package.
type Interface interface {
	// New should open, parse, and process the requested file,
	// allocate a new object implementing Interface, and return it.
	// Value of defaultSection should be used by other functions
	// of the object.
	New(file, defaultSection string) (Interface, error)

	// Merge should join values of the new configuration with the current one.
	// Values of the new configuration should have a priority over
	// the current configuration.
	Merge(newConfig Interface)

	// String should return a string value associated with the requested key
	// from the default section. In case result cannot be returned,
	// a false value should be returned as a second argument.
	String(key string) (string, bool)

	// StringDefault is an equivalent of String except if no value can be
	// returned the one received as a default should be returned instead.
	StringDefault(key, defaultValue string) string

	// StringFromSection is an equivalent of String except a value from
	// the requested section (not the default one) is expected.
	StringFromSection(section, key string) (string, bool)

	// Strings as an equivalent of String but for []string values.
	Strings(key string) ([]string, bool)

	// StringsDefault is an equivalent of StringDefault but for []string values.
	StringsDefault(key string, defaultValue []string) []string

	// StringsFromSection is an equivalent of StringFromSection
	// but for []string values.
	StringsFromSection(section, key string) ([]string, bool)
}
