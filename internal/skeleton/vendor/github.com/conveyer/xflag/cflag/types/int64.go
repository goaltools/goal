package types

import (
	"strconv"
)

// Int64s represents a slice of int64 values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Int64s struct {
	base
	Value []int64
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Int64s) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Int64s) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Int64s) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Int64s) get(i int) string { return strconv.FormatInt(s.Value[i], 10) }

// Alloc allocates a slice of values.
func (s *Int64s) alloc() { s.Value = []int64{} }

// Add adds a new value to the slice.
func (s *Int64s) add(v string) error {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, i)
	return nil
}
