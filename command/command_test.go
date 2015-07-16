package command

import (
	"testing"
)

func TestProcess_IncorrectArgs(t *testing.T) {
	c := NewContext()
	c.Register(Handler{
		Name: "new",
	})

	err := c.Process()
	if err == nil || err == ErrIncorrectArgs {
		t.Errorf(
			"If no arguments are received an error different from IncorrectArgs is expected. Got error: %v.",
			err,
		)
	}

	err = c.Process("new", "path/to/app", "smth")
	if err != nil {
		t.Errorf("Odd number of arguments is allowed, got error: %v", err)
	}

	err = c.Process("run", "path/to/app")
	if err != ErrIncorrectArgs {
		t.Errorf("Unregistered handler requested. Error expected, got %v.", err)
	}
}

func TestRegister(t *testing.T) {
	c := NewContext()
	c.Register(Handler{
		Name: "new",
	})
	if _, ok := (*c)["new"]; !ok {
		t.Errorf("The first handler was not registered. Context is %v.", c)
	}
	c.Register(Handler{
		Name: "create",
	})
	if _, ok := (*c)["create"]; !ok {
		t.Errorf("The second handler was not registered. Context is %v.", c)
	}
}

func TestProcess(t *testing.T) {
	a := ""
	v := ""
	y := false

	Helpers["-x"] = func(val string) {
		if val == "y" {
			y = true
		}
	}

	c := NewContext()
	c.Register(Handler{
		Name: "update",
		Main: func(action string, params Data) {
			if action != "update" {
				t.Errorf(`The first argument should be a subcommand name, "%s" is not.`, action)
			}
			a = action
			v = params[action]
		},
	})

	err := c.Process("update", "test", "-x", "y")
	if err != nil {
		t.Errorf("Correct input arguments are used, but got error: %v.", err)
	}
	if y != true {
		t.Errorf("Helper function was not started. Y is still %v.", y)
	}

	if a != "update" || v != "test" {
		t.Errorf(`Looks like entry function was not started. As a="%s", v="%s".`, a, v)
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
