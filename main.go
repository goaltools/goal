// Package main is used as an entry point of
// the framework. It validates user input parameters
// and runs subcommands (aka tools).
package main

import (
	"flag"
	"os"

	"github.com/colegion/goal/internal/command"
	"github.com/colegion/goal/log"
	"github.com/colegion/goal/tools/create"
	"github.com/colegion/goal/tools/run"
)

var trace = flag.Bool("trace", false, "Show stack trace in case of runtime errors.")

// handlers is a stores information about the registered subcommands (aka tools)
// the framework supports.
var handlers = command.NewContext(
	create.Handler,
	run.Handler,
)

func main() {
	// Do not show stacktrace if something goes wrong
	// but tracing is disabled.
	defer func() {
		if err := recover(); err != nil {
			if *trace {
				log.Warn.Fatalf("TRACE: %v.", err)
			}
		}
	}()

	// Printing a header message with a name of the framework.
	log.Trace.Println(header)

	// Try to run the command user requested.
	// Ignoring the first argument as it is name of the executable.
	flag.Parse()
	err := handlers.Run(flag.Args())
	if err != nil {
		log.Warn.Printf(unknownCmd, err, os.Args[0])
	}
}

var header = `~
~ Goal Framework
~ https://github.com/colegion/goal
~
  ██████╗  ██████╗  █████╗ ██╗
 ██╔════╝ ██╔═══██╗██╔══██╗██║
 ██║  ███╗██║   ██║███████║██║
 ██║   ██║██║   ██║██╔══██║██║
 ╚██████╔╝╚██████╔╝██║  ██║███████╗
  ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝
`

var unknownCmd = `Error: %v.
Run "%s help" for usage.`
