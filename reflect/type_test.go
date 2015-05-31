package reflect

import (
	"testing"
)

func TestTypeString(t *testing.T) {
	expRes := map[string]Type{
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
	for exp, typ := range expRes {
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
	pkg := getPackage(t, `package test
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
	`)
	expRes := []Type{
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
	}

	for i, v := range getFields(t, pkg).List {
		if len(expRes) > i { // Anonymous struct should be skipped.
			break
		}
		typ := processType(v.Type)
		if typ == nil || typ.Name != expRes[i].Name ||
			typ.Package != expRes[i].Package || typ.Star != expRes[i].Star {

			t.Errorf("Field of type %#v expected, got '%#v'.", expRes[i], typ)
		}
	}
}

// deepEqualType is a function that is used by tests to compare types.
func deepEqualType(t1, t2 *Type) bool {
	if t1 == nil || t2 == nil {
		if t1 == t2 {
			return true
		}
		return false
	}
	return t1.String() == t2.String()
}
