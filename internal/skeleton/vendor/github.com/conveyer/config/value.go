package config

// Value implements ValueInterface.
type Value struct {
	data interface{}
}

// NewValue allocates and returns a new Value.
func NewValue(data interface{}) *Value {
	return &Value{data}
}

// Interface returns an inner value as interface{}.
func (v *Value) Interface() interface{} {
	return v.data
}

// String returns a string representation of the Data or false
// as a second argument otherwise.
func (v *Value) String() (string, bool) {
	s, ok := v.data.(string)
	return s, ok
}

// StringDefault is an equivalent of String that returns the specified default
// value if no other string can be returned.
func (v *Value) StringDefault(defaultValue string) string {
	s, ok := v.String()
	if !ok {
		return defaultValue
	}
	return s
}

// Strings is an equivalent of String but for []string values.
func (v *Value) Strings() ([]string, bool) {
	ss, ok := v.data.([]string)
	return ss, ok
}

// StringsDefault is an equivalent of StringDefault but for []string values.
func (v *Value) StringsDefault(defaultValue []string) []string {
	ss, ok := v.Strings()
	if !ok {
		return defaultValue
	}
	return ss
}
