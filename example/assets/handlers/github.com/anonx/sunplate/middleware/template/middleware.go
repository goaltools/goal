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

// Before is a dump function that always returns nil.
func (t Middleware) Before(c *contr.Middleware, w http.ResponseWriter, r *http.Request) a.Result {
	// Continue execution chain.
	return nil
}

// After is a dump function that always returns nil.
func (t Middleware) After(c *contr.Middleware, w http.ResponseWriter, r *http.Request) a.Result {
	// Continue execution chain.
	return nil
}

func init() {
	_ = strconv.MeaningOfLife
}
