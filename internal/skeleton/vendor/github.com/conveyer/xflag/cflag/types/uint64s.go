package types

import (
	"strconv"
)

// Uint64s represents a slice of uint values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Uint64s struct {
	base
	Value []uint64
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Uint64s) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Uint64s) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Uint64s) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Uint64s) get(i int) string { return strconv.FormatUint(s.Value[i], 10) }

// Alloc allocates a slice of values.
func (s *Uint64s) alloc() { s.Value = []uint64{} }

// Add adds a new value to the slice.
func (s *Uint64s) add(v string) error {
	u, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, u)
	return nil
}
