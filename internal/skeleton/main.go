// Package main is an entry point of the application.
package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	"github.com/anonx/sunplate/internal/skeleton/controllers"
	"github.com/anonx/sunplate/internal/skeleton/routes"
)

var (
	addr = flag.String("addr", ":8080", "HTTP address that should be used by the app")
	root = flag.String("root", "./", "Path to the root directory of the project")
)

func main() {
	// Set max procs for multi-thread executing.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Build the routes and handler.
	handler, err := routes.List.Build()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a new server.
	s := &http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	// Starting the server.
	log.Printf(`Listening on "%s".`, s.Addr)
	log.Fatal(serve(s))
}

func init() {
	// Initializing controllers.
	controllers.Init(*root)
}
