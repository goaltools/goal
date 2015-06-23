// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"go/ast"
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/reflect"
)

const (
	// ActionInterfaceImport is a GOPATH to the Result interface
	// that should be implemented by types returned by actions.
	ActionInterfaceImport = "github.com/anonx/sunplate/action"

	// ActionInterfaceName is an interface that should be implemented
	// by types that are returned from actions.
	ActionInterfaceName = "Result"

	// MagicMethodBefore is a name of the magic method that will be executed
	// before every action.
	MagicMethodBefore = "Before"

	// MagicMethodAfter is a name of the magic method that will be executed
	// after every action.
	MagicMethodAfter = "After"

	// MagicMethodFinally is a name of the magic method that will be executed
	// after every action no matter what.
	MagicMethodFinally = "Finally"
)

// Controller is a type that represents application controller,
// a structure that has actions.
type Controller struct {
	Actions reflect.Funcs  // Actions are methods that implement action.Result interface.
	After   reflect.Funcs  // Magic methods that are executed after actions if they return nil.
	Before  reflect.Funcs  // Magic methods that are executed before every action.
	Finally reflect.Funcs  // Finally is executed at the end of every request no matter what.
	Struct  reflect.Struct // Structure of the controller (its name, fields, etc).
}

// Start is an entry point of the generate handlers command.
func Start(basePath string, params command.Data) {
	// Generate and save a new package.
	t := output.NewType(
		params.Default("--package", "handlers"), filepath.Join(basePath, "./handlers.go.template"),
	)
	t.CreateDir(params.Default("--output", "./assets/handlers/"))
	t.Extension = ".go" // Save generated file as a .go source.
	t.Context = map[string]interface{}{
		"rootPath": params.Default("--path", "./controllers/"),
	}
	t.Generate()
}

// extractControllers gets a reflect.Package type and returns
// a slice of controllers that are found there.
func extractControllers(pkg *reflect.Package) (cs []Controller) {
	// Actions are required to return action.Result as the first argument.
	// actionImportName is used to store information on how the action package is named.
	// For illustration, here is an example:
	//	import (
	//		qwerty "github.com/anonx/sunplate/action"
	//	)
	// In the example above action package will be imported as "qwerty".
	// So, we are saving this name to actionImportName[FILE_NAME_WHERE_WE_IMPORT_THIS]
	// to eliminate the need of iterating over all imports over and over again.
	actionImportName := map[string]string{}

	// Files that should be excluded from search of actions
	// as they do not have action package imported.
	// We are using this as a cache.
	ignoreFiles := map[string]bool{}

	// isAction gets a pointer to reflect.Func structure
	// and checks whether it represents action or one of magic methods.
	// If so, it return true.
	isAction := func(f *reflect.Func) bool {
		// Check whether the file where this method located
		// is ignored due to the lack of action subpackage import.
		if ignoreFiles[f.File] {
			return false
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

		// Check whether we already know from previous iterations
		// how action subpackage is imported (its name).
		if _, ok := actionImportName[f.File]; !ok {
			// If not, try to find it out.
			n, ok := pkg.Imports.Name(f.File, ActionInterfaceImport)
			if !ok {
				// Action subpackage import path is not found in this file.
				// So, this is not an action method.
				// Ignore this file (and all methods inside it) in future.
				ignoreFiles[f.File] = true
				return false
			}
			actionImportName[f.File] = n // Save the import name to use in future iterations.
		}

		// Make sure the first result is of type action.Result.
		correctPackage := f.Results[0].Type.Package == actionImportName[f.File]
		correctName := f.Results[0].Type.Name == ActionInterfaceName
		if !correctPackage || !correctName {
			return false
		}
		return true
	}

	// isBefore gets a Func and checks whether it is a Before magic method.
	isBefore := func(f *reflect.Func) bool {
		if f.Name == MagicMethodBefore {
			return true
		}
		return false
	}

	// isAfter gets a Func and checks whether it is an After magic method.
	isAfter := func(f *reflect.Func) bool {
		if f.Name == MagicMethodAfter {
			return true
		}
		return false
	}

	// isFinally gets a Func and checks whether it is a Finally magic method.
	isFinally := func(f *reflect.Func) bool {
		if f.Name == MagicMethodFinally {
			return true
		}
		return false
	}

	// notMagicMethod gets a Func and makes sure it is not a magic method but a usual action.
	notMagicMethod := func(f *reflect.Func) bool {
		if isBefore(f) || isAfter(f) || isFinally(f) {
			return false
		}
		return true
	}

	// Iterating through all available structures and checking
	// whether those structures are controllers (i.e. whether they have actions).
	for i := 0; i < len(pkg.Structs); i++ {
		// Make sure the structure has methods.
		ms, ok := pkg.Methods[pkg.Structs[i].Name]
		if !ok {
			continue
		}

		// Check whether there are actions among those methods.
		as := ms.Filter(isAction, notMagicMethod, isAfter, isBefore, isFinally)
		if len(as) == 0 {
			continue
		}

		// Add a new controller to the list of results.
		cs = append(cs, Controller{
			Actions: as[0],
			After:   as[1],
			Before:  as[2],
			Finally: as[3],
			Struct:  pkg.Structs[i],
		})
	}
	return
}
