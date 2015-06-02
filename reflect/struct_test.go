package reflect

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"

	"github.com/anonx/sunplate/log"
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
	assertDeepEqualStruct(&expRes, r)
}

func TestProcessImportDecl_IncorrectTok(t *testing.T) {
	s := processImportDecl(&ast.GenDecl{
		Tok: token.TYPE,
	})
	if s != nil {
		t.Errorf("Nil should be returned in case genDecl's Tok != token.IMPORT, %#v is returned.", s)
	}
}

func TestProcessImportDecl_EmptySpec(t *testing.T) {
	s := processImportDecl(&ast.GenDecl{
		Tok: token.IMPORT,
	})
	if s == nil || len(s) > 0 {
		t.Errorf("In case of empty Specs, initialized but empty map is expected. Got %#v.", s)
	}
}

func TestProcessImportDecl(t *testing.T) {
	pkg := getPackage(t, `package test
			import (
				"strings"

				"./example"

				"github.com/anonx/sunplate"
				l "github.com/anonx/sunplate/log"
			)
		`,
	)
	expRes := map[string]string{
		"strings":  "strings",
		"example":  "./example",
		"sunplate": "github.com/anonx/sunplate",
		"l":        "github.com/anonx/sunplate/log",
	}
	genDecl, _ := pkg.Decls[0].(*ast.GenDecl)
	imps := processImportDecl(genDecl)
	if !reflect.DeepEqual(expRes, imps) {
		t.Errorf("Incorrect result returned. Expected %#v, got %#v.", expRes, imps)
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
	assertDeepEqualStruct(expRes, res)
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

// assertDeepEqualStruct is used by tests to compare two structures.
func assertDeepEqualStruct(s1, s2 *Struct) {
	if s1 == nil || s2 == nil {
		if s1 != s2 {
			log.Error.Panicf("One of the structs is nil while another is not: %#v != %#v.", s1, s2)
		}
		return
	}
	assertDeepEqualArgSlice(s1.Fields, s2.Fields)
	if !reflect.DeepEqual(s1.Comments, s2.Comments) {
		log.Error.Panicf("Comments of structs are not equal: %#v != %#v.", s1.Comments, s2.Comments)
	}
	if s1.Name != s2.Name || s1.File != s2.File {
		log.Error.Panicf("Structs are not equal: %#v != %#v.", s1, s2)
	}
}

// assertDeepEqualStructSlice is a function that is used in tests for
// comparison of struct slices.
func assertDeepEqualStructSlice(s1, s2 []Struct) {
	if len(s1) != len(s2) {
		log.Error.Panicf(
			"Struct slices %#v and %#v have different length: %d and %d.",
			s1, s2, len(s1), len(s2),
		)
		return
	}
	for i, st := range s1 {
		assertDeepEqualStruct(&st, &s2[i])
	}
}
