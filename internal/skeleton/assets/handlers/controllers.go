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

// Controllers is an instance of tControllers that is automatically generated from
// Controllers controller being found at "github.com/colegion/goal/internal/skeleton/controllers/init.go",
// and contains methods to be used as handler functions.
//
// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
var Controllers tControllers

// tControllers is a type with handler methods of Controllers controller.
type tControllers struct {
}

// newC allocates (github.com/colegion/goal/internal/skeleton/controllers).Controllers controller,
// initializes its parents and returns it.
func (t tControllers) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Controllers {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Controllers{}

	// Allocate its parents. Make sure controller of every type
	// is allocated just once, then reused.
	c.Templates = c.Errors.Templates
	c.Errors = &c5.Errors{}
	c.Static = &c3.Static{}
	c.Sessions = &c2.Sessions{

		Request: r,

		Response: w,
	}
	c.Requests = &c1.Requests{

		Request: r,

		Response: w,
	}
	c.Global = &c0.Global{

		CurrentAction: act,

		CurrentController: ctr,
	}
	c.Errors.Templates = &c4.Templates{}
	c.Errors.Templates.Requests = c.Requests
	c.Errors.Templates.Global = c.Global
	c.Templates.Requests = c.Requests
	c.Templates.Global = c.Global

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tControllers) before(c *contr.Controllers, w http.ResponseWriter, r *http.Request) http.Handler {
	// Call special Before actions of the parent controllers.

	if res := c.Global.Before(); res != nil {
		return res
	}
	if res := c.Requests.Before(); res != nil {
		return res
	}
	if res := c.Sessions.Before(); res != nil {
		return res
	}

	// Call special Before action of (github.com/colegion/goal/internal/skeleton/controllers).Controllers.
	if res := c.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tControllers) after(c *contr.Controllers, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	// Execute magic After methods of embedded controllers.
	if res := c.Sessions.After(); res != nil {
		return res
	}

	return
}

func init() {
	_ = strconv.MeaningOfLife
}
