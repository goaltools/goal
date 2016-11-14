// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	c0 "github.com/goaltools/goal/internal/skeleton/assets/handlers/github.com/goaltools/contrib/controllers/requests"
	c2 "github.com/goaltools/goal/internal/skeleton/assets/handlers/github.com/goaltools/contrib/controllers/sessions"
	c3 "github.com/goaltools/goal/internal/skeleton/assets/handlers/github.com/goaltools/contrib/controllers/static"
	c1 "github.com/goaltools/goal/internal/skeleton/assets/handlers/github.com/goaltools/contrib/controllers/templates"
	contr "github.com/goaltools/goal/internal/skeleton/controllers"

	"github.com/goaltools/goal/strconv"
)

// Controllers is an insance of tControllers that is automatically generated from Controllers controller
// being found at "github.com/goaltools/goal/internal/skeleton/controllers/init.go",
// and contains methods to be used as handler functions.
//
// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
var Controllers tControllers

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tControllers is a type with handler methods of Controllers controller.
type tControllers struct {
}

// New allocates (github.com/goaltools/goal/internal/skeleton/controllers).Controllers controller,
// initializes its parents; then returns the controller.
func (t tControllers) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Controllers {
	c := &contr.Controllers{}
	c.Requests = c0.Requests.New(w, r, ctr, act)
	c.Templates = c1.Templates.New(w, r, ctr, act)
	c.Sessions = c2.Sessions.New(w, r, ctr, act)
	c.Static = c3.Static.New(w, r, ctr, act)
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

	if h := c3.Static.Before(c.Static, w, r); h != nil {
		return h
	}

	// Call magic Before action of (github.com/goaltools/goal/internal/skeleton/controllers).Before.
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

	if h = c3.Static.After(c.Static, w, r); h != nil {
		return h
	}

	return
}

// Init initializes controllers of "github.com/goaltools/goal/internal/skeleton/controllers",
// its parents, and returns a list of routes along
// with handler functions associated with them.
func Init() (routes []struct {
	Method, Pattern, Label string
	Handler                http.HandlerFunc
}) {

	routes = append(routes, initApp()...)

	routes = append(routes, initControllers()...)

	return
}

func initControllers() (rs []struct {
	Method, Pattern, Label string
	Handler                http.HandlerFunc
}) {
	rs = append(rs, c0.Init()...)

	rs = append(rs, c1.Init()...)

	rs = append(rs, c2.Init()...)

	rs = append(rs, c3.Init()...)

	return
}

func init() {
	_ = strconv.MeaningOfLife
}
