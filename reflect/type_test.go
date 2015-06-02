package reflect

import (
	"testing"

	"github.com/anonx/sunplate/log"
)

func TestTypeString(t *testing.T) {
	expRes := map[string]Type{
		"int64": {
			Name: "int64",
		},
		"template.Template": {
			Name:    "Template",
			Package: "template",
		},
		"*template.URL": {
			Name:    "URL",
			Package: "template",
			Star:    true,
		},
		"*Controller": {
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
		{
			Name:    "Cool",
			Package: "something",
			Star:    true,
		},
		{
			Name: "Name",
			Star: true,
		},
		{
			Name: "float64",
		},
		{
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

// assertDeepEqualType is a function that is used by tests to compare types.
func assertDeepEqualType(t1, t2 *Type) {
	if t1 == nil || t2 == nil {
		if t1 != t2 {
			log.Error.Panicf("One of the types is nil while another is not: %#v != %#v.", t1, t2)
		}
		return
	}
	if t1.String() != t2.String() {
		log.Error.Panicf("Types are not equal: %s != %s.", t1, t2)
	}
}
