package handlers

import (
	"testing"

	"github.com/anonx/sunplate/reflect"
)

func TestValidArgument(t *testing.T) {
	if ok := !validArgument(&reflect.Arg{Type: &reflect.Type{Name: "App", Package: "xxx"}}); !ok {
		t.Errorf("xxx.App is not a supported type: %v expected, got %v.", ok, !ok)
	}
	if ok := validArgument(&reflect.Arg{Type: &reflect.Type{Name: "int16"}}); !ok {
		t.Errorf("int16 is a supported type: %v expected, got %v.", ok, !ok)
	}
	if ok := !validArgument(&reflect.Arg{Type: &reflect.Type{Name: "[][]int"}}); !ok {
		t.Errorf("[][]int is not a supported type: %v expected, got %v.", ok, !ok)
	}
	if ok := validArgument(&reflect.Arg{Type: &reflect.Type{Name: "[]float64"}}); !ok {
		t.Errorf("[]float64 is a supported type: %v expected, got %v.", ok, !ok)
	}
}

func TestActionFunc(t *testing.T) {
	f := actionFn
	fn := actionFunc(&reflect.Package{
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
		Params: reflect.Args{
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
	res := before(f)
	if res {
		t.Errorf("Incorrect result: action is not a magic Before method.")
	}

	f.Name = "Before"
	res = before(f)
	if !res {
		t.Errorf("Incorrect result: action is a magic Before method.")
	}
}

func TestAfter(t *testing.T) {
	f := actionFn
	res := after(f)
	if res {
		t.Errorf("Incorrect result: action is not a magic After method.")
	}

	f.Name = "After"
	res = after(f)
	if !res {
		t.Errorf("Incorrect result: action is a magic After method.")
	}
}

func TestFinally(t *testing.T) {
	f := actionFn
	res := finally(f)
	if res {
		t.Errorf("Incorrect result: action is not a magic Finally method.")
	}

	f.Name = "Finally"
	res = finally(f)
	if !res {
		t.Errorf("Incorrect result: action is a magic Finally method.")
	}
}

func TestNotMagicAction(t *testing.T) {
	f := actionFn
	f.Name = "Before"
	res := notMagicAction(f)
	if res {
		t.Errorf("Incorrect result: action is a magic method.")
	}

	f.Name = "Index"
	res = notMagicAction(f)
	if !res {
		t.Errorf("Incorrect result: action is not a magic method.")
	}
}

var actionFn = &reflect.Func{
	Comments: reflect.Comments{
		"// Something is a sample action.",
	},
	Name: "Something",
	File: "app.go",
	Params: reflect.Args{
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
	Results: reflect.Args{
		{
			Type: &reflect.Type{
				Name:    "Result",
				Package: "action",
			},
		},
	},
}
