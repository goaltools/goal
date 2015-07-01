// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"strconv"

	contr "github.com/anonx/sunplate/middleware/template"
)

// Middleware is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/middleware/template/template.go".
//
// Middleware is a main type that should be embeded into controller structs.
type Middleware struct {
}

// New allocates (github.com/anonx/sunplate/middleware/template).Middleware controller,
// initializes its parents; then returns.
func (t Middleware) New() *contr.Middleware {
	c := &contr.Middleware{}
	return c
}

// Before executes magic Before actions of parent controllers.
func (t Middleware) Before(c *contr.Middleware, w http.ResponseWriter, r *http.Request) {
	// Execute magic Before actions of parent controllers.

}

func init() {
	_ = strconv.IntSize
}
