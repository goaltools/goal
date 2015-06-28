package handlers

import (
	r "reflect"
	"testing"

	"github.com/anonx/sunplate/log"
	"github.com/anonx/sunplate/reflect"
)

func TestProcessPackage(t *testing.T) {
	ps := packages{}
	ps.processPackage("github.com/anonx/sunplate/generation/handlers/testdata/controllers")
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
	"github.com/anonx/sunplate/generation/handlers/testdata/controllers": controllers{
		"Controller": controller{
			After: &reflect.Func{
				Comments: reflect.Comments{"// After is a magic method that is executed after every request."},
				File:     "init.go",
				Name:     "After",
				Params: reflect.Args{
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
				Results: reflect.Args{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "a",
						},
					},
				},
			},
			Before: &reflect.Func{
				Comments: reflect.Comments{"// Before is a magic method that is executed before every request."},
				File:     "init.go",
				Name:     "Before",
				Params: reflect.Args{
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
				Results: reflect.Args{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "a",
						},
					},
				},
			},
			Finally: &reflect.Func{
				Comments: reflect.Comments{"// Finally is a magic method that is executed after every request."},
				File:     "app.go",
				Name:     "Finally",
				Params: reflect.Args{
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
				Results: reflect.Args{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "action",
						},
					},
				},
			},

			Comments: reflect.Comments{
				"// Controller is a struct that should be embedded into every controller",
				"// of your app to make methods provided by middleware controllers available.",
			},
			File: "init.go",
			Parents: []parent{
				{
					Import: "github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage",
					Name:   "Controller",
				},
			},
		},
		"App": controller{
			Actions: reflect.Funcs{
				{
					Comments: reflect.Comments{"// Index is a sample action."},
					File:     "init.go",
					Name:     "Index",
					Params: reflect.Args{
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
					Results: reflect.Args{
						{
							Type: &reflect.Type{
								Name:    "Result",
								Package: "a",
							},
						},
					},
				},
				{
					Comments: reflect.Comments{"// HelloWorld is a sample action."},
					File:     "app.go",
					Name:     "HelloWorld",
					Params: reflect.Args{
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
					Results: reflect.Args{
						{
							Type: &reflect.Type{
								Name:    "Result",
								Package: "action",
							},
						},
					},
				},
			},
			After:   &reflect.Func{},
			Before:  &reflect.Func{},
			Finally: &reflect.Func{},

			Comments: reflect.Comments{
				"// App is a sample controller.",
			},
			File: "app.go",
			Parents: []parent{
				{
					Import:  "github.com/anonx/sunplate/generation/handlers/testdata/controllers",
					Name:    "Controller",
					Pointer: true,
				},
			},
		},
	},
	"github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage": controllers{
		"Controller": controller{
			Actions: reflect.Funcs{
				{
					Comments: reflect.Comments{"// Index is a sample action."},
					File:     "app.go",
					Name:     "Index",
					Params: reflect.Args{
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
					Results: reflect.Args{
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
				Comments: reflect.Comments{
					"// After is a magic function that is executed after any request.",
				},
				File:   "app.go",
				Name:   "After",
				Params: reflect.Args{},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: reflect.Args{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "action",
						},
					},
				},
			},
			Before: &reflect.Func{
				Comments: reflect.Comments{
					"// Before is a magic function that is executed before any request.",
				},
				File:   "app.go",
				Name:   "Before",
				Params: reflect.Args{},
				Recv: &reflect.Arg{
					Name: "c",
					Type: &reflect.Type{
						Name: "Controller",
						Star: true,
					},
				},
				Results: reflect.Args{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "action",
						},
					},
				},
			},
			Finally: &reflect.Func{
				Comments: reflect.Comments{
					"// Finally is a magic function that is executed after any request",
					"// no matter what.",
				},
				File: "app.go",
				Name: "Finally",
				Params: reflect.Args{
					{
						Name: "userID",
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
				Results: reflect.Args{
					{
						Type: &reflect.Type{
							Name:    "Result",
							Package: "action",
						},
					},
				},
			},

			Comments: reflect.Comments{
				"// Controller is some controller.",
			},
			File:    "app.go",
			Parents: []parent{},
		},
	},
}
