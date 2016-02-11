// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	contr "github.com/colegion/contrib/controllers/sessions"

	"github.com/colegion/goal/strconv"
)

// Sessions is an insance of tSessions that is automatically generated from Sessions controller
// being found at "github.com/colegion/contrib/controllers/sessions/sessions.go",
// and contains methods to be used as handler functions.
//
// Sessions is a controller that makes Session field
// available for your actions when you're using this
// controller as a parent.
var Sessions tSessions

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tSessions is a type with handler methods of Sessions controller.
type tSessions struct {
}

// New allocates (github.com/colegion/contrib/controllers/sessions).Sessions controller,
// then returns it.
func (t tSessions) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Sessions {
	c := &contr.Sessions{

		Request: r,

		Response: w,
	}
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tSessions) Before(c *contr.Sessions, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call magic Before action of (github.com/colegion/contrib/controllers/sessions).Before.
	if h := c.Before(); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tSessions) After(c *contr.Sessions, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	// Call magic After method of (github.com/colegion/contrib/controllers/sessions).Sessions.
	defer func() {
		if h == nil {
			h = c.After()
		}
	}()

	return
}

// Init is used to initialize controllers of "github.com/colegion/contrib/controllers/sessions"
// and its parents.
func Init() {

	initSessions()

	contr.Init(context)

}

func initSessions() {

}

func init() {
	_ = strconv.MeaningOfLife
}
