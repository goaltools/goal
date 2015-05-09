// Package generation is a runner of subcommands
// related to code generation.
package generation

import (
	"github.com/anonx/sunplate/generation/handlers"
	"github.com/anonx/sunplate/generation/listing"
	"github.com/anonx/sunplate/log"
)

// Start is an entry point of generate command.
func Start(action string, params map[string]string) {
	switch params[action] {
	case "handler", "handlers":
		handlers.Start(params)
	case "listing":
		listing.Start(params)
	default:
		log.Warn.Printf(
			"I do not know how to generate '%s'.\nRun 'sunplate help %s' for more information.",
			params[action], action,
		)
	}
}
