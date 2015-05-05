// Package main is used as an entry point of
// 'ok' toolkit. It validates user input parameters
// and runs subcommands.
package main

import (
	"os"

	"github.com/anonx/ok/command"
	"github.com/anonx/ok/generation"
	"github.com/anonx/ok/help"
	"github.com/anonx/ok/log"
	"github.com/anonx/ok/scan"
)

// Handlers is a map of registered commands
// 'ok' toolkit supports.
var Handlers = map[string]command.Handler{
	"generate": generation.Start,
	"help":     help.Start,
	"scan":     scan.Start,
}

func main() {
	// Show header message.
	log.Trace.Println(header)

	// Validate input parameters and find out what user wants to run.
	ct, err := command.NewType(os.Args[1:])
	if err != nil {
		// Validation failed because of incorrect arguments number,
		// try to show help menu instead.
		ct, err = command.NewType([]string{"help", "info"})
	}
	err = ct.Register(Handlers)
	if err != nil {
		// Validation failed because requested handler does not exist.
		log.Warn.Printf("Unknown command '%s'.\nRun 'ok help' for usage.", os.Args[1])
	}
}

var header = `~
~ https://github.com/anonx/ok
~
    /\  \         /\__\
   /::\  \       /:/  /
  /:/\:\  \     /:/__/
 /:/  \:\  \   /::\__\____
/:/__/ \:\__\ /:/\:::::\__\
\:\  \ /:/  / \/_|:|~~|~
 \:\  /:/  /     |:|  |
  \:\/:/  /      |:|  |
   \::/  /       |:|  |
    \/__/         \|__|
`
