package reflect

// Func is a type that represents information about a function or method.
type Func struct {
	Args     []Arg    // A list of arguments this function receives.
	Comments []string // Comments that are located right above the function declaration.
	Line     int      // Line of code where this function has been found.
	Name     string   // Name of the function, e.g. "Index" or "About".
	Return   []Arg    // A list of arguments the function returns.
}
