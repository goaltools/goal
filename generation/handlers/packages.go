package handlers

import (
	"go/build"
	"path/filepath"

	"github.com/anonx/sunplate/reflect"
)

// packages represents packages of controllers. The format is the following:
//	- Import path:
//		- Controllers
type packages map[string]controllers

// controllers stores information about application controllers
// in the following form:
//	- Name of the controller:
//		- Controller representation itself
type controllers map[string]controller

// parent represents embedded struct that should be scanned for
// actions and magic methods.
type parent struct {
	Import  string // Import path of the structure, e.g. "github.com/anonx/sunplate/template" or "".
	Name    string // Name of the structure, e.g. "Template".
	Pointer bool   // Is this type embedded as a pointer or not, i.e. "*Template" or "Template".
}

// controller is a type that represents application controller,
// a structure that has actions.
type controller struct {
	Actions reflect.Funcs // Actions are methods that implement action.Result interface.
	After   *reflect.Func // Magic method that is executed after actions if they return nil.
	Before  *reflect.Func // Magic method that is executed before every action.
	Finally *reflect.Func // Finally is executed at the end of every request no matter what.

	Comments reflect.Comments // A group of comments right above the controller declaration.
	File     string           // Name of the file where this controller is located.
	Parents  []parent         // A list of embedded structs that should be parsed.
}

// processPackage gets an import path of a package, processes it, and
// extracts controllers + actions. Absolute import path is expected,
// i.e. "github.com/anonx/sunplate/controllers" rather than "./controllers".
func (ps packages) processPackage(importPath string) {
	p := reflect.ParseDir(filepath.Join(build.Default.GOPATH, "src", importPath))
	ps[importPath] = ps.extractControllers(p)
}

// scanAnonEmbStructs expects a package and an index of structure in that package.
// It scans the structure looking for fields that are anonymously embedded types.
// If those types are from other packages, they are processed as well.
// As a result a list of all found types in a form of []parent is returned.
func (ps packages) scanAnonEmbStructs(pkg *reflect.Package, i int) (prs []parent) {
	// Iterating over fields of the structure.
	for j := range pkg.Structs[i].Fields {
		// Make sure current field is embedded anonymously,
		// i.e. there is no arg name.
		if pkg.Structs[i].Fields[j].Name != "" {
			continue
		}

		// Add the field to the list of results.
		imp, _ := pkg.Imports.Value(pkg.Structs[i].File, pkg.Structs[i].Fields[j].Type.Package)
		prs = append(prs, parent{
			Import:  imp,
			Name:    pkg.Structs[i].Fields[j].Type.Name,
			Pointer: pkg.Structs[i].Fields[j].Type.Star,
		})

		// Check whether this import has already been processed.
		// If not, do it now.
		if _, ok := ps[imp]; imp != "" && !ok {
			ps.processPackage(imp)
		}
	}
	return
}

// extractControllers gets a reflect.Package type and returns
// a slice of controllers that are found there.
func (ps packages) extractControllers(pkg *reflect.Package) controllers {
	// Initialize a function that will be used for detection of actions.
	action := actionFunc(pkg)

	// Iterating through all available structures and checking
	// whether those structures are controllers (i.e. whether they have actions).
	cs := controllers{}
	for i := 0; i < len(pkg.Structs); i++ {
		// Make sure the structure has methods.
		ms, ok := pkg.Methods[pkg.Structs[i].Name]
		if !ok {
			continue
		}

		// Check whether there are actions among those methods.
		// If there are no any, this is not a controller; ignore it.
		as, count := ms.FilterGroups(action, notMagicMethod, after, before, finally)
		if count == 0 {
			continue
		}

		// Add a new controller to the list of results.
		cs[pkg.Structs[i].Name] = controller{
			Actions: as[0],
			After:   firstFunc(as[1]),
			Before:  firstFunc(as[2]),
			Finally: firstFunc(as[3]),

			Comments: pkg.Structs[i].Comments,
			File:     pkg.Structs[i].File,
			Parents:  ps.scanAnonEmbStructs(pkg, i),
		}
	}
	return cs
}

// firstFunc gets a list of functions and returns the first element of it.
// If the list is empty, nil is returned.
func firstFunc(fs reflect.Funcs) *reflect.Func {
	if len(fs) == 0 {
		return nil
	}
	return &fs[0]
}
