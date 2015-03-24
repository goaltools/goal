package command

import (
	"testing"
)

func TestNewType_IncorrectArgsNumber(t *testing.T) {
	_, err := NewType([]string{})
	if err != IncorrectArgsErr {
		t.Error("Parameters' absence is not allowed, error expected.")
	}

	_, err = NewType([]string{"run", "path/to/app", "smth"})
	if err != IncorrectArgsErr {
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
