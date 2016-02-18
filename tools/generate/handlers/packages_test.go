package handlers

import (
	r "reflect"
	"testing"

	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/reflect"
	"github.com/colegion/goal/internal/routes"
)

func TestProcessPackage(t *testing.T) {
	psR := packages{}
	psR.processPackage("github.com/colegion/goal/tools/generate/handlers/testdata/controllers", routes.Prefixes{
		{
			Method:  "ROUTE",
			Pattern: "",
		},
	})
	assertDeepEqualPkgs(ps, psR)
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
	"github.com/colegion/goal/tools/generate/handlers/testdata/controllers": controllers{
		init: &reflect.Func{
			Comments: []string{"// Init ..."},
			File:     "smth.go",
			Name:     "Init",
			Params: []reflect.Arg{
				{
					Type: &reflect.Type{
						Name:    "Values",
						Package: "url",
					},
				},
			},
		},
		list: []*controller{
			{
				Name: "App",
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
				Parents: parents{
					"github.com/colegion/goal/tools/generate/handlers/testdata/controllers",
					[]parent{
						{0, "github.com/colegion/goal/tools/generate/handlers/testdata/controllers", "Controller"},
						{0, "github.com/colegion/goal/tools/generate/handlers/testdata/controllers", "NotController"},
						{0, "github.com/colegion/goal/tools/generate/handlers/testdata/controllers", "NotController1"},
					},
				},
			},
			{
				Name: "Controller",
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
				Parents: parents{
					"github.com/colegion/goal/tools/generate/handlers/testdata/controllers",
					[]parent{
						{
							Import: "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage",
							Name:   "Controller",
						},
						{
							Import: "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage",
							Name:   "SubSubPackage",
						},
						{
							Import: "github.com/naoina/denco",
							Name:   "Param",
						},
					},
				},
			},
		},
	},
	"github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage": controllers{
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
		list: []*controller{
			{
				Name: "Controller",
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
				Parents: parents{
					"github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage",
					[]parent{
						{
							Import: "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage",
							Name:   "SubSubPackage",
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
	"github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage": controllers{
		list: []*controller{
			{
				Name: "SubSubPackage",
				Before: &reflect.Func{
					Comments: []string{
						"// Before does nothing.",
					},
					File:   "subsubpackage.go",
					Name:   "Before",
					Params: []reflect.Arg{},
					Recv: &reflect.Arg{
						Name: "c",
						Type: &reflect.Type{
							Name: "SubSubPackage",
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
				Parents: parents{
					childImport: "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage",
				},
				Comments: []string{
					"// SubSubPackage is a controller.",
				},
				File: "subsubpackage.go",
			},
		},
	},
}
