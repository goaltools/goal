// Package main is used as an entry point of
// the framework. It validates user input parameters
// and runs subcommands (aka tools).
package main

import (
	"os"
	"strings"

	"github.com/colegion/goal/internal/command"
	"github.com/colegion/goal/log"
)

// handlers is a stores information about the registered subcommands (aka tools)
// the framework supports.
var handlers = command.NewContext()

func main() {
	// Try to run the command user requested.
	// Ignoring the first argument as it is name of the executable.
	err := handlers.Run(os.Args[1:])
	switch err {
	case nil: // Handler's entry function has been successfully executed.
		// Do nothing.
	case command.ErrIncorrectArgs: // Incorrect command requested.
		log.Warn.Printf(unknownCmd, strings.Join(os.Args, " "))
	default: // Some other error has been received.
		log.Error.Println(err)
	}
}

var unknownCmd = `Unknown command "%s".
Run "goal help" for usage.`
