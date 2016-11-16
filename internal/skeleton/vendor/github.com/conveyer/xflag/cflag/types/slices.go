package types

// slice is an interface that defines methods that
// every slice type must implement.
type slice interface {
	lenght() int
	get(i int) string

	alloc()
	add(val string) error

	initialized() bool
	requireInit(bool)
}

// base is a type that is wrapped by every real slice
// type of the package. It provides basic fields
// and methods.
type base struct {
	inited bool
}

// initialized is a getter of the "inited" field.
func (b *base) initialized() bool {
	return b.inited
}

// requireInit sets an "initialized" mark so the
// slice can be reinitialized.
func (b *base) requireInit(yes bool) {
	b.inited = !yes
}

// str gets a slice and returns it in a human
// readable format.
func str(s slice) string {
	// If there are no elements in the slice,
	// return nothing.
	l := s.lenght()
	if l == 0 {
		return "[]"
	}

	// Otherwise, prepare a list and return it.
	res := s.get(0)
	for i := 1; i < l; i++ {
		res += "; " + s.get(i)
	}
	return "[" + res + "]"
}

func set(s slice, v string) error {
	// Check whether the end of the input
	// is stated.
	if v == EOI {
		s.requireInit(true)
		return nil
	}

	// If the slice is marked as uninitialized,
	// reallocate it, and set as initialized.
	if !s.initialized() {
		s.requireInit(false)
		s.alloc()
	}

	// Add a new value.
	return s.add(v)
}
