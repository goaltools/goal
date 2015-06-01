package reflect

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestProcessStructDecl_IncorrectTok(t *testing.T) {
	s := processStructDecl(&ast.GenDecl{
		Tok: token.IMPORT,
	})
	if s != nil {
		t.Errorf("Nil should be returned in case genDecl's Tok != token.TYPE, %#v is returned.", s)
	}
}

func TestProcessStructDecl_EmptySpec(t *testing.T) {
	s := processStructDecl(&ast.GenDecl{
		Tok: token.TYPE,
	})
	if s != nil {
		t.Errorf("In case of empty Specs, nil is expected. Got %#v.", s)
	}
}

func TestProcessStructDecl(t *testing.T) {
	pkg := getPackage(t, `package test
			// Sample ...
			// Line 2
			type Sample struct {
				Something *something.Cool "something"

				Name struct {
					First, Last string
				}
			}
		`,
	)
	expRes := Struct{
		Comments: []string{"// Sample ...", "// Line 2"},
		Name:     "Sample",
		Fields: []Arg{
			{
				Name: "Something",
				Tag:  "something",
				Type: &Type{
					Name:    "Cool",
					Package: "something",
					Star:    true,
				},
			},
		},
	}
	genDecl, _ := pkg.Decls[0].(*ast.GenDecl)
	r := processStructDecl(genDecl)
	if !deepEqualStruct(&expRes, r) {
		t.Errorf("Error: %#v != %#v.", expRes, r)
	}
}

func TestProcessTypeSpec_IncorrectType(t *testing.T) {
	s := processTypeSpec(&ast.TypeSpec{
		Type: &ast.InterfaceType{},
	})
	if s != nil {
		t.Errorf("StructType is the only supported type and thus nil expected, got %#v.", s)
	}
}

func TestProcessTypeSpec(t *testing.T) {
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
				Tag:  "something",
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
	res := processTypeSpec(typeSpec)
	if !deepEqualStruct(expRes, res) {
		fset := token.NewFileSet()
		ast.Print(fset, typeSpec)
		t.Errorf("Incorrect processStructTypeSpec result. Expected %#v, got %#v.", expRes, res)
	}
}

func TestProcessImportSpec(t *testing.T) {
	pkg := getPackage(t, `package test
			import(
				"github.com/anonx/sunplate"
				l "github.com/anonx/sunplate/log"
				"./example"
				"strings"
			)
		`,
	)
	expRes := map[string]string{
		"sunplate": "github.com/anonx/sunplate",
		"l":        "github.com/anonx/sunplate/log",
		"example":  "./example",
		"strings":  "strings",
	}
	genDecl, _ := pkg.Decls[0].(*ast.GenDecl)
	for _, st := range genDecl.Specs { // Iterating over specs.
		k, v := processImportSpec(st.(*ast.ImportSpec))
		if expRes[k] != v {
			t.Errorf(
				"Incorrect import key-value pair. Expected %s=%s, got %s=%s.",
				k, expRes[k], k, v,
			)
		}
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
