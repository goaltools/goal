package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestParseDir_IncorrectPath(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Incorrect path is expected to cause panic, but nothing happened.")
		}
	}()
	ParseDir("testdata/dir_that_does_not_exist")
}

func TestParseDir(t *testing.T) {
	ParseDir("../example/controllers")
}

func TestProcessDecls(t *testing.T) {
	pkg := getPackage(t, `package test
			import (
				"strings"

				"./example"

				"github.com/anonx/sunplate"
				l "github.com/anonx/sunplate/log"
			)

			import "fmt"

			type Sample struct {
				Something string "something"
			}

			func (s *Sample) Test() bool {
				return true
			}

			func init() {
			}
		`,
	)
	expRes := &Package{
		Funcs: []Func{
			{
				File: "sample.go",
				Name: "init",
			},
		},
		Methods: []Func{
			{
				File: "sample.go",
				Name: "Test",
				Recv: &Arg{
					Name: "s",
					Type: &Type{
						Name: "Sample",
						Star: true,
					},
				},
				Results: []Arg{
					{
						Type: &Type{
							Name: "bool",
						},
					},
				},
			},
		},
		Name: "test",
		Structs: []Struct{
			{
				Fields: []Arg{
					{
						Name: "Something",
						Tag:  "something",
						Type: &Type{
							Name: "string",
						},
					},
				},
				File: "sample.go",
				Name: "Sample",
			},
		},
		Imports: map[string]map[string]string{
			"sample.go": map[string]string{
				"strings":  "strings",
				"example":  "./example",
				"sunplate": "github.com/anonx/sunplate",
				"l":        "github.com/anonx/sunplate/log",
				"fmt":      "fmt",
			},
		},
	}

	_ = pkg
	_ = expRes
}

func TestJoinMaps(t *testing.T) {
	a := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	b := map[string]string{
		"key2": "new_value2",
		"key3": "value3",
	}
	expRes := map[string]string{
		"key1": "value1",
		"key2": "new_value2",
		"key3": "value3",
	}
	c := joinMaps(a, b)
	if !reflect.DeepEqual(expRes, c) {
		t.Errorf("Incorrect result of joinMaps. Expected %#v, got %#v.", expRes, c)
	}

	var d map[string]string
	e := joinMaps(d, a)
	if !reflect.DeepEqual(e, a) {
		t.Errorf("Incorrect result of joinMaps if base is nil. Expected %#v, got %#v.", expRes, e)
	}
}

// getPackage is a function that parses input go source and returns ast tree.
func getPackage(t *testing.T, src string) *ast.File {
	fset := token.NewFileSet() // Positions are relative to fset.
	pkg, err := parser.ParseFile(fset, "sample.go", src, parser.ParseComments)
	if err != nil {
		t.Errorf("Failed to parse test package, error: '%s'.", err)
	}
	return pkg
}

// deepEqualPkg is used by tests to compare two packages.
func deepEqualPkg(p1, p2 *Package) bool {
	if p1 == nil || p2 == nil {
		if p1 == p2 {
			return true
		}
		return false
	}
	if p1.Name != p2.Name || !deepEqualStructSlice(p1.Structs, p2.Structs) ||
		!deepEqualFuncSlice(p1.Funcs, p2.Funcs) || !deepEqualFuncSlice(p1.Methods, p2.Methods) {

		return false
	}
	return true
}

var testPackage = `package controllers

import (
	"github.com/anonx/sunplate/action"
	t "github.com/anonx/sunplate/middleware/template"
)

// Application is a test struct.
type Application struct {
	Test struct {
		HelloWorld string "testtag:'helloworld'"
		Smth struct {
			Yahoo string
		}
	}
	Name string "testtag:'name'"
	Age  int    "testtag:'age'"
	*t.Middleware
}

// Index comment is here.
func (c *Application) Index(firstname, lastname string) action.Type {
	return "Incorrect return type but I do not care"
}

// About comment is here.
func (c *Application) About(page int, t1, t2 template.Middleware, smth *template.Middleware) string {
	return "How are ya?"
}
`
