package command

import (
	"testing"
)

func TestNewType_IncorrectArgsNumber(t *testing.T) {
	_, err := NewType([]string{})
	if err != ErrIncorrectArgs {
		t.Error("Parameters' absence is not allowed, error expected.")
	}

	_, err = NewType([]string{"run", "path/to/app", "smth"})
	if err != ErrIncorrectArgs {
		t.Error("Odd number of arguments is not allowed, error expected.")
	}
}

func TestNewType(t *testing.T) {
	typ, err := NewType([]string{"run", "path/to/app", "--smth", "cool"})
	if err != nil {
		t.Errorf("Error was not expected. Got %#v.", err)
	}
	if typ.params["run"] != "path/to/app" {
		t.Error("Arguments were expected to be saved as dictionary.")
	}
	if typ.action != "run" {
		t.Errorf("Action was expected to be 'run'. Instead it is '%s'.", typ.action)
	}
}

func TestRegister_IncorrectArgs(t *testing.T) {
	typ, _ := NewType([]string{"run", "path/to/app"})
	err := typ.Register(map[string]Handler{
		"handlerX": func(action string, params Data) {
			// ToDo
		},
	})
	if err != ErrIncorrectArgs {
		t.Error("Action 'run' is not a registered handler. And thus incorrect args error expected.")
	}
}

func TestRegister(t *testing.T) {
	typ, _ := NewType([]string{"run", "path/to/app"})

	val := ""
	err := typ.Register(map[string]Handler{
		"run": func(action string, params Data) {
			val = params["run"]
		},
	})
	if err != nil {
		t.Errorf("Error expected to be nil. Got %#v.", err)
	}

	if val != "path/to/app" {
		t.Error("Handler function for 'run' was expected to be called. But apparently it wasn't.")
	}
}

func TestDefault(t *testing.T) {
	d := Data{
		"key1": "value1",
	}
	if d.Default("key1", "smth") != "value1" {
		t.Errorf("Default is expected to return value if it exists.")
	}

	if d.Default("key2", "smth") != "smth" {
		t.Errorf("Default is expected to return default value if key is not found.")
	}
}

// expectPanic is used to make sure there was a panic in program.
// If there wasn't, this function panics with the input message.
// Use it as follows:
//	defer expectPanic("We expected panic, but didn't get it")
//	SomeFunctionThatShouldCausePanic()
func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
