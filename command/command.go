// Package command is used for parsing input parameters.
package command

import (
	"errors"
)

// ErrIncorrectArgs is returned every time a user is trying
// to use input parameters we do not expect.
var ErrIncorrectArgs = errors.New("incorrect arguments received")

// Handler is an entry function of subprograms.
// It expects the command name as a first argument and a map of
// all available parameters as a second one.
type Handler func(string, map[string]string)

// Data is an internal type for representation of user input parameters.
type Data map[string]string

// Type is a main type of command package.
// It is used for storage of parsed parameters.
type Type struct {
	action string
	params Data
}

// NewType initializes and returns Type object. It expects even number of args.
// Otherwise, it an IncorIncorrectArgsErr error will be returned.
func NewType(args []string) (*Type, error) {
	// Make sure the number of arguments is even number
	// and it is more than zero.
	if len(args) == 0 || len(args)%2 != 0 {
		return nil, ErrIncorrectArgs
	}

	// Save the arguments as a dict.
	params := Data{}
	for i := 0; i < len(args); i += 2 {
		params[args[i]] = args[i+1]
	}

	// The first argument is a requested action.
	return &Type{
		action: args[0],
		params: params,
	}, nil
}

// Register gets a list of handlers and tries to call that one
// which was requested by the user.
// It returns ErrIncorrectArgs error if handler does not exist.
func (t *Type) Register(handlers map[string]Handler) error {
	handler, ok := handlers[t.action]
	if !ok {
		return ErrIncorrectArgs
	}
	handler(t.action, t.params)
	return nil
}

// Default expects a key and a value as input parameters.
// If such key exists within params, an associated value is returned.
// Otherwise, the value received as input parameter is returned.
func (t Data) Default(key, value string) string {
	if v, ok := t[key]; ok {
		return v
	}
	return value
}
