// Package main is used as an entry point of
// 'sunplate' toolkit. It validates user input parameters
// and runs subcommands.
package main

import (
	"os"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/create"
	"github.com/anonx/sunplate/generation"
	"github.com/anonx/sunplate/log"
	"github.com/anonx/sunplate/run"
)

// Handlers is a map of registered subcommands
// 'sunplate' toolkit supports.
var Handlers = command.NewContext()

func main() {
	var trace bool

	// Do not show stacktrace if something goes wrong
	// in case tracing is turned off.
	defer func() {
		if !trace {
			if err := recover(); err != nil {
				// Do nothing, error message has already been printed
				// and we do not need stack trace.
			}
		}
	}()

	// Enabling tracing if that is requested by a user.
	command.Helpers["--trace"] = func(val string) {
		if val == "true" {
			trace = true
		}
	}

	// Try to run the subcommand user requested.
	err := Handlers.Process(os.Args[1:]...)
	if err == command.ErrIncorrectArgs { // The arguments were not correct.
		log.Warn.Printf(unknownCmd, os.Args[1])
		return
	}
	if err != nil { // The arguments were omitted.
		Handlers.Process("help", "info") // Show a help message.
		return
	}
}

func init() {
	// Register the supported subcommands.
	Handlers.Register(create.Handler)
	Handlers.Register(run.Handler)
	Handlers.Register(generation.Handler)
	Handlers.Register(helpHandler)

	// Show header message when using new or help
	// commands.
	command.Helpers["new"] = showHeader
	command.Helpers["run"] = showHeader
	command.Helpers["help"] = showHeader
}

// showHeader prints a header message
// with the name of the project.
func showHeader(val string) {
	log.Trace.Println(header)
}

var unknownCmd = `Unknown command "%s".
Run "sunplate help" for usage.`
