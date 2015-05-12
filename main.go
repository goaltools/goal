// Package main is used as an entry point of
// 'sunplate' toolkit. It validates user input parameters
// and runs subcommands.
package main

import (
	"go/build"
	"os"
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation"
	"github.com/anonx/sunplate/help"
	"github.com/anonx/sunplate/log"
)

// Handlers is a map of registered commands
// 'sunplate' toolkit supports.
var Handlers = map[string]command.Handler{
	"generate": generation.Start,
	"help":     help.Start,
}

func main() {
	// Do not show stacktrace if something goes wrong.
	defer catchPanic()

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
		log.Warn.Printf("Unknown command '%s'.\nRun 'sunplate help' for usage.", os.Args[1])
	}
}

func catchPanic() {
	if err := recover(); err != nil {
		log.Error.Fatal(err)
	}
}

func init() {
	// Set a base path of 'generation' package's files.
	generation.BasePath = filepath.Join(
		build.Default.GOPATH, "src/github.com/anonx/sunplate/generation",
	)
}

var header = `~
~ https://github.com/anonx/sunplate
~
                        _       _
                       | |     | |
  ___ _   _ _ __  _ __ | | __ _| |_ ___
 / __| | | | '_ \| '_ \| |/ _' | __/ _ \
 \__ \ |_| | | | | |_) | | (_| | ||  __/
 |___/\__,_|_| |_| .__/|_|\__,_|\__\___|
                 | |
                 |_|
`
