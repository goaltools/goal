package method

import (
	"testing"

	"github.com/colegion/goal/internal/reflect"
)

func TestFunc(t *testing.T) {
	f := magicM
	fn := Func(&reflect.Package{
		Imports: reflect.Imports{
			"app.go": map[string]string{
				"http": "net/http",
			},
			"init.go": map[string]string{},
		},
	})

	res := fn(&reflect.Func{
		Name: "Test",
		File: "app.go",
	})
	if res {
		t.Errorf("Incorrect result: Magic methods must return 2 arguments.")
	}

	res = fn(f)
	if !res {
		t.Errorf(
			"True expected: (%v) %s(%v, %v) %s.", f.Recv, f.Name, f.Params[0], f.Params[1], f.Results[0].Type,
		)
	}

	f.Name = "lowercase"
	res = fn(f)
	if res {
		t.Errorf("Incorrect result: Magic methods must be exported.")
	}

	fn(&reflect.Func{ // Trying a file without action imports for the first time.
		Name: "Test",
		File: "init.go",
	})
	res = fn(&reflect.Func{ // The second try must be ignored immidiately.
		Name: "Test",
		File: "init.go",
	})
	if res {
		t.Errorf("Incorrect result: file without http import cannot contain magic methods.")
	}

	f.Name = "Something"
	f.Params[0].Type.Name = "NotResponseWriter"
	res = fn(f)
	if res {
		t.Errorf("Incorrect first argument, false expected: %#v.", f.Params[0].Type)
	}

	f1 := magicM
	f1.Params[0].Type.Name = "ResponseWriter"
	f1.Results[0].Type.Name = "notBool"
	res = fn(f1)
	if res {
		t.Errorf("Not a bool is returned as a result: %#v. False expected, got true.", f1)
	}
	f1.Results[0].Type.Name = "bool"

	f1.File = "app_test.go"
	res = fn(f1)
	if res {
		t.Error("Test files must be ignored. False expected, got true.")
	}
	f1.File = "app.go"
}

func TestInitially(t *testing.T) {
	f := magicM
	f.Name = "XXX"
	res := Initially(f)
	if res {
		t.Errorf("Incorrect result: function is not a magic Initially method.")
	}

	f.Name = "Initially"
	res = Initially(f)
	if !res {
		t.Errorf("Incorrect result: function is a magic Initially method.")
	}
}

func TestFinally(t *testing.T) {
	f := magicM
	f.Name = "XXX"
	res := Finally(f)
	if res {
		t.Errorf("Incorrect result: function is not a magic Finally method.")
	}

	f.Name = "Finally"
	res = Finally(f)
	if !res {
		t.Errorf("Incorrect result: function is a magic Finally method.")
	}
}

var magicM = &reflect.Func{
	Comments: []string{
		"// Something is a magic method.",
	},
	Name: "Something",
	File: "app.go",
	Params: []reflect.Arg{
		{
			Name: "w",
			Type: &reflect.Type{
				Package: "http",
				Name:    "ResponseWriter",
			},
		},
		{
			Name: "r",
			Type: &reflect.Type{
				Package: "http",
				Name:    "Request",
				Star:    true,
			},
		},
	},
	Recv: &reflect.Arg{
		Name: "c",
		Type: &reflect.Type{
			Name: "App",
			Star: true,
		},
	},
	Results: []reflect.Arg{
		{
			Type: &reflect.Type{
				Name: "bool",
			},
		},
	},
}
