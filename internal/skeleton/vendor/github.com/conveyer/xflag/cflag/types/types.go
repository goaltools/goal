// Package types implements flag.Value interface for a number of
// complex types such as []string, []int, []bool, etc.
// So, it is possible to use them with the standard flag package
// that has a limited number of supported types out of the box.
// NB: Set methods of this package's types work pretty much like Add.
// To make it behave as Set, call Set(EOI) after you are done
// adding new values.
package types

const (
	// EOI (end of input) is a special value of a string that can be passed
	// to the Set method of any type to mark it as uninitialized.
	// This is necessary as Set in current implementation works as Add. I.e.:
	//	Set("a")
	//	Set("b")
	//	Set("c")
	// The code above will produce []string{"a", "b", "c"},
	// not just []string{"c"}.
	// But sometimes original behaviour of Set is required. E.g.
	// we inserted "a", "b", and "c" first and now
	// want to redefine the values to "x", "y", "z".
	// To do it, the following call is required after the code above:
	//	Set(EOI)
	// And then we can readd all the required values:
	//	Set("x")
	//	Set("y")
	//	Set("z")
	EOI = "\t\n\"\000\"\n\t"
)
