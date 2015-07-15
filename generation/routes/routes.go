// Package routes is a subcommand for scan of a directory
// with handlers and generation of a routes package for them.
package routes

import (
	h "github.com/anonx/sunplate/example/assets/handlers"
	r "github.com/anonx/sunplate/routing"
)

func init() {
	_ = r.Routes{
		r.Get("/", h.App.Index),
	}
}
