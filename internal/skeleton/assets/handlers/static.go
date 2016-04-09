// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/static"

	"github.com/colegion/goal/strconv"
)

// Static is an instance of tStatic that is automatically generated from
// Static controller being found at "github.com/colegion/contrib/controllers/static/static.go",
// and contains methods to be used as handler functions.
//
// Static is a controller that brings static
// assets' serving functionality to your app.
var Static tStatic

// tStatic is a type with handler methods of Static controller.
type tStatic struct {
}

// newC allocates (github.com/colegion/contrib/controllers/static).Static controller,
// initializes its parents and returns it.
func (t tStatic) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Static {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Static{}

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tStatic) before(c *contr.Static, w http.ResponseWriter, r *http.Request) http.Handler {

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tStatic) after(c *contr.Static, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	return
}

// Serve is a handler that was generated automatically.
// It calls Before, Serve, and After actions of a <no value> controller.
// Serve action may be found in the "github.com/colegion/contrib/controllers/static/static.go".
//
// Serve is a wrapper around Go's standard FileServer
// and StripPrefix HTTP handlers.
//@get /*filepath
func (t tStatic) Serve(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Static.newC(w, r, "Static", "Serve")
	defer func() {
		// If one of the actions (Before, After or Serve) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Static.after(c, w, r) // Call this at the very end, but before applying result.
	if res := Static.before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Serve(
		strconv.String(r.Form, "filepath"),
	); res != nil {
		h = res
		return
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
