// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/colegion/goal/internal/skeleton/assets/handlers/github.com/colegion/contrib/controllers/requests"
	c2 "github.com/colegion/goal/internal/skeleton/assets/handlers/github.com/colegion/contrib/controllers/sessions"
	c1 "github.com/colegion/goal/internal/skeleton/assets/handlers/github.com/colegion/contrib/controllers/templates"
	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/strconv"
)

// Controllers is an insance of tControllers that is automatically generated from Controllers controller
// being found at "github.com/colegion/goal/internal/skeleton/controllers/init.go",
// and contains methods to be used as handler functions.
//
// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
var Controllers tControllers

// tControllers is a type with handler methods of Controllers controller.
type tControllers struct {
}

// New allocates (github.com/colegion/goal/internal/skeleton/controllers).Controllers controller,
// initializes its parents; then returns the controller.
func (t tControllers) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Controllers {
	c := &contr.Controllers{}
	c.Requests = c0.Requests.New(w, r, ctr, act)
	c.Templates = c1.Templates.New(w, r, ctr, act)
	c.Sessions = c2.Sessions.New(w, r, ctr, act)
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tControllers) Before(c *contr.Controllers, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if h := c0.Requests.Before(c.Requests, w, r); h != nil {
		return h
	}

	if h := c1.Templates.Before(c.Templates, w, r); h != nil {
		return h
	}

	if h := c2.Sessions.Before(c.Sessions, w, r); h != nil {
		return h
	}

	// Call magic Before action of (github.com/colegion/goal/internal/skeleton/controllers).Before.
	if h := c.Before(); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tControllers) After(c *contr.Controllers, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	// Execute magic After methods of embedded controllers.

	if h = c0.Requests.After(c.Requests, w, r); h != nil {
		return h
	}

	if h = c1.Templates.After(c.Templates, w, r); h != nil {
		return h
	}

	if h = c2.Sessions.After(c.Sessions, w, r); h != nil {
		return h
	}

	return
}

func initControllers() {

	c0.Init()

	c1.Init()

	c2.Init()

}

func init() {
	_ = strconv.MeaningOfLife
}
