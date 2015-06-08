// Package routing is a wrapper on naoina/denco router.
// It uses request.Form for params instead of a separate Params
// argument. So, it requires a more memory and it is a bit slower.
// However, the downsides are an acceptable trade off for compatability
// with the standard library.
//
// A sample of its usage is below:
//
//	package main
//
//	import (
//		"log"
//		"net/http"
//
//		"github.com/anonx/sunplate/routing"
//	)
//
//	func main() {
//		r := routing.New()
//		err := r.Handle(routing.Routes{
//			r.Get("/profiles/:username", ShowUserHandleFunc),
//			r.Delete("/profiles/:username", DeleteUserHandleFunc),
//		}).Build()
//		if err != nil {
//			panic(err)
//		}
//		log.Fatal(http.ListenAndServe(":8080", r))
//	}
package routing

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/naoina/denco"
)

// Router represents a multiplexer for HTTP requests.
type Router struct {
	data    *denco.Router  // data stores denco router.
	indexes map[string]int // indexes is used to simplify search of records we need.
	records []denco.Record // records is a list of handlers expected by denco router.
}

// Routes is an alias of []Route.
type Routes []Route

// Route is used to store information about HTTP request's handler
// including a list of allowed methods and pattern.
type Route struct {
	handlers *dict  // HTTP request method -> handler pairs.
	pattern  string // Pattern is a routing path for handler.
}

// dict is a dictionary structure that is used by routing package instead of map
// for small sets of data.
// On average efficency of getting an element from map is O(c + 1).
// At the same time efficency of iterating over a slice is O(n).
// And when n is small, O(n) < O(c + 1). That's why we are using simple loop rather than
// a map function.
type dict struct {
	keys   []string
	values []*http.HandlerFunc
}

// newDict allocates and returns a dict structure.
func newDict() *dict {
	return &dict{
		keys:   []string{},
		values: []*http.HandlerFunc{},
	}
}

// set expects key and value as input parameters that are
// saved to the dict.
func (t *dict) set(k string, v *http.HandlerFunc) {
	// Check whether we have already had such key.
	if _, i := t.get(k); i >= 0 {
		// If so, update it.
		t.values[i] = v
		return
	}
	// Otherwise, add a new key-value pair.
	t.keys = append(t.keys, k)
	t.values = append(t.values, v)
}

// get receives a key as input and returns associated value.
func (t *dict) get(k string) (*http.HandlerFunc, int) {
	for i, v := range t.keys {
		if v == k {
			return t.values[i], i
		}
	}
	return nil, -1
}

// join receives a new dict and appends it to the dict.
func (t *dict) join(d *dict) {
	// Iterate through all keys of a new dict.
	for i, k := range d.keys {
		// Add them to the main dict.
		t.set(k, d.values[i])
	}
}

// NewRouter allocates and returns a new multiplexer.
func NewRouter() *Router {
	return &Router{
		indexes: map[string]int{},
	}
}

// Get is an short form of Route("GET", pattern, handler).
func (t *Router) Get(pattern string, handler http.HandlerFunc) Route {
	return t.Route("GET", pattern, handler)
}

// Post is a short form of Route("POST", pattern, handler).
func (t *Router) Post(pattern string, handler http.HandlerFunc) Route {
	return t.Route("POST", pattern, handler)
}

// Put is a short form of Route("PUT", pattern, handler).
func (t *Router) Put(pattern string, handler http.HandlerFunc) Route {
	return t.Route("PUT", pattern, handler)
}

// Head is a short form of Route("HEAD", pattern, handler).
func (t *Router) Head(pattern string, handler http.HandlerFunc) Route {
	return t.Route("HEAD", pattern, handler)
}

// Delete is a short form of Route("DELETE", pattern, handler).
func (t *Router) Delete(pattern string, handler http.HandlerFunc) Route {
	return t.Route("DELETE", pattern, handler)
}

// ServeHTTP is used to implement http.Handler interface.
// It dispatches the request to the handler whose pattern
// most closely matches the request URL.
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, _ := t.Handler(r)
	h.ServeHTTP(w, r)
}

// Handle registers handlers for given patterns.
// If a handler already exists for pattern, it will be overridden.
// If it exists but with another method, a new method will be added.
func (t *Router) Handle(routes Routes) *Router {
	for _, route := range routes {
		// Check whether we have already had such route.
		index, ok := t.indexes[route.pattern]

		// If we haven't, add the route.
		if !ok {
			// Save pattern's index to simplify its search
			// in next iteration.
			t.indexes[route.pattern] = len(t.records)

			// Add the route to the slice.
			t.records = append(t.records, denco.NewRecord(route.pattern, route))
			continue
		}

		// Otherwise, just add new HTTP methods to the existing route.
		r := t.records[index].Value.(Route)
		r.handlers.join(route.handlers)
	}
	return t
}

// Build compiles registered routes. Routes that are added after building will not
// be handled. A new call to build will be required.
func (t *Router) Build() error {
	t.data = denco.New()
	return t.data.Build(t.records)
}

// Route allocates and returns a Route struct.
func (t *Router) Route(method, pattern string, handler http.HandlerFunc) Route {
	hs := newDict()
	hs.set(strings.ToUpper(method), &handler)
	return Route{
		handlers: hs,
		pattern:  pattern,
	}
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
	route := obj.(Route)
	handler, ok := route.handlers.get(r.Method)
	if ok == -1 {
		return http.HandlerFunc(MethodNotAllowed), route.pattern
	}

	// Add parameters of request to request.Form and return a handler.
	if r.Form == nil { // Make sure Form is initialized.
		r.Form = url.Values{}
	}
	for _, param := range params {
		r.Form.Add(param.Name, param.Value)
	}
	return handler, route.pattern
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
