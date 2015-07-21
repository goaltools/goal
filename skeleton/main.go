// Package main is an entry point of the application.
package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/anonx/sunplate/skeleton/routes"

	"github.com/anonx/sunplate/log"
)

func main() {
	// Set max procs for multi-thread executing.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Use either default HTTP address, or get it
	// from the arguments list (the first one is a file name,
	// the second argument is what we need).
	addr := ":8080"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	// Build the routes and handler.
	handler, err := routes.List.Build()
	if err != nil {
		log.Error.Fatal(err)
	}

	// Prepare a server and run it.
	go log.Info.Printf(`Listening on "%s".`, addr)
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Error.Fatal(s.ListenAndServe())
}
