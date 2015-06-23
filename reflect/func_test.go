package reflect

import (
	"go/ast"
	"reflect"
	"strings"
	"testing"

	"github.com/anonx/sunplate/log"
)

func TestFuncsFilter(t *testing.T) {
	t1 := Funcs{
		{
			Name: "Func1",
		},
		{
			Name: "Func2",
		},
		{
			Name: "Func32",
		},
	}
	expRes := Funcs{
		{
			Name: "Func2",
		},
		{
			Name: "Func32",
		},
	}
	r := t1.Filter(func(f *Func) bool {
		return true
	})
	assertDeepEqualFuncs(t1, r[0])

	r = t1.Filter(func(f *Func) bool {
		if strings.HasSuffix(f.Name, "2") {
			return true
		}
		return false
	}, func(f *Func) bool {
		return true
	})
	assertDeepEqualFuncs(expRes, r[0])
	assertDeepEqualFuncs(t1, r[1])
}

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
	expRes := Funcs{
		{
			Comments: Comments{"// Index is a method of App.", "// Index does something cool."},
			Name:     "Index",
			Results: Args{
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
			Comments: Comments{"// About is used for...", "// Try it."},
			Name:     "About",
			Params: Args{
				{
					Name: "page",
					Type: &Type{
						Name: "int",
					},
				},
			},
			Results: Args{
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
			Comments: Comments{"// HelloWorld is a demo function.", "// It expects a greeting message, name, and your age."},
			Name:     "HelloWorld",
			Params: Args{
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
			Results: Args{
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
		assertDeepEqualFunc(f, &expRes[i])
	}
}

// assertDeepEqualFunc is used by tests to check whether two Func structs are
// equal or not.
func assertDeepEqualFunc(f1, f2 *Func) {
	if f1 == nil || f2 == nil {
		if f1 != f2 {
			log.Error.Panicf("One of the funcs is nil while another is not: %#v != %#v.", f1, f2)
		}
		return
	}
	assertDeepEqualArg(f1.Recv, f2.Recv)
	assertDeepEqualArgs(f1.Params, f2.Params)
	assertDeepEqualArgs(f1.Results, f2.Results)
	if !reflect.DeepEqual(f1.Comments, f2.Comments) {
		log.Error.Panicf("Comments of funcs are not equal: %#v != %#v.", f1.Comments, f2.Comments)
	}
	if f1.Name != f2.Name || f1.File != f2.File {
		log.Error.Panicf("Funcs are not equal: %#v != %#v.", f1, f2)
	}
}

// assertDeepEqualFuncs is a function that is used in tests for
// comparison of functions.
func assertDeepEqualFuncs(f1, f2 Funcs) {
	if len(f1) != len(f2) {
		log.Error.Panicf(
			"Func slices %#v and %#v have different length: %d and %d.",
			f1, f2, len(f1), len(f2),
		)
		return
	}
	for i, fn := range f1 {
		assertDeepEqualFunc(&fn, &f2[i])
	}
}
