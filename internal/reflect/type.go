package reflect

import (
	"fmt"
	"go/ast"
)

// Type represents a type of argument.
type Type struct {
	Name    string // Name of the type, e.g. "URL". It is empty if Decl is not nil.
	Package string // Package name, e.g. "template" in case of "html/template".
	Star    bool   // Star indicates whether it is a pointer.
}

// String prints a type name, e.g. "*template.URL", "template.Template",
// "Controller", "int64", etc.
func (t *Type) String() (name string) {
	name = t.Name
	if t.Package != "" {
		name = t.Package + "." + name
	}
	if t.Star {
		name = "*" + name
	}
	return name
}

// processType parses go ast tree related to types into
// Type, a format that is used by this reflect package.
func processType(typ interface{}) *Type {
	switch v := typ.(type) {
	case *ast.StarExpr:
		// X field contains info about an actual type.
		// Try to receive it.
		t := processType(v.X)
		if t != nil {
			t.Star = true
		}
		return t
	case *ast.Ident:
		// Initialize a name of the type and return it.
		return &Type{
			Name: v.Name,
		}
	case *ast.SelectorExpr:
		// X contains info about a selector's package.
		// Try to extract it.
		t := processType(v.X)
		if t != nil {
			t = &Type{
				Name:    v.Sel.Name,
				Package: t.Name,
			}
		}
		return t
	case *ast.ArrayType:
		// Elt contains info about an actual type, get it.
		t := processType(v.Elt)
		if t != nil {
			t.Name = fmt.Sprintf("[]%s", t.Name) // Add "[]" to the type name.
		}
		return t
	case *ast.MapType:
		// Extract key and value's types.
		t := &Type{}
		kt := processType(v.Key)
		vt := processType(v.Value)
		if kt != nil && vt != nil {
			t.Name = fmt.Sprintf("map[%s]%s", kt.String(), vt.String())
		}
		return t
	case *ast.Ellipsis:
		// Elt contains info about an actual type, extract it.
		t := processType(v.Elt)
		if t != nil {
			t.Name = fmt.Sprintf("...%s", t.Name) // Add "..." to the type name.
		}
		return t
	default: // Ignore unsupported types.
		// TODO: add a better support of map and support of anonymous structs.
	}
	return nil
}
