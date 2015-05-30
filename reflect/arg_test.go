package reflect

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestProcessFieldList(t *testing.T) {
	pkg := getTestPackage(t, `package test
			type Sample struct {
				Something            *something.Cool
				FirstName, LastName  *Name
				GPA                  float64
			}
		`,
	)
	expRes := []Arg{
		Arg{
			Name: "Something",
			Type: &Type{
				Name:    "Cool",
				Package: "something",
				Star:    true,
			},
		},
		Arg{
			Name: "FirstName",
			Type: &Type{
				Name: "Name",
				Star: true,
			},
		},
		Arg{
			Name: "LastName",
			Type: &Type{
				Name: "Name",
				Star: true,
			},
		},
		Arg{
			Name: "GPA",
			Type: &Type{
				Name: "float64",
			},
		},
	}
	f := getFields(t, pkg)

	args := processFieldList(f)
	for i, exp := range expRes {
		if args[i].Name != exp.Name || args[i].Tag != exp.Tag ||
			!deepEqualType(args[i].Type, exp.Type) {

			t.Errorf(
				"Cannot process '%s' field. Expected %#v of type %#v, received: %#v of type %#v.",
				exp.Name, exp, exp.Type, args[i], args[i].Type,
			)
		}
	}
}

func TestProcessField(t *testing.T) {
	pkg := getTestPackage(t, `package test
			type Sample struct {
				Something            *something.Cool
				FirstName, LastName  *Name
				GPA                  float64
			}
		`,
	)
	expRes := [][]Arg{
		[]Arg{
			Arg{
				Name: "Something",
				Type: &Type{
					Name:    "Cool",
					Package: "something",
					Star:    true,
				},
			},
		},
		[]Arg{
			Arg{
				Name: "FirstName",
				Type: &Type{
					Name: "Name",
					Star: true,
				},
			},
			Arg{
				Name: "LastName",
				Type: &Type{
					Name: "Name",
					Star: true,
				},
			},
		},
		[]Arg{
			Arg{
				Name: "GPA",
				Type: &Type{
					Name: "float64",
				},
			},
		},
	}

	for i, v := range getFields(t, pkg).List {
		args := processField(v)
		if len(args) == 0 {
			t.Error("Arguments expected, nothing is returned.")
		}
		for j, arg := range args {
			if arg.Name != expRes[i][j].Name || !deepEqualType(arg.Type, expRes[i][j].Type) {
				t.Errorf(
					"Cannot process '%s' field. Expected %#v of type %#v, received: %#v of type %#v.",
					arg.Name, expRes[i][j], expRes[i][j].Type, arg, arg.Type,
				)
			}
		}
	}
}

// deepEqualType is used by tests to ensure two types are equal.
func deepEqualType(t1, t2 *Type) bool {
	if t1.Name == t2.Name && t1.String() == t2.String() && reflect.DeepEqual(t1.Decl, t2.Decl) {
		return true
	}
	return false
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
