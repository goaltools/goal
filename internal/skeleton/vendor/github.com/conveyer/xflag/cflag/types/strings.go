package types

// Strings represents a slice of string values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Strings struct {
	base
	Value []string
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Strings) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Strings) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Strings) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Strings) get(i int) string { return s.Value[i] }

// Alloc allocates a slice of values.
func (s *Strings) alloc() { s.Value = []string{} }

// Add adds a new value to the slice.
func (s *Strings) add(v string) error {
	s.Value = append(s.Value, v)
	return nil
}
