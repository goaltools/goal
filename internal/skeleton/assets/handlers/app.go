// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c5 "github.com/colegion/contrib/controllers/errors"
	c0 "github.com/colegion/contrib/controllers/global"
	c1 "github.com/colegion/contrib/controllers/requests"
	c2 "github.com/colegion/contrib/controllers/sessions"
	c3 "github.com/colegion/contrib/controllers/static"
	c4 "github.com/colegion/contrib/controllers/templates"
	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/strconv"
)

// App is an instance of tApp that is automatically generated from
// App controller being found at "github.com/colegion/goal/internal/skeleton/controllers/app.go",
// and contains methods to be used as handler functions.
//
// App is a sample controller.
var App tApp

// tApp is a type with handler methods of App controller.
type tApp struct {
}

// newC allocates (github.com/colegion/goal/internal/skeleton/controllers).App controller,
// initializes its parents and returns it.
func (t tApp) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.App {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.App{}

	// Allocate its parents. Make sure controller of every type
	// is allocated just once, then reused.
	c.Controllers = &contr.Controllers{}
	c.Controllers.Templates = c.Controllers.Errors.Templates
	c.Controllers.Errors = &c5.Errors{}
	c.Controllers.Static = &c3.Static{}
	c.Controllers.Sessions = &c2.Sessions{

		Request: r,

		Response: w,
	}
	c.Controllers.Requests = &c1.Requests{

		Request: r,

		Response: w,
	}
	c.Controllers.Global = &c0.Global{

		CurrentAction: act,

		CurrentController: ctr,
	}
	c.Controllers.Errors.Templates = &c4.Templates{}
	c.Controllers.Errors.Templates.Requests = c.Controllers.Requests
	c.Controllers.Errors.Templates.Global = c.Controllers.Global
	c.Controllers.Templates.Requests = c.Controllers.Requests
	c.Controllers.Templates.Global = c.Controllers.Global

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tApp) before(c *contr.App, w http.ResponseWriter, r *http.Request) http.Handler {
	// Call special Before actions of the parent controllers.

	if res := c.Controllers.Global.Before(); res != nil {
		return res
	}
	if res := c.Controllers.Requests.Before(); res != nil {
		return res
	}
	if res := c.Controllers.Sessions.Before(); res != nil {
		return res
	}

	if res := c.Controllers.Before(); res != nil {
		return res
	}

	// Call special Before action of (github.com/colegion/goal/internal/skeleton/controllers).App.
	if res := c.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tApp) after(c *contr.App, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	// Execute magic After methods of embedded controllers.
	if res := c.Controllers.Sessions.After(); res != nil {
		return res
	}

	return
}

// Index is a handler that was generated automatically.
// It calls Before, Index, and After actions of a <no value> controller.
// Index action may be found in the "github.com/colegion/goal/internal/skeleton/controllers/app.go".
//
// Index is an action that renders a home page.
//@get /
func (t tApp) Index(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := App.newC(w, r, "App", "Index")
	defer func() {
		// If one of the actions (Before, After or Index) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer App.after(c, w, r) // Call this at the very end, but before applying result.
	if res := App.before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Index(); res != nil {
		h = res
		return
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
