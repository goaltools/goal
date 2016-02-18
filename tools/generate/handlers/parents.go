package handlers

import (
	"fmt"
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
func (ps parents) All(pkgs packages) (pcs []*controller) {
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
		pcs = append(c.Parents.All(pkgs), pcs...)

		// Add current controller to the bottom.
		pcs = append(pcs, c)
	}
	return pcs
}
