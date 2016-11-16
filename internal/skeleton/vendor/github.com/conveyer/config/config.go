// Package config provides functions that are used for
// parsing of confguration files.
// Package config defines an
package config

// Demonstration of API:
//	// Allocate a new configuration.
//	c, err := MyConfig.New("/path/to/config.ini", "mySection")
//
//	// Merge a new config to the current one.
//	err = c.Join("/path/to/another/config.ini")
//
//	// Extract values from the "mySection" requested above.
//	s, ok := c.Value("myKey").String()
//	s1 := c.Value("myKey1").StringDefault("default value")
//
//	// Extract values from a custom section.
//	s, ok = c.Value("mySection1", "myKey").String()
//	s1 = c.Value("mySection1", "myKey1").StringDefault("default value")

// Interface describes the methods that must be implemented by every
// config parser in order to be compatible with the package.
type Interface interface {
	// New should try to open, parse, and process the requested file, allocate
	// a new config, and return it if that's possible. A non-nil error is expected
	// as a second argument otherwise.
	// Argument pathPrefix specifies section / object where method Value
	// will look for values to return.
	New(file string, pathPrefix ...string) (Interface, error)

	// Join should merge the current configuration with the content of the
	// requested configuration file. Values of the input file should have a
	// higher priority than the values of the current config and thus
	// override them in case similar keys are used.
	Join(file string) error

	// Value should return a value associated with "pathPrefix + path".
	Value(path ...string) ValueInterface

	// ValuePrefixless should be an equivalent of Value that uses "path" as a key,
	// ignoring the pathPrefix defined in the constructor.
	ValuePrefixless(path ...string) ValueInterface
}

// ValueInterface describes a type of value that is expected
// to be returned from configuration.
type ValueInterface interface {
	// Interface should return an inner value as an interface{}.
	Interface() interface{}

	// String should return the value as a string or false as a second
	// argument is expected if that is not possible.
	String() (string, bool)

	// StringDefault is an equiavalent of String that should return a specified
	// default value if no conversion to string is possible.
	StringDefault(string) string

	// Strings is an equivalent of String but for []string data.
	Strings() ([]string, bool)

	// StringsDefault is an equivalent of StringDefault but for []string data.
	StringsDefault([]string) []string
}
