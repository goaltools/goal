package handlers

import (
	"path/filepath"
	r "reflect"
	"testing"

	"github.com/goaltools/goal/internal/log"
	"github.com/goaltools/goal/internal/reflect"
	"github.com/goaltools/goal/internal/routes"
)

func TestProcessPackage(t *testing.T) {
	psR := packages{}
	psR.processPackage("github.com/goaltools/goal/tools/generate/handlers/testdata/controllers", routes.Prefixes{
		{
			Method:  "ROUTE",
			Pattern: "",
		},
	})
	assertDeepEqualPkgs(ps, psR)
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
	a := ps["github.com/goaltools/goal/tools/generate/handlers/testdata/controllers"].data["App"].Actions[0]
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
	if filepath.Base(c1.File) != filepath.Base(c2.File) {
		log.Error.Panicf("Controllers are from different files: %s != %s.", c1.File, c2.File)
	}
	if !r.DeepEqual(c1.Comments, c2.Comments) {
		log.Error.Panicf("Controllers have different comments: %#v != %#v.", c1.Comments, c2.Comments)
	}
	if !r.DeepEqual(c1.Parents, c2.Parents) {
		log.Error.Panicf("Controllers have different parent controllers: %#v != %#v.", c1.Parents, c2.Parents)
	}
	log.Trace.Println("Actions...")
	if err := reflect.AssertEqualFuncs(c1.Actions, c2.Actions); err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Println("Before...")
	if err := reflect.AssertEqualFunc(c1.Before, c2.Before); err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Println("After...")
	if err := reflect.AssertEqualFunc(c1.After, c2.After); err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Println("Fields...")
	if !r.DeepEqual(c1.Fields, c2.Fields) {
		log.Error.Panicf(`Fields %v and %v are not equal.`, c1.Fields, c2.Fields)
	}
	log.Trace.Println("Routes...")
	assertDeepEqualRoutes(c1.Routes, c2.Routes)
}

func assertDeepEqualRoutes(r1, r2 [][]routes.Route) {
	if len(r1) != len(r2) {
		log.Error.Panicf(`Routes %v and %v are of different lengths: %d != %d.`, r1, r2, len(r1), len(r2))
	}
	for i := range r1 {
		if !r.DeepEqual(r1[i], r2[i]) {
			log.Error.Panicf(`Routes of %dth action are different: %v and %v.`, i, r1, r2)
		}
	}
}

func assertDeepEqualControllers(cs1, cs2 controllers) {
	if len(cs1.data) != len(cs2.data) {
		log.Error.Panicf(
			"controllers maps %#v and %#v have different length: %d != %d",
			cs1.data, cs2.data, len(cs1.data), len(cs2.data),
		)
	}
	if err := reflect.AssertEqualFunc(cs1.init, cs2.init); err != nil {
		log.Error.Panic(err)
	}
	for k := range cs1.data {
		c1 := cs1.data[k]
		c2 := cs2.data[k]
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
	for k := range ps1 {
		assertDeepEqualControllers(ps1[k], ps2[k])
	}
}

var ps = packages{
	"github.com/goaltools/goal/tools/generate/handlers/testdata/controllers": controllers{
		data: map[string]controller{
			"App": {
				Actions: []reflect.Func{
					{
						Comments: []string{
							"// HelloWorld is a sample action.", "//@get",
							"// Below is an unsupported method.", "//@someMethodThatDoesntExist /hello/world",
						},
						File: "app.go",
						Name: "HelloWorld",
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
									Name:    "Handler",
									Package: "http",
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
									Name:    "Handler",
									Package: "h",
								},
							},
						},
					},
				},

				Routes: [][]routes.Route{
					{
						{Method: "GET", Pattern: "/App/HelloWorld", HandlerName: "App.HelloWorld"},
					},
				},
				Comments: []string{
					"// App is a sample controller.",
				},
				File: "app.go",
				Parents: []parent{
					{
						Name: "Controller",
					},
					{
						Name: "NotController",
					},
					{
						Name: "NotController1",
					},
				},
			},
			"Controller": {
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
								Name:    "Handler",
								Package: "h",
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
								Name:    "Handler",
								Package: "h",
							},
						},
					},
				},
				Fields: []field{
					{
						Name: "R",
						Type: "request",
					},
					{
						Name: "W",
						Type: "response",
					},
					{
						Name: "A",
						Type: "action",
					},
					{
						Name: "C",
						Type: "controller",
					},
				},

				Comments: []string{
					"// Controller is a struct that should be embedded into every controller",
					"// of your app to make methods provided by middleware controllers available.",
				},
				File: "init.go",
				Parents: []parent{
					{
						Import: "github.com/goaltools/goal/tools/generate/handlers/testdata/controllers/subpackage",
						Name:   "Controller",
					},
					{
						Import: "github.com/naoina/denco",
						Name:   "Param",
					},
				},
			},
		},
	},
	"github.com/goaltools/goal/tools/generate/handlers/testdata/controllers/subpackage": controllers{
		init: &reflect.Func{
			Comments: []string{"// Init ..."},
			File:     "app.go",
			Name:     "Init",
			Params: []reflect.Arg{
				{
					Name: "ctx",
					Type: &reflect.Type{
						Name:    "Values",
						Package: "url",
					},
				},
			},
		},
		data: map[string]controller{
			"Controller": {
				Actions: []reflect.Func{
					{
						Comments: []string{"// Index is a sample action.", "//@post index someindexlabel"},
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
								Name: "Controller",
							},
						},
						Results: []reflect.Arg{
							{
								Type: &reflect.Type{
									Name:    "Handler",
									Package: "http",
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
								Name:    "Handler",
								Package: "http",
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
								Name:    "Handler",
								Package: "http",
							},
						},
					},
				},

				Routes: [][]routes.Route{
					{
						{
							Method:      "POST",
							Pattern:     "/subpackage/index",
							Label:       "someindexlabel",
							HandlerName: "Controller.Index",
						},
					},
				},
				Comments: []string{
					"// Controller is some controller.",
				},
				File: "app.go",
			},
		},
	},
}
