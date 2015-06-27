// Package generation is a runner of subcommands
// related to code generation.
package generation

import (
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/handlers"
	"github.com/anonx/sunplate/generation/listing"
	"github.com/anonx/sunplate/log"
)

// BasePath is a path from directory where executable file
// is located to this file.
var BasePath = "./"

// Start is an entry point of generate command.
func Start(action string, params command.Data) {
	switch params[action] {
	case "handler", "handlers":
		handlers.Start(filepath.Join(BasePath, "./handlers"), params)
	case "listing":
		listing.Start(filepath.Join(BasePath, "./listing"), params)
	default:
		log.Trace.Panicf(errMsg, params[action], action)
	}
}

var errMsg = `I do not know how to generate "%s".
Run "sunplate help %s" for more information.
`
