package types

import (
	"strconv"
)

// Uints represents a slice of uint values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Uints struct {
	base
	Value []uint
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Uints) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Uints) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Uints) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Uints) get(i int) string { return strconv.FormatUint(uint64(s.Value[i]), 10) }

// Alloc allocates a slice of values.
func (s *Uints) alloc() { s.Value = []uint{} }

// Add adds a new value to the slice.
func (s *Uints) add(v string) error {
	u, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, uint(u))
	return nil
}
