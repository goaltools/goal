package handlers

import (
	r "reflect"
	"testing"

	"github.com/anonx/sunplate/internal/reflect"
	"github.com/anonx/sunplate/log"
)

func TestProcessPackage(t *testing.T) {
	ps := packages{}
	ps.processPackage("github.com/anonx/sunplate/internal/programs/generate/handlers/testdata/controllers")
}

func TestParentPackage(t *testing.T) {
	p := parent{}
	s := p.Package()
	if s != "" {
		// E.g. if we are using it for generation of:
		//	uniquePkgName.Application.Index.
		// In case the Application is local (i.e. its import is empty) we need:
		//	Application.Index.
		// I.e. the method must return empty string.
		t.Errorf("Packages with empty imports must have no names.")
	}
	p = parent{
		ID:     1,
		Import: "net/http",
		Name:   "Request",
	}
	s = p.Package(".XXX")
	if s != "c1.XXX" {
		t.Errorf(`Incorrect package name. Expected "c1.XXX", got "%s".`, s)
	}
}

func TestControllerIgnoredArgs(t *testing.T) {
	c := controller{}
	a := ps["github.com/anonx/sunplate/internal/programs/generate/handlers/testdata/controllers"]["App"].Actions[1]
	exp := ", _, _"
	if r := c.IgnoredArgs(&a); r != exp {
		t.Errorf(`Incorrect IgnoreArgs result. Expected "%s", got "%s".`, exp, r)
	}
}

func assertDeepEqualController(c1, c2 *controller) {
	if c1 == nil || c2 == nil {
		if c1 != c2 {
			log.Error.Panicf(
				"One of the controllers is equal to nil while another is not: %#v != %#v.", c1, c2,
			)
		}
		return
	}
	if c1.File != c2.File {
		log.Error.Panicf("Controllers are from different files: %s != %s.", c1.File, c2.File)
	}
	if !r.DeepEqual(c1.Comments, c2.Comments) {
		log.Error.Panicf("Controllers have different comments: %#v != %#v.", c1.Comments, c2.Comments)
	}
	if !r.DeepEqual(c1.Parents, c2.Parents) {
		log.Error.Panicf("Controllers have different parent controllers: %#v != %#v.", c1.Parents, c2.Parents)
	}
	if err := reflect.AssertEqualFuncs(c1.Actions, c2.Actions); err != nil {
		log.Error.Panic(err)
	}
	if err := reflect.AssertEqualFunc(c1.After, c2.After); err != nil {
		log.Error.Panic(err)
	}
	if err := reflect.AssertEqualFunc(c1.Before, c2.Before); err != nil {
		log.Error.Panic(err)
	}
	if err := reflect.AssertEqualFunc(c1.Finally, c2.Finally); err != nil {
		log.Error.Panic(err)
	}
}

func assertDeepEqualControllers(cs1, cs2 controllers) {
	if len(cs1) != len(cs2) {
		log.Error.Panicf(
			"controllers maps %#v and %#v have different length: %d != %d",
			cs1, cs2, len(cs1), len(cs2),
		)
	}
	for i := range cs1 {
		c1 := cs1[i]
		c2 := cs2[i]
		assertDeepEqualController(&c1, &c2)
	}
}

func assertDeepEqualPkgs(ps1, ps2 packages) {
	if len(ps1) != len(ps2) {
		log.Error.Panicf(
			"packages maps %#v and %#v have different length: %d != %d",
			ps1, ps2, len(ps1), len(ps2),
		)
	}
	for i := range ps1 {
		assertDeepEqualControllers(ps1[i], ps2[i])
	}
}

var ps = packages{
	"github.com/anonx/sunplate/internal/programs/generate/handlers/testdata/controllers": controllers{
		"Controller": controller{
			After: &reflect.Func{
				Comments: []string{"// After is a magic method that is executed after every request."},
				File:     "init.go",
				Name:     "After",
				Params: []reflect.Arg{
					{
						Name: "name",
						Type: &reflect.Type{
							Name: "string",
						},
					},
				},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "a",
						},
					},
				},
			},
			Before: &reflect.Func{
				Comments: []string{"// Before is a magic method that is executed before every request."},
				File:     "init.go",
				Name:     "Before",
				Params: []reflect.Arg{
					{
						Name: "uid",
						Type: &reflect.Type{
							Name: "string",
						},
					},
				},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "a",
						},
					},
				},
			},
			Finally: &reflect.Func{
				Comments: []string{"// Finally is a magic method that is executed after every request."},
				File:     "app.go",
				Name:     "Finally",
				Params: []reflect.Arg{
					{
						Name: "w",
						Type: &reflect.Type{
							Package: "http",
							Name:    "ResponseWriter",
						},
					},
					{
						Name: "r",
						Type: &reflect.Type{
							Package: "http",
							Name:    "Request",
							Star:    true,
						},
					},
				},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name: "bool",
						},
					},
				},
			},
			Initially: &reflect.Func{
				Comments: []string{"// Initially is a magic method that is executed before every request."},
				File:     "app.go",
				Name:     "Initially",
				Params: []reflect.Arg{
					{
						Name: "w",
						Type: &reflect.Type{
							Package: "http",
							Name:    "ResponseWriter",
						},
					},
					{
						Name: "r",
						Type: &reflect.Type{
							Package: "http",
							Name:    "Request",
							Star:    true,
						},
					},
				},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name: "bool",
						},
					},
				},
			},

			Comments: []string{
				"// Controller is a struct that should be embedded into every controller",
				"// of your app to make methods provided by middleware controllers available.",
			},
			File: "init.go",
			Parents: []parent{
				{
					Import: "github.com/anonx/sunplate/internal/programs/generate/handlers/testdata/controllers/subpackage",
					Name:   "Controller",
				},
			},
		},
		"App": controller{
			Actions: []reflect.Func{
				{
					Comments: []string{"// Index is a sample action."},
					File:     "init.go",
					Name:     "Index",
					Params: []reflect.Arg{
						{
							Name: "page",
							Type: &reflect.Type{
								Name: "int",
							},
						},
					},
					Recv: &reflect.Arg{
						Name: "c",
						Type: &reflect.Type{
							Name: "App",
							Star: true,
						},
					},
					Results: []reflect.Arg{
						{
							Type: &reflect.Type{
								Name:    "Result",
								Package: "a",
							},
						},
					},
				},
				{
					Comments: []string{"// HelloWorld is a sample action."},
					File:     "app.go",
					Name:     "HelloWorld",
					Params: []reflect.Arg{
						{
							Name: "page",
							Type: &reflect.Type{
								Name: "int",
							},
						},
					},
					Recv: &reflect.Arg{
						Name: "c",
						Type: &reflect.Type{
							Name: "App",
						},
					},
					Results: []reflect.Arg{
						{
							Type: &reflect.Type{
								Name:    "Result",
								Package: "action",
							},
						},
						{
							Type: &reflect.Type{
								Name: "bool",
							},
						},
						{
							Type: &reflect.Type{
								Name: "error",
							},
						},
					},
				},
			},
			After:   &reflect.Func{},
			Before:  &reflect.Func{},
			Finally: &reflect.Func{},

			Comments: []string{
				"// App is a sample controller.",
			},
			File: "app.go",
			Parents: []parent{
				{
					Import: "github.com/anonx/sunplate/internal/programs/generate/handlers/testdata/controllers",
					Name:   "Controller",
				},
			},
		},
	},
	"github.com/anonx/sunplate/internal/programs/generate/handlers/testdata/controllers/subpackage": controllers{
		"Controller": controller{
			Actions: []reflect.Func{
				{
					Comments: []string{"// Index is a sample action."},
					File:     "app.go",
					Name:     "Index",
					Params: []reflect.Arg{
						{
							Name: "page",
							Type: &reflect.Type{
								Name: "int",
							},
						},
					},
					Recv: &reflect.Arg{
						Name: "c",
						Type: &reflect.Type{
							Name: "App",
							Star: true,
						},
					},
					Results: []reflect.Arg{
						{
							Type: &reflect.Type{
								Name:    "Result",
								Package: "action",
							},
						},
					},
				},
			},
			After: &reflect.Func{
				Comments: []string{
					"// After is a magic function that is executed after any request.",
				},
				File:   "app.go",
				Name:   "After",
				Params: []reflect.Arg{},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "action",
						},
					},
				},
			},
			Before: &reflect.Func{
				Comments: []string{
					"// Before is a magic function that is executed before any request.",
				},
				File:   "app.go",
				Name:   "Before",
				Params: []reflect.Arg{},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "action",
						},
					},
				},
			},
			Finally: &reflect.Func{
				Comments: []string{
					"// Finally is a magic function that is executed after any request",
					"// no matter what.",
				},
				File: "app.go",
				Name: "Finally",
				Params: []reflect.Arg{
					{
						Name: "w",
						Type: &reflect.Type{
							Package: "http",
							Name:    "ResponseWriter",
						},
					},
					{
						Name: "r",
						Type: &reflect.Type{
							Package: "http",
							Name:    "Request",
							Star:    true,
						},
					},
				},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: []reflect.Arg{
					{
						Type: &reflect.Type{
							Name: "bool",
						},
					},
				},
			},

			Comments: []string{
				"// Controller is some controller.",
			},
			File:    "app.go",
			Parents: []parent{},
		},
	},
}
