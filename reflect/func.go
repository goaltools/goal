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

// Filter returns a list of functions from members of a list
// fulfilling a condition given by the fn argument.
func (fs Funcs) Filter(fn func(f *Func) bool) (res Funcs) {
	for _, v := range fs {
		if fn(&v) {
			res = append(res, v)
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
