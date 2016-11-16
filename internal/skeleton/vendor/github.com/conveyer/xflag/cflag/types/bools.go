package types

import (
	"strconv"
)

// Bools represents a slice of bool values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Bools struct {
	base
	Value []bool
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Bools) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Bools) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Bools) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Bools) get(i int) string { return strconv.FormatBool(s.Value[i]) }

// Alloc allocates a slice of values.
func (s *Bools) alloc() { s.Value = []bool{} }

// Add adds a new value to the slice.
func (s *Bools) add(v string) error {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, b)
	return nil
}
