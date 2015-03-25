package main

import (
	"os"

	"github.com/anonx/ok/command"
	"github.com/anonx/ok/help"
	"github.com/anonx/ok/log"
)

// Handlers is a list of registered commands
// 'ok' toolkit supports.
var Handlers = map[string]command.Handler{
	"help": help.Start,
}

func main() {
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
