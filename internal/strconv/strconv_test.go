package strconv

import (
	"testing"

	r "github.com/anonx/sunplate/internal/reflect"
)

/*
	Below are tests for code generation related functions and methods
	of strconv package.
*/

func TestRender(t *testing.T) {
	c := Context()
	a := r.Arg{Name: "names", Type: &r.Type{Name: "[]string"}}
	exp := `strconv.Strings(r.Form, "names")`
	var expErr error
	if r, err := c.Render("strconv", "r.Form", a); err != expErr || r != exp {
		t.Errorf("Incorrect result of Render. Expected `%v`, `%v`;\ngot `%v`, `%v`.", exp, expErr, r, err)
	}

	a = r.Arg{Name: "names", Type: &r.Type{Name: "YoHoHo"}}
	exp = ""
	expErr = ErrUnsupportedType
	if r, err := c.Render("strconv", "r.Form", a); err != expErr || r != exp {
		t.Errorf("Incorrect result of Render. Expected `%v`, `%v`;\ngot `%v`, `%v`.", exp, expErr, r, err)
	}
}

func TestContext(t *testing.T) {
	c := Context()
	supportedTypes := []string{
		"bool", "string", "int", "int8", "int16", "int32", "int64",
		"float32", "float64", "uint", "uint8", "uint16", "uint32", "uint64",

		"[]bool", "[]string", "[]int", "[]int8", "[]int16", "[]int32", "[]int64",
		"[]float32", "[]float64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64",
	}
	num := len(supportedTypes)
	if l := len(c); l != num {
		t.Errorf("Context returns incorrect number of arguments. Expected %d, got %d.", num, l)
	}
	for _, k := range supportedTypes {
		if _, ok := c[k]; !ok {
			t.Errorf(`Incorrect result of Context. Key "%s" is not found in %v.`, k, c)
		}
	}
}

func TestStrconvFunc(t *testing.T) {
	if ok := !strconvFunc(r.Func{Name: "local"}); !ok {
		t.Errorf(errMsg, ok, !ok)
	}
	if ok := !strconvFunc(r.Func{Name: "E"}); !ok {
		t.Errorf(errMsg, ok, !ok)
	}
	ps := r.Args{
		{
			Name: "vs",
			Type: &r.Type{
				Name:    "Values",
				Package: "url",
			},
		},
		{
			Name: "k",
			Type: &r.Type{
				Name: "string",
			},
		},
		{
			Name: "is",
			Type: &r.Type{
				Name: "...int",
			},
		},
	}
	if ok := !strconvFunc(r.Func{Name: "E", Params: ps}); !ok {
		t.Errorf(errMsg, ok, !ok)
	}
	ps[0].Type = &r.Type{
		Name: "SmthElse",
	}
	if ok := !strconvFunc(r.Func{Name: "E", Params: ps}); !ok {
		t.Errorf(errMsg, ok, !ok)
	}
}

var errMsg = `Incorrect result. Expected "%v", got "%v".`
