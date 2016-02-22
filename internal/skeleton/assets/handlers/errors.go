// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/errors"
	c0 "github.com/colegion/contrib/controllers/global"
	c1 "github.com/colegion/contrib/controllers/requests"
	c4 "github.com/colegion/contrib/controllers/templates"

	"github.com/colegion/goal/strconv"
)

// Errors is an instance of tErrors that is automatically generated from
// Errors controller being found at "github.com/colegion/contrib/controllers/errors/errors.go",
// and contains methods to be used as handler functions.
//
// Errors is a controller that brings support of errors
// processing to your application.
var Errors tErrors

// tErrors is a type with handler methods of Errors controller.
type tErrors struct {
}

// newC allocates (github.com/colegion/contrib/controllers/errors).Errors controller,
// initializes its parents and returns it.
func (t tErrors) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Errors {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Errors{}

	// Allocate its parents. Make sure controller of every type
	// is allocated just once, then reused.
	c.Templates = &c4.Templates{}
	c.Templates.Requests = &c1.Requests{

		Request: r,

		Response: w,
	}
	c.Templates.Global = &c0.Global{

		CurrentAction: act,

		CurrentController: ctr,
	}

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tErrors) before(c *contr.Errors, w http.ResponseWriter, r *http.Request) http.Handler {
	// Call special Before actions of the parent controllers.
	if res := c.Templates.Global.Before(); res != nil {
		return res
	}
	if res := c.Templates.Requests.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tErrors) after(c *contr.Errors, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	// Execute magic After methods of embedded controllers.

	return
}

// NotFound is a handler that was generated automatically.
// It calls Before, NotFound, and After actions of a <no value> controller.
// NotFound action may be found in the "github.com/colegion/contrib/controllers/errors/errors.go".
//
// NotFound is an action that renders a 404 page not found error.
//@route /404 404
func (t tErrors) NotFound(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.newC(w, r, "Errors", "NotFound")
	defer func() {
		// If one of the actions (Before, After or NotFound) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Errors.after(c, w, r) // Call this at the very end, but before applying result.
	if res := Errors.before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.NotFound(); res != nil {
		h = res
		return
	}
}

// MethodNotAllowed is a handler that was generated automatically.
// It calls Before, MethodNotAllowed, and After actions of a <no value> controller.
// MethodNotAllowed action may be found in the "github.com/colegion/contrib/controllers/errors/errors.go".
//
// MethodNotAllowed is an action that renders a 405 method not allowed error.
//@route /405 405
func (t tErrors) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.newC(w, r, "Errors", "MethodNotAllowed")
	defer func() {
		// If one of the actions (Before, After or MethodNotAllowed) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Errors.after(c, w, r) // Call this at the very end, but before applying result.
	if res := Errors.before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.MethodNotAllowed(); res != nil {
		h = res
		return
	}
}

// InternalServerError is a handler that was generated automatically.
// It calls Before, InternalServerError, and After actions of a <no value> controller.
// InternalServerError action may be found in the "github.com/colegion/contrib/controllers/errors/errors.go".
//
// InternalServerError is an action that renders a 500 internal server error.
//@route /500 500
func (t tErrors) InternalServerError(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.newC(w, r, "Errors", "InternalServerError")
	defer func() {
		// If one of the actions (Before, After or InternalServerError) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Errors.after(c, w, r) // Call this at the very end, but before applying result.
	if res := Errors.before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.InternalServerError(); res != nil {
		h = res
		return
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
