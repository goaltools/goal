// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/middleware/template"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Middleware is an insance of tMiddleware that is automatically generated from Middleware controller
// being found at "github.com/anonx/sunplate/middleware/template/template.go",
// and contains methods to be used as handler functions.
//
// Middleware is a main type that should be embeded into controller structs.
var Middleware tMiddleware

// tMiddleware is a type with handler methods of Middleware controller.
type tMiddleware struct {
}

// New allocates (github.com/anonx/sunplate/middleware/template).Middleware controller,
// then returns it.
func (t tMiddleware) New() *contr.Middleware {
	c := &contr.Middleware{}
	return c
}

// Before is a dump method that always returns nil.
func (t tMiddleware) Before(c *contr.Middleware, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// After is a dump method that always returns nil.
func (t tMiddleware) After(c *contr.Middleware, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// Finally is a dump method that does nothing.
func (t tMiddleware) Finally(c *contr.Middleware, w http.ResponseWriter, r *http.Request) {
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderTemplate action found at
// github.com/anonx/sunplate/middleware/template/template.go
// in appropriate order.
//
// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t tMiddleware) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	c := Middleware.New()
	defer Middleware.Finally(c, w, r)
	if res := Middleware.Before(c, w, r); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := c.RenderTemplate( // "Binding" parameters.
		strconv.String(r.Form, "templatePath"),
	); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := Middleware.After(c, w, r); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
