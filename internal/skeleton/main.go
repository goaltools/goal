// Package main is an entry point of the application.
package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/colegion/goal/internal/skeleton/assets/handlers"
	"github.com/colegion/goal/internal/skeleton/routes"

	c "github.com/colegion/goal/config"
)

var config = c.New()

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
		Addr:    config.StringDefault("http.addr", ":8080"),
		Handler: handler,
	}

	// Starting the server.
	log.Printf(`Listening on "%s".`, s.Addr)
	log.Fatal(serve(s))
}

func init() {
	// Openning configuration file.
	err := config.ParseFile("config/config.ini", c.ReadFromFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialization of handlers.
	handlers.Init(config)
}
