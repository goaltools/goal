// Package routing is a wrapper on naoina/denco router.
// It uses request.Form for params instead of a separate Params
// argument. So, it requires a more memory and it is a bit slower.
// However, the downsides are an acceptable trade off for compatability
// with the standard library.
package routing

import (
	"net/http"

	"github.com/naoina/denco"
)

// Router represents a multiplexer for HTTP request.
type Router struct {
	data *denco.Router
}

// info stores information about HTTP request's handler, its
// pattern and methods.
type info struct {
	handler http.HandlerFunc // HTTP request handler.
	methods []string         // A list of allowed HTTP methods (e.g. "GET" or "POST").
	pattern string           // Pattern is a routing path for handler.
}

// New allocates and returns a new multiplexer.
func New() *Router {
	return &Router{}
}

// HasMethod checks whether specific method request is allowed.
func (t *info) HasMethod(name string) bool {
	// We are iterating through a slice of method strings
	// rather than using a map as there are only a few possible values.
	// So, hash function will require more time than a simple loop.
	for _, v := range t.methods {
		if v == name {
			return true
		}
	}
	return false
}

// ServeHTTP is used to implement http.Handler interface.
// It dispatches the request to the handler whose pattern
// most closely matches the request URL.
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, it is being overridden.
func (t *Router) Handle(method, pattern string, handler http.Handler) {
}

// HandleFunc registers the handler function for the given pattern.
func (t *Router) HandleFunc(method, pattern string, handler http.HandlerFunc) {
}

// Handler returns the handler to use for the given request, consulting r.Method
// and r.URL.Path. It always returns a non-nil handler. If there is no registered handler
// that applies to the request, Handler returns a “page not found” handler and empty pattern.
// If there is a registered handler but requested method is not allowed,
// "method not allowed" and a pattern are returned.
func (t *Router) Handler(r *http.Request) (handler http.Handler, pattern string) {
	// Make sure we have a handler for this request.
	obj, params, found := t.data.Lookup(r.URL.Path)
	if !found {
		return http.HandlerFunc(NotFound), ""
	}

	// Check whether requested method is allowed.
	data := obj.(info)
	if !data.HasMethod(r.Method) {
		return http.HandlerFunc(MethodNotAllowed), data.pattern
	}

	// Add parameters of request to request.Form and return a handler.
	for _, param := range params {
		r.Form.Add(param.Name, param.Value)
	}
	return data.handler, data.pattern
}

// MethodNotAllowed replies to the request with an HTTP 405 method not allowed
// error. If you want to use your own MethodNotAllowed handler, please override
// this variable.
var MethodNotAllowed = func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
}

// NotFound replies to the request with an HTTP 404 not found error.
// NotFound is called when unknown HTTP method or a handler not found.
// If you want to use the your own NotFound handler, please overwrite this variable.
var NotFound = func(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
