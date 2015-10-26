// Package server provides functions for starting a new HTTP server.
// When on Windows a standard http.Serve is used. On other platforms
// a server that supports graceful restarts and shutdowns is started.
package server

import (
	"flag"
	"net/http"
	"net/http/httptest"

	"github.com/colegion/goal/internal/skeleton/assets/handlers"
	"github.com/colegion/goal/internal/skeleton/routes"

	"github.com/colegion/contrib/configs/iniflag"
	"github.com/colegion/contrib/servers/grace"
)

var addr = flag.String("http.addr", ":8080", "HTTP address the app must listen on")

// Start starts a new server using parameters from the
// requested INI configuration files.
func Start(configFiles ...string) error {
	// Prepare an HTTP handler.
	h, err := initHandler(configFiles)
	if err != nil {
		return err
	}

	// Allocate a new HTTP server.
	s := &http.Server{
		Addr:    *addr,
		Handler: h,
	}

	// Start the server or return an error.
	return grace.Serve(s)
}

// StartTest is similar to Start but starts an httptest server that listens on
// a random address and may be used for testing of your application.
// The caller should call Close on the returned Server when finished.
func StartTest(configFiles ...string) (*httptest.Server, error) {
	// Prepare an HTTP handler.
	h, err := initHandler(configFiles)
	if err != nil {
		return nil, err
	}

	// Start the server and return it.
	return httptest.NewServer(h), nil
}

// initHandler parses configuration files, populates flags of controllers,
// and returns an initialized go HTTP handler that is based on the app's
// routes and handler functions automatically generated from controllers.
func initHandler(configFiles []string) (http.Handler, error) {
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
