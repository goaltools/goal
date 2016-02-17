package handlers

import (
	"fmt"
	"go/ast"
	r "reflect"

	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/reflect"
)

// field represents a field of a structure that must be automatically binded.
type field struct {
	Name string // Name of the field, e.g. "Request".
	Type string // Type of the binding, e.g. "request" or "action".
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
