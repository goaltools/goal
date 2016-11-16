package types

import (
	"strconv"
)

// Ints represents a slice of int values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Ints struct {
	base
	Value []int
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Ints) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Ints) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Ints) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Ints) get(i int) string { return strconv.Itoa(s.Value[i]) }

// Alloc allocates a slice of values.
func (s *Ints) alloc() { s.Value = []int{} }

// Add adds a new value to the slice.
func (s *Ints) add(v string) error {
	i, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, i)
	return nil
}
