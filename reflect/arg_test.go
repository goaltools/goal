package reflect

import (
	"go/ast"
	"testing"

	"github.com/anonx/sunplate/log"
)

func TestProcessFieldList_EmptyInput(t *testing.T) {
	args := processFieldList(nil)
	if len(args) != 0 {
		t.Errorf("Empty arguments list expected, got %#v instead.", args)
	}
}

func TestProcessFieldList(t *testing.T) {
	pkg := getPackage(t, `package test
			type Sample struct {
				Something            *something.Cool
				FirstName, LastName  *Name
				GPA                  float64
			}
		`,
	)
	expRes := []Arg{
		{
			Name: "Something",
			Type: &Type{
				Name:    "Cool",
				Package: "something",
				Star:    true,
			},
		},
		{
			Name: "FirstName",
			Type: &Type{
				Name: "Name",
				Star: true,
			},
		},
		{
			Name: "LastName",
			Type: &Type{
				Name: "Name",
				Star: true,
			},
		},
		{
			Name: "GPA",
			Type: &Type{
				Name: "float64",
			},
		},
	}
	f := getFields(t, pkg)

	args := processFieldList(f)
	for i, exp := range expRes {
		assertDeepEqualArg(&exp, &args[i])
	}
}

func TestProcessField_UnknownType(t *testing.T) {
	pkg := getPackage(t, `package test
			type Sample struct {
				Smth struct {
					Field1 string
					Field2 string
				}
			}
		`,
	)
	f := getFields(t, pkg).List
	args := processField(f[0])
	if len(args) != 0 {
		t.Errorf("Fields of anonymous struct type should be skipped. Instead received %#v.", args)
	}
}

func TestProcessField_EmptyName(t *testing.T) {
	pkg := getPackage(t, `package test
			func Test() string {
				return ""
			}
		`,
	)
	expRes := Arg{
		Type: &Type{
			Name: "string",
		},
	}
	funcDecl := pkg.Decls[0].(*ast.FuncDecl)
	l := processFieldList(funcDecl.Type.Results)
	assertDeepEqualArg(&expRes, &l[0])
}

func TestProcessField(t *testing.T) {
	pkg := getPackage(t, `package test
			type Sample struct {
				Something            *something.Cool "tag:something"
				FirstName, LastName  *Name
				GPA                  float64
			}
		`,
	)
	expRes := [][]Arg{
		{
			{
				Name: "Something",
				Tag:  "tag:something",
				Type: &Type{
					Name:    "Cool",
					Package: "something",
					Star:    true,
				},
			},
		},
		{
			{
				Name: "FirstName",
				Type: &Type{
					Name: "Name",
					Star: true,
				},
			},
			{
				Name: "LastName",
				Type: &Type{
					Name: "Name",
					Star: true,
				},
			},
		},
		{
			{
				Name: "GPA",
				Type: &Type{
					Name: "float64",
				},
			},
		},
	}

	for i, v := range getFields(t, pkg).List {
		args := processField(v)
		assertDeepEqualArgSlice(expRes[i], args)
	}
}

// getFields is a test function that receives test package file and
// returns a list of fields of the first struct having been found there.
func getFields(t *testing.T, pkg *ast.File) *ast.FieldList {
	decl, ok := pkg.Decls[0].(*ast.GenDecl)
	if !ok {
		t.Error("Incorrect test package, cannot find general declaration.")
	}

	spec, ok := decl.Specs[0].(*ast.TypeSpec)
	if !ok {
		t.Error("Incorrect test package, cannot find spec.")
	}

	s, ok := spec.Type.(*ast.StructType)
	if !ok {
		t.Error("Incorrect test package, cannot find a struct.")
	}

	return s.Fields
}

// assertDeepEqualArg is a function that is used by tests to compare two arguments.
func assertDeepEqualArg(a1, a2 *Arg) {
	if a1 == nil || a2 == nil {
		if a1 != a2 {
			log.Error.Panicf("One of the arguments is nil, while other is not: %#v != %#v", a1, a2)
		}
		return
	}
	assertDeepEqualType(a1.Type, a2.Type)
	if a1.Name != a2.Name || a1.Tag != a2.Tag {
		log.Error.Panicf("Arguments are not equal: %#v != %#v.", a1, a2)
	}
}

// assertDeepEqualArgSlice is a function that is used in tests for
// comparison of struct slices.
func assertDeepEqualArgSlice(a1, a2 []Arg) {
	if len(a1) != len(a2) {
		log.Error.Panicf(
			"Argument slices %#v and %#v have different length: %d and %d.",
			a1, a2, len(a1), len(a2),
		)
		return
	}
	for i, a := range a1 {
		assertDeepEqualArg(&a, &a2[i])
	}
}
