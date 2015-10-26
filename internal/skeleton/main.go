// Package main is an entry point of the application.
package main

import (
	"log"

	"github.com/colegion/goal/internal/skeleton/server"
)

func main() {
	log.Fatal(server.Start("config/app.ini"))
}
