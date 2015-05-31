package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestParseDir(t *testing.T) {
	ParseDir("../example/controllers")
}

func TestProcessStructs(t *testing.T) {
	pkg := getPackage(t, testPackage)
	fset := token.NewFileSet() // Positions are relative to fset.
	ast.Print(fset, pkg)
}

func TestProcessCommentGroup_EmptyGroup(t *testing.T) {
	if c := processCommentGroup(nil); len(c) != 0 {
		t.Errorf("Zero length slice expected. Gor %#v instead.", c)
	}
}

func TestProcessCommentGroup(t *testing.T) {
	c := processCommentGroup(&ast.CommentGroup{
		List: []*ast.Comment{
			&ast.Comment{
				Text: "// This is line 1",
			},
			&ast.Comment{
				Text: "// This is line 2",
			},
			&ast.Comment{
				Text: "// This is line 3",
			},
		},
	})
	expRes := []string{
		"// This is line 1", "// This is line 2", "// This is line 3",
	}
	if !reflect.DeepEqual(c, expRes) {
		t.Errorf("Incorrect result of processCommentGroup. Expected %#v, got %#v.", expRes, c)
	}
}

// getPackage is a function that parses input go source and returns ast tree.
func getPackage(t *testing.T, src string) *ast.File {
	fset := token.NewFileSet() // Positions are relative to fset.
	pkg, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		t.Errorf("Failed to parse test package, error: '%s'.", err)
	}
	return pkg
}

var testPackage = `package controllers

import (
	"github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/middleware/template"
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
	*template.Middleware
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
