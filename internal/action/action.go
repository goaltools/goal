// Package action provides functions for search of actions
// among methods of a package and checking whether they are
// actions with special meaning (such as Before or After)
// or just plain actions.
package action

import (
	"go/ast"

	"github.com/anonx/sunplate/internal/reflect"
	"github.com/anonx/sunplate/internal/strconv"
	"github.com/anonx/sunplate/log"
)

const (
	// InterfaceImport is a GOPATH to the Result interface
	// that should be implemented by types returned by actions.
	InterfaceImport = "github.com/anonx/sunplate/action"

	// InterfaceName is an interface that should be implemented
	// by types that are returned from actions.
	InterfaceName = "Result"

	// MagicActionBefore is a name of the magic method that will be executed
	// before every action.
	MagicActionBefore = "Before"

	// MagicActionAfter is a name of the magic method that will be executed
	// after every action.
	MagicActionAfter = "After"
)

// StrconvContext is a mapping of supported by strconv types and reflect functions.
var StrconvContext = strconv.Context()

// Func returns a function that may be used to check whether
// specific Func represents an action (or one of magic method) or not.
func Func(pkg *reflect.Package) func(f *reflect.Func) bool {
	// Actions are required to return action.Result as the first argument.
	// actionImportName is used to store information on how the action package is named.
	// For illustration, here is an example:
	//	import (
	//		qwerty "github.com/anonx/sunplate/action"
	//	)
	// In the example above action package will be imported as "qwerty".
	// So, we are saving this name to actionImportName[FILE_NAME_WHERE_WE_IMPORT_THIS]
	// to eliminate the need of iterating through all imports over and over again.
	actionImportName := map[string]string{}

	// Files that should be excluded from search of actions
	// as they do not have action package imported.
	// We are using this as a cache.
	ignoreFiles := map[string]bool{}

	// Return the function that will define whether the function is an action.
	return func(f *reflect.Func) bool {
		// Check whether the file where this method located
		// is ignored due to the lack of action subpackage import.
		if ignoreFiles[f.File] {
			return false
		}

		// Check whether we already know from previous iterations
		// how action subpackage is imported (its name).
		if _, ok := actionImportName[f.File]; !ok {
			// If not, try to find it out.
			n, ok := pkg.Imports.Name(f.File, InterfaceImport)
			if !ok {
				// Action subpackage import path is not found in this file.
				// So, this is not an action method.
				// Ignore this file (and all methods inside it) in future.
				ignoreFiles[f.File] = true
				return false
			}
			actionImportName[f.File] = n // Save the import name to use in future iterations.
		}

		// Check whether the method returns at least one parameter.
		if len(f.Results) == 0 {
			return false
		}

		// Make sure the method we are checking is Exported.
		// Private ones are ignored.
		if !ast.IsExported(f.Name) {
			return false
		}

		// Make sure the first result is of type action.Result.
		correctPackage := f.Results[0].Type.Package == actionImportName[f.File]
		correctName := f.Results[0].Type.Name == InterfaceName
		if !correctPackage || !correctName {
			return false
		}

		// Check whether only builtin types are among input parameters.
		return builtin(f)
	}
}

// builtin gets a function and makes sure its arguments are of builtin type.
// If not, it prints a warning message and returns false.
func builtin(f *reflect.Func) bool {
	recv := ""
	if f.Recv != nil {
		recv = f.Recv.Type.String() + "."
	}
	fn := func(a *reflect.Arg) bool {
		if _, ok := StrconvContext[a.Type.String()]; !ok {
			log.Warn.Printf(
				`Method "%s%s" in file "%s" cannot be treated as action because argument "%s" is of unsupported type "%s".`,
				recv, f.Name, f.File, a.Name, a.Type,
			)
			return false
		}
		return true
	}
	return len(f.Params.Filter(fn)) == len(f.Params)
}

// Before gets a Func and checks whether it is a Before magic action.
func Before(f *reflect.Func) bool {
	if f.Name == MagicActionBefore {
		return true
	}
	return false
}

// After gets a Func and checks whether it is an After magic action.
func After(f *reflect.Func) bool {
	if f.Name == MagicActionAfter {
		return true
	}
	return false
}

// NotMagicAction gets a Func and makes sure it is not a magic action but a usual one.
func NotMagicAction(f *reflect.Func) bool {
	if Before(f) || After(f) {
		return false
	}
	return true
}
