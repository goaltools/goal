// Package strconv contains code generation staff related
// to github.com/goaltools/goal/strconv package.
package strconv

import (
	"errors"
	"fmt"
	"go/ast"

	"github.com/goaltools/goal/internal/reflect"

	"github.com/conveyer/importpath"
)

// ErrUnsupportedType is an error that indicates that there is no conversion
// function for the requested type.
var ErrUnsupportedType = errors.New("unsupported type")

// FnMap is a mapping between type names and appropriate conversion functions.
type FnMap map[string]reflect.Func

// Render gets a package name, name of url.Values variable and an argument, and renders
// an appropriate strconv Function, e.g. `strconv.Int(r.Form, "number")`.
// If the argument is not supported error will be returned as a second argument.
// This is used for code generation.
func (m FnMap) Render(pkgName, vsName string, a reflect.Arg) (string, error) {
	// Make sure argument is of supported type.
	t := a.Type.String()
	f, ok := m[t]
	if !ok {
		return "", ErrUnsupportedType
	}

	// If argument is of slice type (e.g. []int), make sure the argument name
	// ends with []. For illustration, there is a variable:
	//	var names []string
	// In order to bind it, we have to retrieve "names[]" from the request's Form.
	n := a.Name
	if s := a.Type.String(); len(s) > 2 && s[0] == '[' && s[1] == ']' {
		n += "[]"
	}

	// Return the fragment of code we need.
	return fmt.Sprintf(`%s.%s(%s, "%s")`, pkgName, f.Name, vsName, n), nil
}

// Context returns mappings between types that can be parsed using
// strconv package and functions for that conversions.
// All conversion functions meet the following criteria:
// 1. They are exported.
// 2. They expect 3 arguments: url.Values, string, ...int.
// 3. They return 1 argument.
// This is useful for code generation.
func Context() FnMap {
	p, _ := importpath.ToPath("github.com/goaltools/goal/strconv")
	fs := FnMap{}
	pkg := reflect.ParseDir(p, false)
	for i := range pkg.Funcs {
		if !strconvFunc(pkg.Funcs[i]) {
			continue
		}
		fs[pkg.Funcs[i].Results[0].Type.String()] = pkg.Funcs[i]
	}
	return fs
}

// strconvFunc gets a reflect.Func and detects whether it is
// a string conversion function.
func strconvFunc(f reflect.Func) (r bool) {
	// Make sure the function is exported.
	if !ast.IsExported(f.Name) {
		return
	}

	// There are should be 3 arguments: url.Values, string, ...int.
	if len(f.Params) < 3 {
		return
	}
	ps := []string{"url.Values", "string", "...int"}
	for i := range ps {
		if f.Params[i].Type.String() != ps[i] {
			return
		}
	}

	// It should return 1 parameter.
	if len(f.Results) != 1 {
		return
	}
	return true
}
