// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
)

// Start is an entry point of the generate handlers command.
func Start(basePath string, params command.Data) {
	// Generate and save a new package.
	t := output.NewType(
		params.Default("--package", "handlers"), filepath.Join(basePath, "./handlers.go.template"),
	)
	t.CreateDir(params.Default("--output", "./assets/handlers/"))
	t.Extension = ".go" // Save generated file as a .go source.
	t.Context = map[string]interface{}{
		"rootPath": params.Default("--path", "./controllers/"),
	}
	t.Generate()
}
