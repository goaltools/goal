package action

import (
	"testing"

	"github.com/colegion/goal/internal/reflect"
)

func TestFunc(t *testing.T) {
	f := actionFn
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
		t.Errorf("Incorrect result: actions are methods that return at least one result.")
	}

	res = fn(f)
	if !res {
		t.Errorf("Incorrect result: true should be returned when receiving action as an argument.")
	}

	f.Name = "index"
	res = fn(f)
	if res {
		t.Errorf("Unexported methods cannot be actions.")
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
		t.Errorf("Incorrect result: file without action import cannot contain actions.")
	}

	f.Name = "Initially"
	res = fn(f)
	if res {
		t.Errorf("Use of reserved words for actions should not be allowed. False expected, got true.")
	}

	f1 := actionFn
	f1.Name = "Something"
	f1.Results[0].Type.Name = "NotActionInterface"
	res = fn(f1)
	if res {
		t.Errorf("Not an action interface is returned as a first result: %#v. False expected, got true.", f1)
	}

	f1.Results[0].Type.Name = "Handler"
	f1.File = "app_test.go"
	res = fn(f1)
	if res {
		t.Errorf("Test files must be ignored. False expected, got true.")
	}
}

func TestBuiltin(t *testing.T) {
	f := &reflect.Func{
		Name: "Test",
		Params: []reflect.Arg{
			{
				Name: "name",
				Type: &reflect.Type{
					Name: "string",
				},
			},
		},
	}
	if builtin(f) != true {
		t.Errorf("Parameters of %#v are builtin. True expected, got false.", f)
	}

	f = &reflect.Func{
		Name: "Test",
		Params: []reflect.Arg{
			{
				Name: "name",
				Type: &reflect.Type{
					Name:    "Test",
					Package: "test",
				},
			},
		},
		Recv: &reflect.Arg{
			Type: &reflect.Type{},
		},
	}
	if builtin(f) != false {
		t.Errorf("Parameter `test.Test` of %#v is not builtin. False expected, got true.", f)
	}
}

func TestBefore(t *testing.T) {
	f := actionFn
	res := Before(f)
	if res {
		t.Errorf("Incorrect result: action is not a magic Before method.")
	}

	f.Name = "Before"
	res = Before(f)
	if !res {
		t.Errorf("Incorrect result: action is a magic Before method.")
	}
}

func TestAfter(t *testing.T) {
	f := actionFn
	res := After(f)
	if res {
		t.Errorf("Incorrect result: action is not a magic After method.")
	}

	f.Name = "After"
	res = After(f)
	if !res {
		t.Errorf("Incorrect result: action is a magic After method.")
	}
}

func TestRegular(t *testing.T) {
	f := actionFn
	f.Name = "Before"
	res := Regular(f)
	if res {
		t.Errorf("Incorrect result: action is a magic method.")
	}

	f.Name = "Index"
	res = Regular(f)
	if !res {
		t.Errorf("Incorrect result: action is not a magic method.")
	}
}

var actionFn = &reflect.Func{
	Comments: []string{
		"// Something is a sample action.",
	},
	Name: "Something",
	File: "app.go",
	Params: []reflect.Arg{
		{
			Name: "page",
			Type: &reflect.Type{
				Name: "int",
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
				Name:    "Handler",
				Package: "http",
			},
		},
	},
}
