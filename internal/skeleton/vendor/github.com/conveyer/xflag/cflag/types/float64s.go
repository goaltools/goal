package types

import (
	"strconv"
)

// Float64s represents a slice of float64 values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Float64s struct {
	base
	Value []float64
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Float64s) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Float64s) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Float64s) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Float64s) get(i int) string { return strconv.FormatFloat(s.Value[i], 'f', -1, 64) }

// Alloc allocates a slice of values.
func (s *Float64s) alloc() { s.Value = []float64{} }

// Add adds a new value to the slice.
func (s *Float64s) add(v string) error {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, f)
	return nil
}
