package reflect

import (
	"go/ast"
	"testing"
)

func TestProcessFieldList(t *testing.T) {
	pkg := getTestPackage(t, `package test
			type Sample struct {
				Something            *something.Cool
				FirstName, LastName  *Name
				GPA                  int
			}
		`,
	)
	f := getFields(t, pkg)

	res := processFieldList(f)
	if l := len(res); l != 4 {
		t.Errorf("Expected to get 4 elements, got %d instead.", l)
	}

	if res[0].Name != "Something" || res[0].Type.String() != "*something.Cool" {
		t.Errorf("Cannot process 'Something' field, received: %#v.", res[0])
	}

	if res[1].Name != "FirstName" || res[2].Name != "LastName" ||
		res[1].Type.String() != "*Name" || res[2].Type.String() != "*Name" {

		t.Errorf("Cannot process 'FirstName' and  'LastName' fields, received: %#v.", res[1])
	}

	if res[3].Name != "GPA" || res[3].Type.String() != "int" {
		t.Errorf("Cannot process 'GPA' field, received: %#v.", res[2])
	}
}

func TestProcessField(t *testing.T) {
	pkg := getTestPackage(t, `package test
			type Sample struct {
				Something            *something.Cool
				FirstName, LastName  *Name
				GPA                  int
			}
		`,
	)
	f := getFields(t, pkg).List

	args := processField(f[0])
	if len(args) != 1 || args[0].Name != "Something" || args[0].Type.String() != "*something.Cool" {
		t.Errorf("Cannot process 'Something' field, received: %#v.", args)
	}

	args = processField(f[1])
	if len(args) != 2 || args[0].Name != "FirstName" || args[1].Name != "LastName" ||
		args[0].Type.String() != "*Name" || args[1].Type.String() != "*Name" {

		t.Errorf("Cannot process 'FirstName' and  'LastName' fields, received: %#v.", args)
	}

	args = processField(f[2])
	if len(args) != 1 || args[0].Name != "GPA" || args[0].Type.String() != "int" {
		t.Errorf("Cannot process 'GPA' field, received: %#v.", args)
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
