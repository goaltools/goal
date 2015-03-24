package command

import (
	"testing"
)

func TestNewType_IncorrectArgsNumber(t *testing.T) {
	defer expectPanic("We are using odd number of arguments and thus expecting panic.")
	NewType([]string{"run", "path/to/app", "smth"})
}

func TestNewType(t *testing.T) {
	typ := NewType([]string{"run", "path/to/app"})
	_ = typ
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
