package types

import (
	"time"
)

// Durations represents a slice of int values,
// a type that implements flag.Value and thus
// can be used with flag.Var.
type Durations struct {
	base
	Value []time.Duration
}

//
// Methods below implement flag.Value interface.
//

// String returns the type in a human readable format.
func (s *Durations) String() string { return str(s) }

// Set gets a string value and adds it to the slice.
func (s *Durations) Set(v string) error { return set(s, v) }

//
// Methods below implement slice interface.
//

// Len returns a number of elements in the slice.
func (s *Durations) lenght() int { return len(s.Value) }

// Get returns a value by its index.
func (s *Durations) get(i int) string { return s.Value[i].String() }

// Alloc allocates a slice of values.
func (s *Durations) alloc() { s.Value = []time.Duration{} }

// Add adds a new value to the slice.
func (s *Durations) add(v string) error {
	d, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	s.Value = append(s.Value, d)
	return nil
}
