// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	c2 "github.com/colegion/contrib/controllers/sessions"
	c4 "github.com/colegion/contrib/controllers/templates"
)

// controllers stores names of all controllers and actions of the application.
var controllers = url.Values{}

// routes were extracted from controllers of the application.
var routes = []struct {
	Method, Pattern, Label string
	Handler                http.HandlerFunc
}{
	{"CONNECT", "/errors/404", "404", Errors.NotFound},
	{"CONNECT", "/errors/405", "405", Errors.MethodNotAllowed},
	{"CONNECT", "/errors/500", "500", Errors.InternalServerError},
	{"DELETE", "/errors/404", "404", Errors.NotFound},
	{"DELETE", "/errors/405", "405", Errors.MethodNotAllowed},
	{"DELETE", "/errors/500", "500", Errors.InternalServerError},
	{"GET", "/", "", App.Index},
	{"GET", "/*filepath", "", Static.Serve},
	{"GET", "/errors/404", "404", Errors.NotFound},
	{"GET", "/errors/405", "405", Errors.MethodNotAllowed},
	{"GET", "/errors/500", "500", Errors.InternalServerError},
	{"HEAD", "/errors/404", "404", Errors.NotFound},
	{"HEAD", "/errors/405", "405", Errors.MethodNotAllowed},
	{"HEAD", "/errors/500", "500", Errors.InternalServerError},
	{"OPTIONS", "/errors/404", "404", Errors.NotFound},
	{"OPTIONS", "/errors/405", "405", Errors.MethodNotAllowed},
	{"OPTIONS", "/errors/500", "500", Errors.InternalServerError},
	{"PATCH", "/errors/404", "404", Errors.NotFound},
	{"PATCH", "/errors/405", "405", Errors.MethodNotAllowed},
	{"PATCH", "/errors/500", "500", Errors.InternalServerError},
	{"POST", "/errors/404", "404", Errors.NotFound},
	{"POST", "/errors/405", "405", Errors.MethodNotAllowed},
	{"POST", "/errors/500", "500", Errors.InternalServerError},
	{"PUT", "/errors/404", "404", Errors.NotFound},
	{"PUT", "/errors/405", "405", Errors.MethodNotAllowed},
	{"PUT", "/errors/500", "500", Errors.InternalServerError},
	{"TRACE", "/errors/404", "404", Errors.NotFound},
	{"TRACE", "/errors/405", "405", Errors.MethodNotAllowed},
	{"TRACE", "/errors/500", "500", Errors.InternalServerError},
}

// Init initializes controllers of "<no value>",
// its parents, and returns a list of routes along
// with handler functions associated with them.
func Init() []struct {
	Method, Pattern, Label string
	Handler                http.HandlerFunc
} {
	c2.Init(controllers)
	c4.Init(controllers)
	c4.Init(controllers)
	return routes
}
