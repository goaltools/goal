// Package utils provides building blocks that are reused across
// the application (e.g. in tests and main package) but are not related
// to the business logics. Everything that is used for initialization and
// start of the app but hasn't got its own package must be located here.
package utils

import (
	"net/http"

	"github.com/colegion/goal/internal/skeleton/assets/handlers"
	"github.com/colegion/goal/internal/skeleton/routes"

	"github.com/colegion/contrib/configs/iniflag"
)

// InitHandler gets a number of INI configuration file paths.
// It parses them, populates flags with the data extracted from the configs
// and input arguments to the program, and returns an initialized
// go HTTP handler that is based on the app's routes and handler functions
// automatically generated from controllers.
// Example:
//	h, err := InitHandler("basic.ini", "more/important.ini", "most/important.ini")
//	if err != nil {
//		panic(err)
//	}
func InitHandler(configFiles ...string) (http.Handler, error) {
	// Parse configuration files and flags.
	err := iniflag.Parse(configFiles...)
	if err != nil {
		return nil, err
	}

	// Initialize handler functions and controllers.
	handlers.Init()

	// Build a standard HTTP handler from routes and handler functions.
	return routes.List.Build()
}
