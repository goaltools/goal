package reflect

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestProcessFuncDecl(t *testing.T) {
	pkg := getPackage(t, `package test
			// Index is a method of App.
			// Index does something cool.
			func (c *App) Index() action.Result {
				return c.RenderTemplate("index.html")
			}

			// Some comment here...

			// About is used for...
			// Try it.
			func (c App) About(page int) (res *action.Result) {
				return c.RenderTemplate("about.html")
			}

			// HelloWorld is a demo function.
			// It expects a greeting message, name, and your age.
			func HelloWorld(greeting, name string, age int) (string, bool) {
				return fmt.Sprintf(
					"%s, %s. You're %d y.o.", greeting, name, age,
				), true
			}

			func init() {
				// TODO: add something cool...
			}
		`,
	)
	expRes := []Func{
		{
			Comments: []string{"// Index is a method of App.", "// Index does something cool."},
			Name:     "Index",
			Results: []Arg{
				{
					Type: &Type{
						Name:    "Result",
						Package: "action",
					},
				},
			},
			Recv: &Arg{
				Name: "c",
				Type: &Type{
					Name: "App",
					Star: true,
				},
			},
		},
		{
			Comments: []string{"// About is used for...", "// Try it."},
			Name:     "About",
			Params: []Arg{
				{
					Name: "page",
					Type: &Type{
						Name: "int",
					},
				},
			},
			Results: []Arg{
				{
					Name: "res",
					Type: &Type{
						Name:    "Result",
						Package: "action",
						Star:    true,
					},
				},
			},
			Recv: &Arg{
				Name: "c",
				Type: &Type{
					Name: "App",
				},
			},
		},
		{
			Comments: []string{"// HelloWorld is a demo function.", "// It expects a greeting message, name, and your age."},
			Name:     "HelloWorld",
			Params: []Arg{
				{
					Name: "greeting",
					Type: &Type{
						Name: "string",
					},
				},
				{
					Name: "name",
					Type: &Type{
						Name: "string",
					},
				},
				{
					Name: "age",
					Type: &Type{
						Name: "int",
					},
				},
			},
			Results: []Arg{
				{
					Type: &Type{
						Name: "string",
					},
				},
				{
					Type: &Type{
						Name: "bool",
					},
				},
			},
		},
		{
			Name: "init",
		},
	}
	for i, decl := range pkg.Decls {
		funcDecl := decl.(*ast.FuncDecl)
		f := processFuncDecl(funcDecl)
		if !deepEqualFunc(f, &expRes[i]) {
			t.Errorf("Incorrect func value. Expected %#v, got %#v.", expRes[i], f)
		}
	}
}

// deepEqualFunc is used by tests to check whether two Func structs are
// equal or not.
func deepEqualFunc(f1, f2 *Func) bool {
	if f1 == nil || f2 == nil {
		if f1 == f2 {
			return true
		}
		return false
	}
	if f1.Name != f2.Name || f1.File != f2.File ||
		!reflect.DeepEqual(f1.Comments, f2.Comments) || !deepEqualArg(f1.Recv, f2.Recv) ||
		len(f1.Params) != len(f2.Params) || len(f1.Results) != len(f2.Results) {

		return false
	}
	if !deepEqualArgSlice(f1.Params, f2.Params) || !deepEqualArgSlice(f1.Results, f2.Results) {
		return false
	}
	return true
}

// deepEqualFuncSlice is a function that is used in tests for
// comparison of func slices.
func deepEqualFuncSlice(f1, f2 []Func) bool {
	if len(f1) != len(f2) {
		return false
	}
	for i, fn := range f1 {
		if !deepEqualFunc(&fn, &f2[i]) {
			return false
		}
	}
	return true
}
