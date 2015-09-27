// Package command is used for parsing input parameters,
// and starting subcommands (aka tools).
package command

import (
	"errors"
	"flag"
	"strings"
)

const commandWordSep = " "

// Context stores information about available subcommands (aka tools)
// and other related things that are necessary for their start.
type Context struct {
	list []Handler // A list of registered subcommands (tools).
	def  *Handler  // Command that will be started if no arguments received.
}

// NewContext gets a number of handlers as arguments, allocates
// a new Context and returns it.
func NewContext(handlers ...Handler) *Context {
	// Allocate a new context with handlers
	// as a list.
	c := &Context{
		list: handlers,
	}

	// Find default handler and add it to the context.
	for i := 0; i < len(handlers); i++ {
		if handlers[i].Default {
			c.def = &handlers[i]
			return c
		}
	}
	return c
}

// Handler is a type for representation of a subcommand (aka tool)
// such as "cli run" or "cli new".
type Handler struct {
	// Run is an entry function of the handler.
	// The args are the arguments after the command name.
	Run func(h *Handler, args []string) error

	// Default means the handler must be executed if no arguments are
	// received from user (in addition to when it is called explicitly).
	// Only first default handler is used, others will be ignored.
	Default bool

	Name  string // Name of the handler, e.g. "new" or "generate stuff".
	Usage string // Possible arguments of the command, e.g. "[input] [output]".
	Info  string // One line description of the command.
	Desc  string // Detailed description of what the command does.

	Flags flag.FlagSet // Set of flags specific to the command.
}

// ErrIncorrectArgs is returned every time a user is trying
// to request a command that does not exist.
var ErrIncorrectArgs = errors.New("incorrect command requested")

// Run gets a list of arguments and either starts an entry function of the
// requested subcommand (aka tool) or returns an error.
func (c *Context) Run(args []string) error {
	// Start default handlers if no arguments are received.
	if len(args) == 0 {
		return c.def.Run(c.def, args)
	}

	// Otherwise, iterating over all available handlers of subcommands (aka tools).
	for i := 0; i < len(c.list); i++ {
		// Check whether current handler belongs to the subcommand (tool)
		// that is requested by user.
		if lst, ok := c.list[i].requested(args); ok {
			// Start the entry function of the handler.
			return c.list[i].Run(&c.list[i], lst)
		}
	}
	return ErrIncorrectArgs
}

// requested checks whether the handler is the one that is requested by user,
// i.e. handler's name or alias is a part of args.
// It returns arguments (not including the handler name) and true in case
// of success, and nil, false otherwise.
func (h Handler) requested(args []string) ([]string, bool) {
	// Calculate the number of words in the handler's name.
	// It is equal to the number of spaces plus one.
	num := strings.Count(h.Name, commandWordSep) + 1

	// If the number of arguments is less than the number of words
	// in handler's name that means this is not the command user wants.
	if len(args) < num {
		return nil, false
	}

	// Make sure the handler's name is equal to the one user requested.
	if h.Name != strings.Join(args[:num], commandWordSep) {
		return nil, false
	}

	return args[num:], true
}
