package handlers

import (
	"fmt"
	"sort"
	"strings"
)

// parents represents a set of relative controllers.
type parents struct {
	childImport string
	list        []parent
}

// parent represents embedded struct that should be scanned for
// actions and magic methods and/or the opposite of that (i.e.
// all controllers that import the current one).
type parent struct {
	Import string // Import path of the structure, e.g. "github.com/colegion/goal/template".
	Name   string // Name of the structure, e.g. "Template".
}

// parentController represents information about a parent controller,
// how to allocate and access it.
type parentController struct {
	// Accessor is a unique name. It may be used for generation of imports:
	//	import (
	//		uniquePkgName "github.com/path/to/the/project/controllers"
	//	)
	// The package may be accessed later as:
	//	uniquePkgName.Application.Index()
	// The Accessor is empty for the local package.
	Accessor string

	// Controller is a parent controller itself.
	Controller *controller

	// Prefix represents fields that must be accessed before the current controller's
	// name. E.g. "Child.Parent." for a "GrandParent" controller.
	// So, when allocating we'll have to access "Child.Parent.GrandParent".
	Prefix string

	// instance stores information about a controller that has already been allocated
	// and is of the same type (we're allocating every type just once).
	//	c := App{}
	//	c.A = A{}
	//	c.A.B = B{}
	//	c.B = c.A.B // instance in this case is equal to "A.B".
	// If no other controllers of the same type were allocated before, instance is empty.
	instance string
}

// parentControllers is a set of parentControllers in the order their special
// Before and/or After methods must be called. For allocation, reverse it.
type parentControllers []parentController

// Called returns true if special methods on the controller of the same type
// were already called.
func (pc parentController) Called() bool {
	return pc.instance != ""
}

// Access generates code for accessing a parent controller.
func (pc parentController) Access() string {
	return pc.Prefix + pc.Controller.Name
}

// Allocate gets a variable name and return code for the parent
// controller's allocation. E.g. "pkgName.Controller" or "Controller"
// or "Child.Parent.Controller". The way to use it in template:
//	c := App{}
//	... = {{ parentController.Allocate "c" "ctrPackage" }}{}
func (pc parentController) Allocate(varName, currContrPackage string) string {
	// Check whether there is already and instance of this type.
	if pc.instance != "" {
		return varName + "." + pc.instance // E.g. "c.Child.Parent".
	}
	// Find out from what package the controller must be imported.
	accessor := currContrPackage
	if pc.Accessor != "" {
		accessor = pc.Accessor
	}
	return "&" + accessor + "." + pc.Controller.Name // E.g. "pkgName.Parent".
}

// Reverse returns a reversed version of parent controllers.
// They must be reversed in order to be allocated.
func (pcs parentControllers) Reverse() (res parentControllers) {
	for i := len(pcs) - 1; i >= 0; i-- {
		res = append(res, pcs[i])
	}
	return
}

// Imports returns all import paths of parent controllers. It may be
// includes in the templates as follows:
//	import (
//		{{ .parentControllers.Imports }}
//	)
// The result is:
//	uniqueName "some/import/path/1"
//	uniqueName1 "some/import/path/2"
//	...
func (pcs parentControllers) Imports() string {
	results := make([]string, 0, len(pcs))
	for i := range pcs {
		// Ignore repeated packages and the main one (with empty accessor).
		if pcs[i].instance == "" && pcs[i].Accessor != "" {
			results = append(results, fmt.Sprintf(`%s "%s"`, pcs[i].Accessor, pcs[i].Controller.Parents.childImport))
		}
	}
	sort.Strings(results)
	return strings.Join(results, "\n") + "\n"
}

// All returns all parent controllers of a controller including
// grandparents, grandgrandparents, and so forth.
// The result is in the order the controllers must be initialized
// and their special actions must be called.
// I.e. grandparents first, then parents, then children.
func (ps parents) All(pkgs packages, prefix string, ctx *context) (pcs parentControllers) {
	// Iterate over all available parents. Register a new instance of the controller if one doesn't already exist.
	for i := range ps.list {
		// Make sure current parent is a controller rather than
		// some embedded struct.
		pkg, ok := pkgs[ps.list[i].Import] // Make sure it's from the package with controllers.
		if !ok {
			continue // Controllers were not detected in the package of this parent, skip it.
		}
		c := pkg.Controller(ps.list[i].Name) // Make sure it is a controller, not some struct.
		if c == nil {
			continue // This parent is from a package with controllers but not a controller, skip it.
		}

		// Register a new instance of the controller if one doesn't already exist.
		n := fmt.Sprintf("(%s).%s", c.Parents.childImport, c.Name)
		_, ok = ctx.instances[n]
		if !ok {
			ctx.instances[n] = prefix + c.Name // Instance is equal to prefix + current controller name.
		}
	}

	// Iterate over all available parents. Check parents of their parents recursively.
	for i := range ps.list {
		// Make sure current parent is a controller rather than
		// some embedded struct.
		pkg, ok := pkgs[ps.list[i].Import] // Make sure it's from the package with controllers.
		if !ok {
			continue // Controllers were not detected in the package of this parent, skip it.
		}
		c := pkg.Controller(ps.list[i].Name) // Make sure it is a controller, not some struct.
		if c == nil {
			continue // This parent is from a package with controllers but not a controller, skip it.
		}

		// Defining the accessor of the controller package.
		// Local packages must have empty accessors, the same packages must have
		// the same accessors.
		accessor, ok := ctx.packages[c.Parents.childImport] // Getting accessor of the package.
		if !ok {                                            // Accessor hasn't been registered for the package yet.
			accessor = pkg.accessor
			ctx.packages[c.Parents.childImport] = accessor
		}
		if prefix == "" && c.Parents.childImport == ps.childImport {
			accessor = ""
		}

		// Register a new instance of the controller if one doesn't already exist.
		n := fmt.Sprintf("(%s).%s", c.Parents.childImport, c.Name)
		instance, ok := ctx.instances[n]
		if !ok {
			ctx.instances[n] = prefix + c.Name // Instance is equal to prefix + current controller name.
		} else if instance == prefix+c.Name {

			// Add parents' parents to the top of results ("grandparents first" rule).
			// Use old prefix + current controller's name as a new prefix.
			pcs = append(c.Parents.All(pkgs, prefix+c.Name+".", ctx), pcs...)

			instance = ""
		}

		// Add current controller to the bottom.
		pcs = append(pcs, parentController{
			Accessor:   accessor,
			Controller: c,
			Prefix:     prefix,
			instance:   instance,
		})
	}
	return pcs
}

// context is a system type that stores information about parsed packages
// and allocated controllers. It is passed between "parents.All" methods.
type context struct {
	// instances are in the format:
	//	(packageImport).controllerName => ""
	// or:
	//	(packageImport).anotherControllerName => "A.B.C"
	instances map[string]string

	// packages are in the format:
	//	packageImport => "uniqueAccessorX"
	packages map[string]string
}

// newContext allocates and returns a new context.
func newContext() *context {
	return &context{
		instances: map[string]string{},
		packages:  map[string]string{},
	}
}
