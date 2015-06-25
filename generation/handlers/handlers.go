// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
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
	Actions reflect.Funcs  // Actions are methods that implement action.Result interface.
	After   reflect.Funcs  // Magic methods that are executed after actions if they return nil.
	Before  reflect.Funcs  // Magic methods that are executed before every action.
	Finally reflect.Funcs  // Finally is executed at the end of every request no matter what.
	Parents []parent       // A list of embedded structs that should be parsed.
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

func (t *packages) processPackage(importPath string) {
}

// scanAnonEmbStructs expects a package and an index of structure in that package.
// It scans the structure looking for fields that are anonymously embedded types.
// If those types are from other packages, they are processed as well.
// As a result a list of all found types in a form of []parent is returned.
func scanAnonEmbStructs(pkg *reflect.Package, i int) (ps []parent) {
	// Iterating over fields of the structure.
	for j := range pkg.Structs[i].Fields {
		// Make sure current field is embedded anonymously,
		// i.e. there is no arg name.
		if pkg.Structs[i].Fields[j].Name != "" {
			continue
		}

		// Add the field to the list of results.
		imp, _ := pkg.Imports.Value(pkg.Structs[i].File, pkg.Structs[i].Fields[j].Type.Package)
		ps = append(ps, parent{
			Import:  imp,
			Name:    pkg.Structs[i].Fields[j].Type.Name,
			Pointer: pkg.Structs[i].Fields[j].Type.Star,
		})
	}
	return
}

// extractControllers gets a reflect.Package type and returns
// a slice of controllers that are found there.
func extractControllers(pkg *reflect.Package) (cs controllers) {
	// Initialize a function that will be used for detection of actions.
	action := actionFunc(pkg)

	// Iterating through all available structures and checking
	// whether those structures are controllers (i.e. whether they have actions).
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
			After:   as[1],
			Before:  as[2],
			Finally: as[3],
			Parents: scanAnonEmbStructs(pkg, i),
			Struct:  pkg.Structs[i],
		}
	}
	return
}
