// Package main is an entry point of the application.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/colegion/goal/internal/skeleton/utils"

	"github.com/colegion/contrib/servers/grace"
)

var addr = flag.String("http.addr", ":8080", "address the application will listen on")

func main() {
	// Initialize app's HTTP handler.
	handler, err := utils.InitHandler("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a new server.
	s := &http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	// Start the server.
	log.Printf(`Listening on "%s".`, s.Addr)
	log.Fatal(grace.Serve(s))
}
