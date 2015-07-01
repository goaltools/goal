package handlers

import (
	"fmt"
	"go/build"
	"path/filepath"

	"github.com/anonx/sunplate/path"
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

// parents represents a set of parent controllers.
type parents []parent

// parent represents embedded struct that should be scanned for
// actions and magic methods.
type parent struct {
	ID     int    // Unique number that is used for generation of import names.
	Import string // Import path of the structure, e.g. "github.com/anonx/sunplate/template" or "".
	Name   string // Name of the structure, e.g. "Template".
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
	Parents  parents          // A list of embedded structs that should be parsed.
}

// Package returns a unique package name that may be used in templates
// concatenated with some arbitarry suffix strings.
func (p parent) Package(suffices ...string) string {
	s := fmt.Sprintf("c%d", p.ID)
	for i := range suffices {
		s += suffices[i]
	}
	return s
}

// processPackage gets an import path of a package, processes it, and
// extracts controllers + actions.
func (ps packages) processPackage(importPath string) {
	p := reflect.ParseDir(filepath.Join(build.Default.GOPATH, "src", importPath))
	cs := ps.extractControllers(p)
	if len(cs) > 0 {
		ps[importPath] = cs
	}
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

		// Ensure the struct is embedded as a pointer.
		if !pkg.Structs[i].Fields[j].Type.Star {
			continue
		}

		// Add the field to the list of results.
		imp, _ := pkg.Imports.Value(pkg.Structs[i].File, pkg.Structs[i].Fields[j].Type.Package)
		prs = append(prs, parent{
			Import: path.AbsoluteImport(imp),
			Name:   pkg.Structs[i].Fields[j].Type.Name,
		})

		// Check whether this import has already been processed.
		// If not, do it now.
		if _, ok := ps[imp]; imp != "" && !ok {
			ps.processPackage(path.AbsoluteImport(imp))
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
		as, count := ms.FilterGroups(action, notMagicAction, after, before, finally)
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
