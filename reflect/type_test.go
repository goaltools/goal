package reflect

import (
	"testing"
)

func TestTypeString(t *testing.T) {
	for exp, typ := range testData {
		if got := typ.String(); got != exp {
			t.Errorf("Incorrect output of Type.String() for %#v. Got '%s', expected '%s'.", typ, got, exp)
		}
	}
}

func TestProcessType_IncorrectInput(t *testing.T) {
	typ := processType("String is an incorrect input")
	if typ != nil {
		t.Error("Incorrect input received. *StarExpr, *Ident, or *SelectorExpr of ast were expected.")
	}
}

func TestProcessType(t *testing.T) {
	pkg := getTestPackage(t, `package test
			type Sample struct {
				Something *something.Cool
				Fullname  *Name
				GPA       int
				Grade     grade.Type
			}
		`,
	)
	f := getFields(t, pkg).List

	typ := processType(f[0].Type)
	if typ == nil || typ.Name != "Cool" || typ.Package != "something" || !typ.Star {
		t.Errorf("Field of type *something.Cool expected, got '%v'.", typ)
	}

	typ = processType(f[1].Type)
	if typ == nil || typ.Name != "Name" || typ.Package != "" || !typ.Star {
		t.Errorf("Field of type *Name expected, got '%v'.", typ)
	}

	typ = processType(f[2].Type)
	if typ == nil || typ.Name != "int" || typ.Package != "" || typ.Star {
		t.Errorf("Field of type int expected, got '%v'.", typ)
	}

	typ = processType(f[3].Type)
	if typ == nil || typ.Name != "Type" || typ.Package != "grade" || typ.Star {
		t.Errorf("Field of type grade.Type expected, got '%v'.", typ)
	}
}

var testData = map[string]Type{
	"int64": Type{
		Name: "int64",
	},
	"template.Template": Type{
		Name:    "Template",
		Package: "template",
	},
	"*template.URL": Type{
		Name:    "URL",
		Package: "template",
		Star:    true,
	},
	"*Controller": Type{
		Name: "Controller",
		Star: true,
	},
}
