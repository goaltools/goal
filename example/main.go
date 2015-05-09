// Package main is an entry point of the application.
package main

import (
	"net/http"
	"runtime"

	"github.com/anonx/sunplate/log"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Set max procs for multi-thread executing.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Include handlers and run the app.
	router := httprouter.New()
	log.Error.Fatal(http.ListenAndServe(":8080", router))
}
