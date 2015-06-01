package reflect

import (
	"go/ast"
	"testing"
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
		if !deepEqualArg(&exp, &args[i]) {
			t.Errorf(
				"Cannot process '%s' field. Expected %#v of type %#v, received: %#v of type %#v.",
				exp.Name, exp, exp.Type, args[i], args[i].Type,
			)
		}
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
	if !deepEqualArg(&expRes, &l[0]) {
		t.Errorf("Incorrect fieldList result. Expected field %#v, got %#v.", expRes, l[0])
	}
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

// deepEqualArg is a function that is used by tests to compare two arguments.
func deepEqualArg(a1, a2 *Arg) bool {
	if a1 == nil || a2 == nil {
		if a1 == a2 {
			return true
		}
		return false
	}
	if a1.Name == a2.Name && a1.Tag == a2.Tag && deepEqualType(a1.Type, a2.Type) {
		return true
	}
	return false
}
