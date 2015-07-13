// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/middleware/template"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Middleware is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/middleware/template/template.go".
//
// Middleware is a main type that should be embeded into controller structs.
type Middleware struct {
}

// New allocates (github.com/anonx/sunplate/middleware/template).Middleware controller,
// then returns it.
func (t Middleware) New() *contr.Middleware {
	c := &contr.Middleware{}
	return c
}

// Before is a dump method that always returns nil.
func (t Middleware) Before(c *contr.Middleware, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// After is a dump method that always returns nil.
func (t Middleware) After(c *contr.Middleware, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// Finally is a dump method that does nothing.
func (t Middleware) Finally(c *contr.Middleware, w http.ResponseWriter, r *http.Request) {
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderTemplate action found at
// github.com/anonx/sunplate/middleware/template/template.go
// in appropriate order.
//
// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t Middleware) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	c := Middleware{}.New()
	defer Middleware{}.Finally(c, w, r)
	if res := (Middleware{}.Before(c, w, r)); res != nil {
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
	if res := (Middleware{}.After(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

func init() {
	_ = strconv.MeaningOfLife
}