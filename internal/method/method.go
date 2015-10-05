// Package method provides functions for search of Initially and Finally
// methods with special meaning among a list of methods of received package.
package method

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/colegion/goal/internal/reflect"
)

const (
	// InitiallyName is a name of the method with special meaning that is
	// started automatically with any request before any other
	// magic or regular action.
	InitiallyName = "Initially"

	// FinallyName is similar to Initially but is started at the very end.
	// It will be executed even if some of the actions or Initially
	// methods panics.
	FinallyName = "Finally"

	netHTTPImport = "net/http"
	request       = "Request"
	respWriter    = "ResponseWriter"
	result        = "bool"
)

// Func returns a function that gets a Func and checks
// whether it is valid Initially or Finally method.
// The following criterion must be met to make sure specific function
// is a magic method with special meaning:
// 1. It is exported.
// 2. It gets 2 input arguments: http.ResponseWriter and *http.Request;
// 3. It returns 1 result: bool.
// It is assumed Func with receiver is received (this isn't checked).
func Func(pkg *reflect.Package) func(f *reflect.Func) bool {
	// Magic methods are required to get http.ResponseWriter and *http.Request as parameters.
	// importName is used to store information on how the http package is named.
	// For illustration, here is an example:
	//	import (
	//		qwerty "net/http"
	//	)
	// In the example above http package will be imported as "qwerty".
	// So, we are saving this name to importName[FILE_NAME_WHERE_WE_IMPORT_THIS]
	// to eliminate the need of iterating through all imports over and over again.
	importName := map[string]string{}

	// Files that should be excluded from search of magic methods
	// as they do not have http package imported.
	// We are using this as a cache.
	ignoreFiles := map[string]bool{}

	return func(f *reflect.Func) bool {
		// Check whether the file where this method located
		// is ignored due to the lack of http package import.
		if ignoreFiles[f.File] {
			return false
		}

		// Make sure this is not a test file.
		if strings.HasSuffix(f.File, "_test.go") {
			return false
		}

		// Check whether we already know from previous iterations
		// how http package is imported (its name).
		if _, ok := importName[f.File]; !ok {
			// If not, try to find it out.
			n, ok := pkg.Imports.Name(f.File, netHTTPImport)
			if !ok {
				// http package import path is not found in this file.
				// So, this is not an magic method.
				// Ignore this file (and all methods inside it) in future.
				ignoreFiles[f.File] = true
				return false
			}
			importName[f.File] = n // Save the import name to use in future iterations.
		}

		// Make sure the function is exported.
		if !ast.IsExported(f.Name) {
			return false
		}

		// Make sure there are 3 arguments and they are
		// http.ResponseWriter, *http.Request, []interface{}.
		if len(f.Params) != 3 {
			return false
		}
		respWr := f.Params[0].Type.String() == fmt.Sprintf(
			"%s.%s", importName[f.File], respWriter,
		)
		req := f.Params[1].Type.String() == fmt.Sprintf(
			"*%s.%s", importName[f.File], request,
		)
		sl := f.Params[2].Type.String() == "[]string"
		if !respWr || !req || !sl {
			return false
		}

		// Check whether there is only one result and it
		// is of the type we need.
		if len(f.Results) != 1 || f.Results[0].Type.String() != result {
			return false
		}

		return true
	}
}

// Initially gets a magic Func and checks whether it is an Initially method.
func Initially(f *reflect.Func) bool {
	if f.Name == InitiallyName {
		return true
	}
	return false
}

// Finally gets a magic Func and checks whether it is a Finally method.
func Finally(f *reflect.Func) bool {
	if f.Name == FinallyName {
		return true
	}
	return false
}
