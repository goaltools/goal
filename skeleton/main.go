// Package main is an entry point of the application.
package main

import (
	"net/http"
	"runtime"

	"github.com/anonx/sunplate/example/routes"

	"github.com/anonx/sunplate/log"
)

// Comments below are used by `go generate`.
// Please, DO NOT EDIT if you do not know what you are doing.
//
//go:generate sunplate generate handlers
//go:generate sunplate generate listing

func main() {
	// Set max procs for multi-thread executing.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Build the routes and run the app.
	handler, err := routes.List.Build()
	if err != nil {
		log.Error.Fatal(err)
	}
	log.Error.Fatal(http.ListenAndServe(":8080", handler))
}
