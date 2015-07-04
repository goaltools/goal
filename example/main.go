// Package main is an entry point of the application.
package main

import (
	"net/http"
	"runtime"

	"github.com/anonx/sunplate/log"
	"github.com/anonx/sunplate/routing"
)

// Comments below are used by `go generate`.
// Please, DO NOT EDIT if you do not know what you are doing.
//
//go:generate sunplate generate handlers
//go:generate sunplate generate listing

func main() {
	// Set max procs for multi-thread executing.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Include handlers and run the app.
	r := routing.NewRouter()
	err := r.Handle(routing.Routes{
		nil,
	}).Build()
	if err != nil {
		log.Error.Fatal(err)
	}
	log.Error.Fatal(http.ListenAndServe(":8080", r))
}
