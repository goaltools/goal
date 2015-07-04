package reflect

import (
	"go/ast"
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
	r, count := t1.FilterGroups(func(f *Func) bool {
		return true
	}, func(f *Func) bool {
		return true
	})
	if count != 3 {
		t.Errorf("Func Filter: incorrect count result. Expected %d, got %d.", 3, count)
	}
	assertDeepEqualFuncs(t1, r[0])

	r, count = t1.FilterGroups(func(f *Func) bool {
		if strings.HasSuffix(f.Name, "2") {
			return true
		}
		return false
	}, func(f *Func) bool {
		return true
	})
	if count != 2 {
		t.Errorf("Func Filter: incorrect count result. Expected %d, got %d.", 2, count)
	}
	assertDeepEqualFuncs(expRes, r[0])
}

func TestProcessFuncDecl(t *testing.T) {
	pkg := getPackage(t, `package test
			// Index is a method of App.
			// Index does something cool.
			func (c *App) Index(smth map[string]string) action.Result {
				return c.RenderTemplate("index.html")
			}

			// Some comment here...

			// About is used for...
			// Try it.
			func (c App) About(pages ...int) (res *action.Result) {
				return c.RenderTemplate("about.html")
			}

			// HelloWorld is a demo function.
			// It expects a greeting message, name, and your age.
			func HelloWorld(greeting, name string, ages []int) (string, bool) {
				return fmt.Sprintf(
					"%s, %s. You're %d y.o.", greeting, name, age[0],
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
			Params: Args{
				{
					Name: "smth",
					Type: &Type{
						Name: "map[string]string",
					},
				},
			},
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
					Name: "pages",
					Type: &Type{
						Name: "...int",
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
					Name: "ages",
					Type: &Type{
						Name: "[]int",
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
	if err := AssertEqualFunc(f1, f2); err != nil {
		log.Error.Panic(err)
	}
}

// assertDeepEqualFuncs is a function that is used in tests for
// comparison of functions.
func assertDeepEqualFuncs(fs1, fs2 Funcs) {
	if err := AssertEqualFuncs(fs1, fs2); err != nil {
		log.Error.Panic(err)
	}
}
