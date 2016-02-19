package handlers

import (
	"fmt"
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

// All returns all parent controllers of a controller including
// grandparents, grandgrandparents, and so forth.
// The result is in the order the controllers must be initialized
// and their special actions must be called.
// I.e. grandparents first, then parents, then children.
func (ps parents) All(pkgs packages, prefix string, ctx *context) (pcs parentControllers) {
	// Calculate the current level of embedding
	// That is equal to the number of dots in prefix (e.g. in "Child.Parent.").
	level := strings.Count(prefix, ".")

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

		// Add parents' parents to the top of results ("grandparents first" rule).
		// Use old prefix + current controller's name as a new prefix.
		pcs = append(c.Parents.All(pkgs, prefix+c.Name+".", ctx), pcs...)

		// Defining the accessor of the controller package.
		// Local packages must have empty accessors, the same packages must have
		// the same accessors.
		accessor, ok := ctx.packages[c.Parents.childImport] // Getting accessor of the package.
		if !ok {                                            // Accessor hasn't been registered for the package yet.
			ctx.packages[c.Parents.childImport] = fmt.Sprintf("c%dx%d", level, i)
		}
		if level == 0 && c.Parents.childImport == ps.childImport {
			accessor = ""
		}

		// Register a new instance of the controller if one doesn't already exist.
		n := fmt.Sprintf("(%s).%s", c.Parents.childImport, c.Name)
		instance, ok := ctx.instances[n]
		if !ok {
			ctx.instances[n] = prefix + c.Name // Instance is equal to prefix + current controller name.
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
