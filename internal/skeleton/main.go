// Package main is an entry point of the application.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/colegion/goal/internal/skeleton/assets/handlers"
	"github.com/colegion/goal/internal/skeleton/routes"

	"github.com/colegion/contrib/configs/iniflag"
	"github.com/colegion/contrib/servers/grace"
)

var addr = flag.String("http.addr", ":8080", "address the application will listen on")

func main() {
	// Parse configuration files and flags.
	err := iniflag.Parse("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	// Initialization of handlers.
	handlers.Init()

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
	log.Fatal(grace.Serve(s))
}
