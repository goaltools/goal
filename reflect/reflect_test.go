package reflect

import (
	"testing"

	"go/ast"
	"go/parser"
	"go/token"
)

func TestParseDir(t *testing.T) {
	ParseDir("../example/controllers")
}

func TestProcessStructs(t *testing.T) {
}

// getTestPackage is a function that returns test package file.
func getTestPackage(t *testing.T, src string) *ast.File {
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
