package reflect

import (
	"go/ast"
	"go/token"
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
	fset := token.NewFileSet()
	ast.Print(fset, typ)
	switch v := typ.(type) {
	case *ast.StarExpr:
		// X field contains info about an actual type.
		// Try to receive it.
		t := processType(v.X)
		t.Star = true
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
		return &Type{
			Name:    v.Sel.Name,
			Package: t.Name,
		}
	case *ast.ArrayType:
		// Elt contains info about an actual type, get it.
		t := processType(v.Elt)
		t.Name = "[]" + t.Name // Add "[]" to the type name.
		return t
	case *ast.Ellipsis:
		// Elt contains info about an actual type, extract it.
		t := processType(v.Elt)
		t.Name = "..." + t.Name // Add "..." to the type name.
		return t
	}
	return nil
}
