// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	contr "github.com/colegion/contrib/controllers/requests"

	"github.com/colegion/goal/strconv"
)

// Requests is an insance of tRequests that is automatically generated from Requests controller
// being found at "github.com/colegion/contrib/controllers/requests/requests.go",
// and contains methods to be used as handler functions.
//
// Requests is a controller that does two things:
// 1. Calls Request.ParseForm to parse GET / POST requests;
// 2. Makes Request available in your controller (use c.Request).
var Requests tRequests

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tRequests is a type with handler methods of Requests controller.
type tRequests struct {
}

// New allocates (github.com/colegion/contrib/controllers/requests).Requests controller,
// then returns it.
func (t tRequests) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Requests {
	c := &contr.Requests{

		Request: r,
	}
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tRequests) Before(c *contr.Requests, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call magic Before action of (github.com/colegion/contrib/controllers/requests).Before.
	if h := c.Before(); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tRequests) After(c *contr.Requests, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	return
}

// Init initializes controllers of "github.com/colegion/contrib/controllers/requests",
// its parents, and returns a list of routes along
// with handler functions associated with them.
func Init() (routes []struct {
	Method, Pattern string
	Handler         http.HandlerFunc
}) {

	routes = append(routes, initRequests()...)

	return
}

func initRequests() (rs []struct {
	Method, Pattern string
	Handler         http.HandlerFunc
}) {
	return
}

func init() {
	_ = strconv.MeaningOfLife
}
