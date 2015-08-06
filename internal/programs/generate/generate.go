// Package generate is a runner of subcommands
// related to code generation.
package generate

import (
	"github.com/anonx/sunplate/internal/command"
	"github.com/anonx/sunplate/internal/programs/generate/handlers"
	"github.com/anonx/sunplate/internal/programs/generate/views"
	"github.com/anonx/sunplate/log"
)

// Handler is an instance of generate subcommand.
var Handler = command.Handler{
	Name:  "generate",
	Info:  `automatic generation of golang code (for use with "go generate")`,
	Usage: "generate {command} [arguments]",
	Desc: `The commands are:
	handlers    generate golang handler functions for a controllers package
	views       generate a list of file paths found in a requested directory

Possible arguments include:
	--input     a path to directory that should be scanned
	--output    a directory where to save generated .go files
	--package   name of the package that should be created

All paths are expected to be relative to the root of your project.

If no arguments are received the default values will be used:
	generate handlers --input ./controllers --output ./assets/handlers --package handlers
	generate listing --input ./views --output ./assets/views --package views
`,

	Main: start,
}

// start is an entry point of generate command.
func start(action string, params command.Data) {
	switch params[action] {
	case "handler", "handlers":
		handlers.Start(params)
	case "view", "views":
		views.Start(params)
	default:
		log.Error.Panicf(errMsg, params[action], action)
	}
}

var errMsg = `I do not know how to generate "%s".
Run "sunplate help %s" for more information.
`
