package routes

import (
	"net/http"

	h "github.com/colegion/goal/internal/skeleton/assets/handlers"

	r "github.com/colegion/contrib/routers/denco"
)

// List is a slice of routes of the following form:
//	Route:
//		Pattern
//		Handlers:
//			Method: Handler
// If using a standard router just call Context.Build() to get http handler
// as the first argument and an error (or nil) as the second one.
var List = r.Routes{
	r.Get("/", h.App.Index),
	r.Get("/greet/:name", h.App.PostGreet),
	r.Post("/greet/:name", h.App.PostGreet),

	// Serve static files of ./static directory.
	r.Get("/static/*filepath", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP),
}
