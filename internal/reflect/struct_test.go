package reflect

import (
	"go/ast"
	"go/token"
	"reflect"
	"strings"
	"testing"

	"github.com/goaltools/goal/internal/log"
)

func TestStructsFilter(t *testing.T) {
	t1 := Structs{
		{
			Name: "Struct1",
		},
		{
			Name: "Struct2",
		},
		{
			Name: "Struct12",
		},
	}
	expRes := Structs{
		{
			Name: "Struct2",
		},
		{
			Name: "Struct12",
		},
	}
	r := t1.Filter(func(s *Struct) bool {
		return true
	})
	assertDeepEqualStructs(t1, r)

	r = t1.Filter(func(s *Struct) bool {
		if strings.HasSuffix(s.Name, "2") {
			return true
		}
		return false
	})
	assertDeepEqualStructs(expRes, r)
}

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
		Comments: Comments{"// Sample ...", "// Line 2"},
		Name:     "Sample",
		Fields: Args{
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

				"github.com/goaltools/goal"
				l "github.com/goaltools/goal/log"
			)
		`,
	)
	expRes := map[string]string{
		"strings": "strings",
		"example": "./example",
		"goal":    "github.com/goaltools/goal",
		"l":       "github.com/goaltools/goal/log",
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
		Fields: Args{
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
				"github.com/goaltools/goal"
				l "github.com/goaltools/goal/log"
				"./example"
				"strings"
			)
		`,
	)
	expRes := map[string]string{
		"goal":    "github.com/goaltools/goal",
		"l":       "github.com/goaltools/goal/log",
		"example": "./example",
		"strings": "strings",
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
	if err := AssertEqualStruct(s1, s2); err != nil {
		log.Error.Panic(err)
	}
}

// assertDeepEqualStructs is a function that is used in tests for
// comparison of structs.
func assertDeepEqualStructs(ss1, ss2 Structs) {
	if err := AssertEqualStructs(ss1, ss2); err != nil {
		log.Error.Panic(err)
	}
}
