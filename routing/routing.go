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

	"github.com/naoina/denco"
)

// Router represents a multiplexer for HTTP requests.
type Router struct {
	data    *denco.Router  // data stores denco router.
	indexes map[string]int // indexes are used to check whether a record exists.
	records []denco.Record // records is a list of handlers expected by denco router.
}

// Routes is an alias of []Route.
type Routes []Route

// Route is used to store information about HTTP request's handler
// including a list of allowed methods and pattern.
type Route struct {
	handler *http.HandlerFunc // HTTP request handler function.
	methods []string          // A list of allowed HTTP methods (e.g. "GET" or "POST").
	pattern string            // Pattern is a routing path for handler.
}

// NewRouter allocates and returns a new multiplexer.
func NewRouter() *Router {
	return &Router{
		indexes: map[string]int{},
	}
}

// HasMethod checks whether specific method request is allowed.
func (t *Route) HasMethod(name string) bool {
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
	h, _ := t.Handler(r)
	h.ServeHTTP(w, r)
}

// Handle registers handlers for given patterns.
// If a handler already exists for pattern, it will be overridden.
// If it exists but with another method, a new method will be added.
func (t *Router) Handle(routes Routes) {
	for _, route := range routes {
		// Check whether we have already had such route.
		index, ok := t.indexes[route.pattern]
		if ok {
			// If we haven't, add it.
			t.records = append(t.records, denco.NewRecord(route.pattern, route))
			continue
		}

		// Check whether existing route has the same handler and
		// we are just trying to add a new method.
		r := t.records[index].Value.(Route)
		if r.handler != route.handler {
			// If we aren't, override an old route.
			t.records[index] = denco.NewRecord(route.pattern, route)
			continue
		}

		// Otherwise, add all methods that haven't added yet.
		for _, v := range r.methods {
			if !route.HasMethod(v) {
				route.methods = append(route.methods, v)
			}
		}
	}
}

// Build compiles registered routes. Routes that are added after building will not
// be handled. A new call to build will be required.
func (t *Router) Build() error {
	router := denco.New()
	err := router.Build(t.records)
	if err != nil {
		return err
	}
	return nil
}

// Route allocates and returns a Route struct.
func (t *Router) Route(method, pattern string, handler http.HandlerFunc) Route {
	return Route{
		handler: &handler,
		methods: []string{method},
		pattern: pattern,
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
	if !route.HasMethod(r.Method) {
		return http.HandlerFunc(MethodNotAllowed), route.pattern
	}

	// Add parameters of request to request.Form and return a handler.
	for _, param := range params {
		r.Form.Add(param.Name, param.Value)
	}
	return route.handler, route.pattern
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
