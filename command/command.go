// Package command is used for parsing input parameters.
package command

import (
	"errors"
)

// Context stores information about supported subcommands.
// Data has the following format:
//	Command name:
//		Handler of the command
type Context map[string]Handler

// Handler is a type for representation of a subprogram.
// It contains an entry function of the subprogram,
// and its info, description, help messages, etc.
type Handler struct {
	Name  string // Name of the handler, e.g. "new".
	Usage string // How to use the subprogram, e.g. "new [path]".
	Info  string // One line description of the command.
	Desc  string // Detailed description of what the command does.

	Main HandlerFunc // Entry function of the handler.
}

// HandlerFunc is an entry function of a subprogram.
// It expects the command name as a first argument and a map of
// all available parameters as a second one.
type HandlerFunc func(string, Data)

// Data is an internal type for representation of user input parameters.
type Data map[string]string

// errIncorrectArgs is returned every time a user is trying
// to use input parameters we do not expect.
var errIncorrectArgs = errors.New("incorrect arguments received")

// NewContext allocates and returns a new instance of Context.
func NewContext() *Context {
	return &Context{}
}

// Register gets a handler and adds it to the list of supported ones.
func (c Context) Register(h Handler) {
	c[h.Name] = h
}

// Process gets a number of arguments, validates, and either
// starts executing a requested handler or returns an error. The first
// argument has a special meaning, it is a handler's name.
func (c Context) Process(args ...string) error {
	// Make sure the number of arguments is even number
	// and it is more than zero.
	if len(args) == 0 || len(args)%2 != 0 {
		return errIncorrectArgs
	}

	// Save the arguments as a dict.
	params := Data{}
	for i := 0; i < len(args); i += 2 {
		params[args[i]] = args[i+1]
	}

	// Check whether requested subcommand exists.
	// First argument is its name.
	if h, ok := c[args[0]]; ok {
		// Call the handler's entry function.
		h.Main(args[0], params)
		return nil
	}

	// Otherwise, return Incorrect Arguments error.
	return errIncorrectArgs
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

// Type is a main type of command package.
// It is used for storage of parsed parameters.
type Type struct {
	action string
	params Data
}
