package reflect

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestProcessStructTypeSpec(t *testing.T) {
	pkg := getPackage(t, `package test
			type Sample struct {
				Something *something.Cool "something"

				Name struct {
					First, Last string
				}
			}
		`,
	)
	expRes := &Struct{
		Fields: []Arg{
			{
				Name: "Something",
				Tag:  "\"something\"",
				Type: &Type{
					Name:    "Cool",
					Package: "something",
					Star:    true,
				},
			},
		},
		Name: "Sample",
	}
	genDecl, _ := pkg.Decls[0].(*ast.GenDecl)
	typeSpec, _ := genDecl.Specs[0].(*ast.TypeSpec)
	res := processStructTypeSpec(typeSpec)
	if !deepEqualStruct(expRes, res) {
		fset := token.NewFileSet()
		ast.Print(fset, typeSpec)
		t.Errorf("Incorrect processStructTypeSpec result. Expected %#v, got %#v.", expRes, res)
	}
}

// deepEqualStruct is used by tests to compare two structures.
func deepEqualStruct(s1, s2 *Struct) bool {
	if s1 == nil || s2 == nil {
		if s1 == s2 {
			return true
		}
		return false
	}
	if !reflect.DeepEqual(s1.Comments, s2.Comments) || s1.Name != s2.Name {
		return false
	}
	if len(s1.Fields) != len(s2.Fields) {
		return false
	}
	for i, field := range s1.Fields {
		if !deepEqualArg(&field, &s2.Fields[i]) {
			return false
		}
	}
	return true
}
