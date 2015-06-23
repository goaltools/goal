package reflect

import (
	"go/ast"
)

// Funcs is a type that represents a list of functions.
type Funcs []Func

// Func is a type that represents information about a function or method.
type Func struct {
	Comments Comments // Comments that are located right above the function declaration.
	File     string   // Name of the file where the function is located.
	Name     string   // Name of the function, e.g. "Index" or "About".
	Params   Args     // A list of arguments this function receives.
	Recv     *Arg     // Receiver if it is a method and nil otherwise.
	Results  Args     // A list of arguments the function returns.
}

// Filter gets a condition function and a number of group functions.
// It cuts off those Funcs that do not satisfy condition.
// And then groups the rest of them.
// For illustration:
//	res := myFuncs.Filter(isExported, withArguments, withoutArguments)
// The result will be:
//	// All this functions are satisfying isExported condition.
//	[]Funcs{
//		Funcs{ these are functions withArguments },
//		Funcs{ these are functions withoutArguments },
//	}
func (fs Funcs) Filter(cond func(f *Func) bool, groups ...func(f *Func) bool) []Funcs {
	res := make([]Funcs, len(groups))

	// Iterating over all available Funcs.
	for _, f := range fs {
		// Make sure they satisfy requested condition.
		if !cond(&f) {
			continue
		}

		// Group them into categories.
		for i := range groups {
			if groups[i](&f) {
				res[i] = append(res[i], f)
			}
		}
	}

	return res
}

// processFuncDecl receives an ast function declaration and
// transforms it into Func structure that is returned.
func processFuncDecl(decl *ast.FuncDecl) *Func {
	// Check whether there is a receiver.
	var recv *Arg
	args := processFieldList(decl.Recv)
	if len(args) > 0 {
		recv = &args[0]
	}

	return &Func{
		Comments: processCommentGroup(decl.Doc),
		Name:     decl.Name.Name,
		Params:   processFieldList(decl.Type.Params),
		Results:  processFieldList(decl.Type.Results),
		Recv:     recv,
	}
}
