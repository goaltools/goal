// Package generation is a runner of subcommands
// related to code generation.
package generation

import (
	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/handlers"
	"github.com/anonx/sunplate/generation/listing"
	"github.com/anonx/sunplate/log"
)

// Handler is an instance of generate subcommand.
var Handler = command.Handler{
	Name: "generate",

	Main: start,
}

// start is an entry point of generate command.
func start(action string, params command.Data) {
	switch params[action] {
	case "handler", "handlers":
		handlers.Start(params)
	case "listing":
		listing.Start(params)
	default:
		log.Trace.Panicf(errMsg, params[action], action)
	}
}

var errMsg = `I do not know how to generate "%s".
Run "sunplate help %s" for more information.
`
