package handlers

import (
	"fmt"

	"github.com/colegion/goal/internal/reflect"
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
	ID     int    // Unique number that is used for generation of import names.
	Import string // Import path of the structure, e.g. "github.com/colegion/goal/template".
	Name   string // Name of the structure, e.g. "Template".
}

// parentControllers represents information necessary for generating a code
// for controller's parents and (grand)parents allocation, their
// special "Before", "After" methods, and "Init" functions calls.
// All elements are in the order they must be allocated/initialized:
// grandparents, parents, children, and so forth.
type parentControllers struct {
	list []*controller // A list of parent controllers.

	// inits is a list of init functions, not including the one
	// of main controllers package.
	inits []reflect.Func
}

// Package returns a unique package name that may be used in templates
// concatenated with some arbitarry suffix strings.
// If a parent is from the same package, empty string will be returned.
// This method is useful to generate things like:
//	import (
//		// {{ .parent.Package .import }} "{{ .parent.Import }}"
//		uniquePkgName "github.com/user/project"
//	)
// and (accessor & dot suffix):
//	// {{ .parent.Package "." }}Application.Index()
//	uniquePkgName.Application.Index()
func (p parent) Package(impPath string, suffixes ...string) string {
	// If the parent is from the same package as the child,
	// we don't need an accessor.
	if p.Import == impPath {
		return ""
	}
	// Otherwise, use package ID and some suffix.
	s := fmt.Sprintf("c%d", p.ID)
	for i := range suffixes {
		s += suffixes[i]
	}
	return s
}

// All returns all parent controllers of a controller including
// grandparents, grandgrandparents, and so forth.
// The result is in the order the controllers must be initialized
// and their special actions must be called.
// I.e. grandparents first, then parents, then children.
func (ps parents) All(pkgs packages) *parentControllers {
	// Allocate a new parent controllers structure.
	pcs := &parentControllers{
		list:  []*controller{},
		inits: []reflect.Func{},
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

		// Add parents' parents and their "Init" functions to the
		// top of results ("grandparents first" rule).
		parents := c.Parents.All(pkgs)
		pcs.list = append(parents.list, pcs.list...)
		pcs.inits = append(parents.inits, pcs.inits...)

		// Add current controller and init function, if presented, to the bottom.
		pcs.list = append(pcs.list, c)
		if pkg.init != nil {
			pcs.inits = append(pcs.inits, *pkg.init)
		}
	}
	return pcs
}
