package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"

	"github.com/anonx/sunplate/log"
)

func TestImportsValue(t *testing.T) {
	var t1 Imports
	v, ok := t1.Value("somefile.go", "somename")
	if v != "" || ok != false {
		t.Errorf("Incorrect imports value. Expected '', false. Got '%s', %v.", v, ok)
	}

	t2 := Imports{
		"sample.go": {
			"l": "github.com/anonx/sunplate/log",
		},
	}
	v, ok = t2.Value("sample.go", "l")
	if v != "github.com/anonx/sunplate/log" || ok != true {
		t.Errorf("Incorrect imports value. Expected 'github.com/anonx/sunplate/log', true. Got '%s', %v.", v, ok)
	}

	v, ok = t2.Value("sample.go", "key_that_does_not_exist")
	if v != "" || ok != false {
		t.Errorf("Incorrect imports value. Expected '', false. Got '%s', %v.", v, ok)
	}
}

func TestImportsName(t *testing.T) {
	var t1 Imports
	v, ok := t1.Name("somefile.go", "github.com/anonx/sunplate")
	if v != "" || ok != false {
		t.Errorf("Incorrect import name value. Expected '', false. Got '%s', %v.", v, ok)
	}

	t2 := Imports{
		"sample.go": {
			"action":  "github.com/anonx/sunplate/action",
			"example": "github.com/anonx/sunplate/example",
			"l":       "github.com/anonx/sunplate/log",
		},
	}
	v, ok = t2.Name("sample.go", "sunplate/example")
	if v != "example" || ok != true {
		t.Errorf("Incorrect import name value. Expected 'example', true. Got '%s', %v.", v, ok)
	}

	v, ok = t2.Name("sample.go", "github.com/import_that_does_not_exist")
	if v != "" || ok != false {
		t.Errorf("Incorrect imports value. Expected '', false. Got '%s', %v.", v, ok)
	}
}

func TestParseDir_IncorrectPath(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Incorrect path is expected to cause panic, but nothing happened.")
		}
	}()
	ParseDir("testdata/dir_that_does_not_exist", false)
}

func TestParseDir(t *testing.T) {
	p := ParseDir("./testdata", false)
	expRes := &Package{
		Funcs: Funcs{
			{
				File: "testdata/sample2.go",
				Name: "init",
			},
		},
		Methods: Methods{
			"Test": Funcs{
				{
					Comments: Comments{"// Hello is a method."},
					File:     "testdata/sample1.go",
					Name:     "Hello",
					Recv: &Arg{
						Name: "t",
						Type: &Type{
							Name: "Test",
						},
					},
					Params: Args{
						{
							Name: "names",
							Type: &Type{
								Name: "[]string",
							},
						},
						{
							Name: "args",
							Type: &Type{
								Name: "...int",
							},
						},
					},
					Results: Args{
						{
							Type: &Type{
								Name: "string",
							},
						},
					},
				},
			},
		},
		Name: "sample",
		Structs: Structs{
			{
				Comments: Comments{"// Test is a type."},
				Fields: Args{
					{
						Name: "Name",
						Tag:  `tag:"name"`,
						Type: &Type{
							Name: "string",
						},
					},
				},
				File: "testdata/sample1.go",
				Name: "Test",
			},
		},

		Imports: map[string]map[string]string{
			"testdata/sample1.go": {
				"fmt":     "fmt",
				"l":       "github.com/anonx/sunplate/log",
				"reflect": "github.com/anonx/sunplate/internal/reflect",
			},
			"testdata/sample2.go": {
				"log": "github.com/anonx/sunplate/log",
			},
		},
	}

	assertDeepEqualPkg(expRes, p)
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
		Funcs: Funcs{
			{
				File: "sample.go",
				Name: "init",
			},
		},
		Methods: Methods{
			"Sample": Funcs{
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
					Results: Args{
						{
							Type: &Type{
								Name: "bool",
							},
						},
					},
				},
			},
		},
		Name: "test",
		Structs: Structs{
			{
				Fields: Args{
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
			"sample.go": {
				"strings":  "strings",
				"example":  "./example",
				"sunplate": "github.com/anonx/sunplate",
				"l":        "github.com/anonx/sunplate/log",
				"fmt":      "fmt",
			},
		},
	}
	fs, ms, ss, is := processDecls(pkg.Decls, "sample.go")
	if !reflect.DeepEqual(expRes.Imports["sample.go"], is) {
		t.Errorf("Incorrect imports returned. Expected %#v, got %#v.", expRes.Imports, is)
	}
	assertDeepEqualFuncs(expRes.Funcs, fs)
	assertDeepEqualMethods(expRes.Methods, ms)
	assertDeepEqualStructs(expRes.Structs, ss)
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

// assertDeepEqualMethods is used by tests to compare two Methods values.
func assertDeepEqualMethods(ms1, ms2 Methods) {
	if err := AssertEqualMethods(ms1, ms2); err != nil {
		log.Error.Panic(err)
	}
}

// assertDeepEqualPkg is used by tests to compare two packages.
func assertDeepEqualPkg(p1, p2 *Package) {
	if err := AssertEqualPkg(p1, p2); err != nil {
		log.Error.Panic(err)
	}
}
