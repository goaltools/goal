// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/reflect"
)

const (
	// ActionInterfaceImport is a GOPATH to the Result interface
	// that should be implemented by types returned by actions.
	ActionInterfaceImport = "github.com/anonx/sunplate/action"

	// ActionInterfaceName is an interface that should be implemented
	// by types that are returned from actions.
	ActionInterfaceName = "Result"
)

// Controller is a type that represents application controller,
// a structure that has actions.
type Controller struct {
	Actions reflect.Funcs // Actions are methods that implement action.Result interface.
}

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

// extractControllers gets a reflect.Package type and returns
// a slice of controllers that are found there.
func extractControllers(pkg *reflect.Package) []Controller {
	return nil
}
