package handlers

import (
	"fmt"
	"go/ast"
	r "reflect"
	"strings"

	a "github.com/goaltools/goal/internal/action"
	"github.com/goaltools/goal/internal/log"
	"github.com/goaltools/goal/internal/reflect"
	"github.com/goaltools/goal/internal/routes"
	"github.com/goaltools/goal/utils/path"
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
	Import string // Import path of the structure, e.g. "github.com/goaltools/goal/template" or "".
	Name   string // Name of the structure, e.g. "Template".
}

// field represents a field of a structure that must be automatically binded.
type field struct {
	Name string // Name of the field, e.g. "Request".
	Type string // Type of the binding, e.g. "request" or "action".
}

// controller is a type that represents application controller,
// a structure that has actions.
type controller struct {
	Actions reflect.Funcs // Actions are methods that implement action.Result interface.
	After   *reflect.Func // Magic method that is executed after actions if they return nil.
	Before  *reflect.Func // Magic method that is executed before every action.

	Comments reflect.Comments // A group of comments right above the controller declaration.
	File     string           // Name of the file where this controller is located.
	Parents  parents          // A list of embedded structs that should be parsed.

	Fields []field          // A list of fields that require binding.
	Routes [][]routes.Route // Routes concatenated with prefixes. len(Routes) = len(Actions)
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

// processPackage gets an import path of a package and its
// route prefixes, processes this data, and
// extracts controllers + actions.
func (ps packages) processPackage(importPath string, prefs routes.Prefixes) {
	log.Trace.Printf(`Parsing "%s"...`, importPath)
	dir, err := path.ImportToAbsolute(importPath)
	if err != nil {
		log.Error.Panic(err)
	}
	p := reflect.ParseDir(dir, false)
	cs := ps.extractControllers(p, prefs)
	if len(cs.data) > 0 {
		ps[importPath] = controllers{
			data: cs.data,
			init: ps.extractInitFunc(p),
		}
	}
}

// needBindingField gets a package, an index of struct and index of field
// in the struct. The field is checked whether it has a reserved tag
// and it is of correct type.
func (ps packages) needBindingField(pkg *reflect.Package, i, j int) *field {
	f := &field{}
	t := pkg.Structs[i].Fields[j]
	switch st := r.StructTag(t.Tag).Get("bind"); st {
	case "response":
		// Make sure "http" package is imported.
		n, ok := pkg.Imports.Name(pkg.Structs[i].File, "net/http")
		if !ok || t.Type.String() != fmt.Sprintf("%s.ResponseWriter", n) {
			log.Warn.Printf(
				`Field "%s" in controller "%s" cannot be binded. Response must be of type "(net/http).ResponseWriter".`,
				t.Name, pkg.Structs[i].Name,
			)
			return nil
		}
		f.Type = st
	case "request":
		// Make sure "http" package is imported.
		n, ok := pkg.Imports.Name(pkg.Structs[i].File, "net/http")
		if !ok || t.Type.String() != fmt.Sprintf("*%s.Request", n) {
			log.Warn.Printf(
				`Field "%s" in controller "%s" cannot be binded. Request must be of type "*(net/http).Request".`,
				t.Name, pkg.Structs[i].Name,
			)
			return nil
		}
		f.Type = st
	case "controller":
		if t.Type.String() != "string" {
			log.Warn.Printf(
				`Field "%s" in controller "%s" cannot be binded. Controller name must be of type "string".`,
				t.Name, pkg.Structs[i].Name,
			)
			return nil
		}
		f.Type = "controller"
	case "action":
		if t.Type.String() != "string" {
			log.Warn.Printf(
				`Field "%s" in controller "%s" cannot be binded. Action name must be of type "string".`,
				t.Name, pkg.Structs[i].Name,
			)
			return nil
		}
		f.Type = st
	default:
		return nil
	}
	f.Name = t.Name
	if !ast.IsExported(f.Name) {
		log.Warn.Printf(
			`Field "%s" in controller "%s" must be public in order to be binded.`,
			f.Name, pkg.Structs[i].Name,
		)
		return nil
	}
	log.Trace.Printf(`Field %s will be binded to "%s".`, f.Name, f.Type)
	return f
}

// scanFields expects a package and an index of structure in that package.
// It scans the structure looking for two kinds of fields:
// anonymously embedded types and named fields with special tags.
// Every anonymously embedded type is checked recursively regarding being a controller.
// As a result a list of all found fields with the tags and
// types in a form of []parent are returned.
func (ps packages) scanFields(pkg *reflect.Package, i int) (fs []field, prs []parent) {
	// Iterating over fields of the structure.
	for j := range pkg.Structs[i].Fields {
		// Check whether the field requires binding.
		if f := ps.needBindingField(pkg, i, j); f != nil {
			fs = append(fs, *f)
			continue
		}

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
		p, _ := path.CleanImport(imp)
		prs = append(prs, parent{
			Import: p,
			Name:   pkg.Structs[i].Fields[j].Type.Name,
		})

		// Check whether this import has already been processed.
		// If not, do it now.
		if _, ok := ps[imp]; imp != "" && !ok {
			ps.processPackage(p, routes.ParseTag(pkg.Structs[i].Fields[j].Tag))
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
		if f.Params[0].Type.Name != "Values" {
			return false
		}
		v, _ := pkg.Imports.Value(f.File, f.Params[0].Type.Package)
		if v != "net/url" {
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
func (ps packages) extractControllers(pkg *reflect.Package, prefs routes.Prefixes) controllers {
	// Initialize function that will be used for detection of actions.
	action := a.Func(pkg)

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

		// Check whether there are actions among those methods.
		rs := [][]routes.Route{}
		as, count := ms.FilterGroups(func(f *reflect.Func) bool {
			// Ignore non-actions.
			res := action(f)
			if !res {
				return false
			}

			// Skip non-regular actions.
			if !a.Regular(f) {
				return true
			}

			// Parse action's routes.
			if r := prefs.ParseRoutes(pkg.Structs[i].Name, f); len(r) > 0 {
				rs = append(rs, r)
			}
			return true
		}, a.Regular, a.After, a.Before)

		// If there are no any, this is not a controller; ignore it.
		if count == 0 {
			continue
		}

		// Parse parent controllers and fields that require binding.
		fs, prs := ps.scanFields(pkg, i)

		// Add a new controller to the list of results.
		cs.data[pkg.Structs[i].Name] = controller{
			Actions: as[0],
			After:   firstFunc(as[1]),
			Before:  firstFunc(as[2]),

			Comments: pkg.Structs[i].Comments,
			File:     pkg.Structs[i].File,
			Parents:  prs,

			Fields: fs,
			Routes: rs,
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
