package reflect

import (
	"reflect"
	"testing"
)

func TestTypeString(t *testing.T) {
	expectedResults := map[string]Type{
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
	for exp, typ := range expectedResults {
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
	testPackage := `package test
		type Sample struct {
			Something *something.Cool
			Fullname  *Name
			GPA float64
			Grade grade.Type
			Contact struct {
				Email string "email"
				Phone int64
			}
		}
	`
	expectedResults := []Type{
		Type{
			Name:    "Cool",
			Package: "something",
			Star:    true,
		},
		Type{
			Name: "Name",
			Star: true,
		},
		Type{
			Name: "float64",
		},
		Type{
			Name:    "Type",
			Package: "grade",
		},
		Type{
			Decl: &Struct{
				Fields: []Arg{
					Arg{Name: "Email", Type: &Type{Name: "string"}, Tag: "\"email\""},
					Arg{Name: "Phone", Type: &Type{Name: "int64"}},
				},
			},
		},
	}
	pkg := getTestPackage(t, testPackage)
	f := getFields(t, pkg).List

	for i, v := range f {
		typ := processType(v.Type)
		if typ == nil || !reflect.DeepEqual(typ.Decl, expectedResults[i].Decl) ||
			typ.Name != expectedResults[i].Name || typ.Package != expectedResults[i].Package ||
			typ.Star != expectedResults[i].Star {

			t.Errorf("Field of type %#v expected, got '%#v'.", expectedResults[i], typ)
		}
	}
}
