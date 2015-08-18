// Package main is an entry point of the application.
package main

import (
	"flag"
	"net/http"
	"runtime"

	"github.com/anonx/sunplate/internal/skeleton/controllers"
	"github.com/anonx/sunplate/internal/skeleton/routes"

	"github.com/anonx/sunplate/log"
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
		log.Error.Fatal(err)
	}

	// Prepare a new server.
	s := &http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	// Starting the server.
	log.Info.Printf(`Listening on "%s".`, s.Addr)
	log.Error.Fatal(serve(s))
}

func init() {
	// Initializing controllers.
	controllers.Init(*root)
}
