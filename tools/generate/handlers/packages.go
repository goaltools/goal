package handlers

import (
	a "github.com/colegion/goal/internal/action"
	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/reflect"
	"github.com/colegion/goal/internal/routes"
	"github.com/colegion/goal/utils/path"
)

// packages represents packages of controllers. The format is the following:
//	- Import path:
//		- Controllers
type packages map[string]controllers

// AllInits gets an import path of a main controllers package and
// returns all init function in the order they must be called. I.e. grandparents
// first, then parents, then children, and so forth.
func (ps packages) AllInits(importPath string) (fs reflect.Funcs) {
	// Make sure the input import path belongs to a controllers package.
	cs, ok := ps[importPath]
	if !ok {
		return nil
	}

	// Generate a list of imports to check next by visiting parents
	// of every's controller.
	checked := map[string]bool{}
	parents := []string{}
	for i := range cs.list { // Visiting every controller of the package.
		for j := range cs.list[i].Parents.list { // Checking every parent of every controller.
			// Make sure current parent's import is not in the list yet.
			imp := cs.list[i].Parents.list[j].Import
			if checked[imp] || imp == importPath {
				continue
			}

			// If not, add it and mark as checked.
			checked[imp] = true
			parents = append(parents, imp)
		}
	}

	// Iterate over all extracted import paths and get their inits.
	for i := range parents {
		fs = append(fs, ps.AllInits(parents[i])...)
	}

	// Add current package's init, if presented, to the end of the result.
	if cs.init != nil {
		fs = append(fs, *cs.init)
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
	cs := ps.extractControllers(importPath, p, prefs)
	if len(cs.list) > 0 {
		ps[importPath] = controllers{
			list: cs.list,
			init: ps.extractInitFunc(p),
		}
	}
}

// scanFields gets an import path, a package itself,
// and an index of structure in that package.
// It scans the structure looking for two kinds of fields:
// anonymously embedded types and named fields with special tags.
// Every anonymously embedded type is checked recursively regarding being a controller.
// As a result a list of all found fields with the tags and
// types in a form of []parent are returned.
func (ps packages) scanFields(impPath string, pkg *reflect.Package, i int) ([]field, parents) {
	// Allocate result []field and parents types.
	fs := []field{}
	prs := parents{
		childImport: impPath,
		list:        []parent{},
	}

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
		if imp == "" { // If the import is empty, the embedded structure is from the same package.
			imp = impPath
		}
		prs.list = append(prs.list, parent{
			Import: imp,
			Name:   pkg.Structs[i].Fields[j].Type.Name,
		})

		// Check whether this import has already been processed.
		// If not and this is not the import we've got above, do it now.
		if _, ok := ps[imp]; !ok && imp != impPath {
			ps.processPackage(imp, routes.ParseTag(pkg.Structs[i].Fields[j].Tag))
		}
	}
	return fs, prs
}

// extractControllers gets an import path, a parsed reflect.Package itself, and returns
// a slice of controllers that are found there.
func (ps packages) extractControllers(impPath string, pkg *reflect.Package, prefs routes.Prefixes) controllers {
	// Initialize function that will be used for detection of actions.
	action := a.Func(pkg)

	// Iterating through all available structures and checking
	// whether those structures are controllers (i.e. whether they have actions).
	cs := controllers{
		list: []*controller{},
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
		fs, prs := ps.scanFields(impPath, pkg, i)

		// Add a new controller to the list of results.
		cs.list = append(cs.list, &controller{
			Name: pkg.Structs[i].Name,

			Actions: as[0],
			After:   firstFunc(as[1]),
			Before:  firstFunc(as[2]),

			Comments: pkg.Structs[i].Comments,
			File:     pkg.Structs[i].File,
			Parents:  prs,

			Fields: fs,
			Routes: rs,
		})
	}
	return cs
}

// extractInitFunction returns an "Init" function found in a controller package,
// if it does exist.
func (ps packages) extractInitFunc(pkg *reflect.Package) *reflect.Func {
	// Iterate over all available functions.
	res, _ := pkg.Funcs.FilterGroups(func(f *reflect.Func) bool {
		// Find the one with name "Init".
		if f.Name != "Init" {
			return false
		}

		// Make sure it is a function rather than a method.
		if f.Recv != nil {
			return false // This must never happen, pkg.Funcs does not contain methods.
		}

		// Make sure the function takes one argument.
		if len(f.Params) != 1 {
			return false
		}

		// The argument must be of types Values.
		if f.Params[0].Type.Name != "Values" {
			return false
		}

		// The Values type must be from package "net/url".
		v, _ := pkg.Imports.Value(f.File, f.Params[0].Type.Package)
		if v != "net/url" {
			return false
		}

		// If we've achieved this point, "Init" function is found.
		log.Trace.Printf(`Magic "%s" function will be added to generated "%s" file.`, f.Name, f.File)
		return true
	}, func(f *reflect.Func) bool {
		// Return every of the found "Init" functions.
		return true
	})
	// If result is not empty, return the first function.
	if len(res[0]) > 0 {
		return firstFunc(res[0])
	}
	return nil
}

// firstFunc gets a list of functions and returns the first element of it.
// If the list is empty, nil is returned.
func firstFunc(fs reflect.Funcs) *reflect.Func {
	if len(fs) == 0 {
		return nil
	}
	return &fs[0]
}
