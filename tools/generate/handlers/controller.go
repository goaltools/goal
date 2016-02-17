package handlers

import (
	"strings"

	"github.com/colegion/goal/internal/reflect"
	"github.com/colegion/goal/internal/routes"
)

// controllers stores information about application's controllers
// and its "Init" function.
type controllers struct {
	list []*controller
	init *reflect.Func
}

// controller is a type that represents application controller,
// a structure that has actions.
type controller struct {
	Name string // Name of the controller, e.g. "App" or "Templates".

	Actions reflect.Funcs // Actions are methods that implement action.Result interface.
	After   *reflect.Func // Magic method that is executed after actions if they return nil.
	Before  *reflect.Func // Magic method that is executed before every action.

	Comments reflect.Comments // A group of comments right above the controller declaration.
	File     string           // Name of the file where this controller is located.
	Parents  parents          // A list of embedded structs that should be parsed.

	Fields []field          // A list of fields that require binding.
	Routes [][]routes.Route // Routes concatenated with prefixes. len(Routes) = len(Actions)
}

// Controller gets a controller name and returns it if it
// does exist. Nil is returned otherwise.
func (cs controllers) Controller(name string) *controller {
	for i := 0; i < len(cs.list); i++ {
		if cs.list[i].Name == name {
			return cs.list[i]
		}
	}
	return nil
}

// IgnoredArgs gets an action Func as input parameter
// and returns blank identifiers for parameters
// other than the first one.
// E.g. if the action returns action.Result, error, bool,
// this method will return ", _, _".
// So it can be used during code generation.
func (c controller) IgnoredArgs(f *reflect.Func) (s string) {
	n := len(f.Results) - 1 // Ignore action.Result.
	if n > 0 {
		s = strings.Repeat(", _", n)
	}
	return
}
