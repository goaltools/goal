package reflect

import (
	"go/ast"
)

// Func is a type that represents information about a function or method.
type Func struct {
	Comments []string // Comments that are located right above the function declaration.
	Name     string   // Name of the function, e.g. "Index" or "About".
	Params   []Arg    // A list of arguments this function receives.
	Results  []Arg    // A list of arguments the function returns.
	Recv     *Arg     // Receiver if it is a method and nil otherwise.
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
