// Package routes defines application routes.
package routes

import (
	"net/http"

	h "github.com/colegion/goal/internal/skeleton/assets/handlers"

	r "github.com/colegion/contrib/routers/denco"
)

// List is a slice of routes. If using a default router call Build
// to get an HTTP handler.
var List = r.Routes{
	r.Get("/", h.App.Index),

	// Serve static files of "./static" directory.
	r.Get("/static/*filepath", http.StripPrefix(
		"/static/", http.FileServer(http.Dir("./static")),
	).ServeHTTP),
}

func init() {
	r.NotFound = h.Errors.NotFound
}
