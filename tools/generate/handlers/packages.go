package handlers

import (
	"fmt"
	"strings"

	a "github.com/colegion/goal/internal/action"
	m "github.com/colegion/goal/internal/method"
	"github.com/colegion/goal/internal/path"
	"github.com/colegion/goal/internal/reflect"
	"github.com/colegion/goal/log"
)

// packages represents packages of controllers. The format is the following:
//	- Import path:
//		- Controllers
type packages map[string]controllers

// controllers stores information about application controllers
// in the following form:
//	- Name of the controller:
//		- Controller representation itself
// and their Init functions.
type controllers struct {
	data map[string]controller
	init *reflect.Func
}

// parents represents a set of parent controllers.
type parents []parent

// parent represents embedded struct that should be scanned for
// actions and magic methods.
type parent struct {
	ID     int    // Unique number that is used for generation of import names.
	Import string // Import path of the structure, e.g. "github.com/colegion/goal/template" or "".
	Name   string // Name of the structure, e.g. "Template".
}

// controller is a type that represents application controller,
// a structure that has actions.
type controller struct {
	Actions   reflect.Funcs // Actions are methods that implement action.Result interface.
	After     *reflect.Func // Magic method that is executed after actions if they return nil.
	Before    *reflect.Func // Magic method that is executed before every action.
	Finally   *reflect.Func // Finally is executed at the end of every request no matter what.
	Initially *reflect.Func // Initially is executed at the beginning of every request.

	Comments reflect.Comments // A group of comments right above the controller declaration.
	File     string           // Name of the file where this controller is located.
	Parents  parents          // A list of embedded structs that should be parsed.
}

// Package returns a unique package name that may be used in templates
// concatenated with some arbitarry suffix strings.
// If parent is from the local package, empty string will be returned.
// This method is useful to generate in templates things like:
//	uniquePkgName "github.com/user/project"
// and:
//	uniquePkgName.Application.Index() // Package name and dot suffix.
func (p parent) Package(suffixes ...string) string {
	if p.Import == "" {
		return ""
	}
	s := fmt.Sprintf("c%d", p.ID)
	for i := range suffixes {
		s += suffixes[i]
	}
	return s
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

// processPackage gets an import path of a package, processes it, and
// extracts controllers + actions.
func (ps packages) processPackage(importPath string) {
	log.Trace.Printf(`Parsing "%s"...`, importPath)
	p := reflect.ParseDir(path.PackageDir(importPath), false)
	cs := ps.extractControllers(p)
	if len(cs.data) > 0 {
		ps[importPath] = controllers{
			data: cs.data,
			init: ps.extractInitFunc(p),
		}
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

func (ps packages) extractInitFunc(pkg *reflect.Package) *reflect.Func {
	res, _ := pkg.Funcs.FilterGroups(func(f *reflect.Func) bool {
		if f.Name != "Init" {
			return false
		}
		if f.Recv != nil {
			return false
		}
		if len(f.Params) != 1 {
			return false
		}
		impName, ok := pkg.Imports.Name(f.File, f.Params[0].Type.Package)
		if !ok || f.Params[0].Type.Package != impName || f.Params[0].Type.Star {
			return false
		}
		if f.Params[0].Type.Name != "Getter" {
			return false
		}
		log.Trace.Printf(`Magic "%s" function will be added to generated "%s" file.`, f.Name, f.File)
		return true
	}, func(f *reflect.Func) bool {
		return true
	})
	if len(res[0]) > 0 {
		return &res[0][0]
	}
	return nil
}

// extractControllers gets a reflect.Package type and returns
// a slice of controllers that are found there.
func (ps packages) extractControllers(pkg *reflect.Package) controllers {
	// Initialize functions that will be used for detection of actions
	// and magic methods.
	action := a.Func(pkg)
	method := m.Func(pkg)

	// Iterating through all available structures and checking
	// whether those structures are controllers (i.e. whether they have actions).
	cs := controllers{
		data: map[string]controller{},
	}
	for i := 0; i < len(pkg.Structs); i++ {
		// Make sure the structure has methods.
		ms, ok := pkg.Methods[pkg.Structs[i].Name]
		if !ok {
			continue
		}

		// Check whether there are actions and/or magic method among those methods.
		as, count1 := ms.FilterGroups(action, a.Regular, a.After, a.Before)
		mms, count2 := ms.FilterGroups(method, m.Initially, m.Finally)

		// If there are no any, this is not a controller; ignore it.
		if count1 == 0 && count2 == 0 {
			continue
		}

		// Add a new controller to the list of results.
		cs.data[pkg.Structs[i].Name] = controller{
			Actions:   as[0],
			After:     firstFunc(as[1]),
			Before:    firstFunc(as[2]),
			Initially: firstFunc(mms[0]),
			Finally:   firstFunc(mms[1]),

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
