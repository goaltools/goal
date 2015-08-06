package action

import (
	"testing"

	"github.com/anonx/sunplate/internal/reflect"
)

func TestActionFunc(t *testing.T) {
	f := actionFn
	fn := Func(&reflect.Package{
		Imports: reflect.Imports{
			"app.go": map[string]string{
				"action": "github.com/anonx/sunplate/action",
			},
		},
	})

	res := fn(&reflect.Func{
		Name: "Test",
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

func TestFinally(t *testing.T) {
	f := actionFn
	res := Finally(f)
	if res {
		t.Errorf("Incorrect result: action is not a magic Finally method.")
	}

	f.Name = "Finally"
	res = Finally(f)
	if !res {
		t.Errorf("Incorrect result: action is a magic Finally method.")
	}
}

func TestNotMagicAction(t *testing.T) {
	f := actionFn
	f.Name = "Before"
	res := NotMagicAction(f)
	if res {
		t.Errorf("Incorrect result: action is a magic method.")
	}

	f.Name = "Index"
	res = NotMagicAction(f)
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
				Name:    "Result",
				Package: "action",
			},
		},
	},
}
