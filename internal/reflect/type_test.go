package reflect

import (
	"testing"

	"github.com/colegion/goal/log"
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
			Numbers []int
			IDs []info.ID
			Contact struct {
				Email string "email"
				Phone int64
			}
		}
	`)
	expRes := []*Type{
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
		{
			Name: "[]int",
		},
		{
			Name:    "[]ID",
			Package: "info",
		},
		nil,
	}

	for i, v := range getFields(t, pkg).List {
		typ := processType(v.Type)
		if typ == nil && expRes[i] == nil {
			continue
		}
		if typ == nil || typ.Name != expRes[i].Name ||
			typ.Package != expRes[i].Package || typ.Star != expRes[i].Star {

			t.Errorf("Field of type %#v expected, got '%#v'.", expRes[i], typ)
		}
	}
}

func assertDeepEqualType(t1, t2 *Type) {
	if err := AssertEqualType(t1, t2); err != nil {
		log.Error.Panic(err)
	}
}
