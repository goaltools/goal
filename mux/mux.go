// Package mux is a wrapper on naoina/denco router.
// It uses request.Form for params instead of a separate Params
// argument. So, it requires a bit more memory and it is a bit slower.
// However, the downsides are an acceptable trade off for compatability
// with the standard library.
package mux

import (
	"net/http"

	"github.com/naoina/denco"
)

// Router represents a multiplexer for HTTP request.
type Router struct {
	routers map[string]*denco.Router
}

// Handler represents a handler of HTTP request.
type Handler struct {
	Method  string           // Method is an HTTP methods, e.g. "GET" or "POST".
	Pattern string           // Pattern is a routing path for handler.
	Func    http.HandlerFunc // Func is a function of handler of HTTP request.
}

// New allocates and returns a new multiplexer.
func New() *Router {
	return &Router{}
}

// ServeHTTP is used to implement http.Handler interface.
// It dispatches the request to the handler whose pattern
// most closely matches the request URL.
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, it is being overridden.
func (t *Router) Handle(pattern string, handler Handler) {
}

// HandleFunc registers the handler function for the given pattern.
func (t *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
}
