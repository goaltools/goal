// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/sessions"

	"github.com/colegion/goal/strconv"
)

// Sessions is an instance of tSessions that is automatically generated from
// Sessions controller being found at "github.com/colegion/contrib/controllers/sessions/sessions.go",
// and contains methods to be used as handler functions.
//
// Sessions is a controller that makes Session field
// available for your actions when you're using this
// controller as a parent.
var Sessions tSessions

// tSessions is a type with handler methods of Sessions controller.
type tSessions struct {
}

// newC allocates (github.com/colegion/contrib/controllers/sessions).Sessions controller,
// initializes its parents and returns it.
func (t tSessions) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Sessions {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Sessions{

		Request: r,

		Response: w,
	}

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tSessions) before(c *contr.Sessions, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call special Before action of (github.com/colegion/contrib/controllers/sessions).Sessions.
	if res := c.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tSessions) after(c *contr.Sessions, w http.ResponseWriter, r *http.Request) (res http.Handler) {
	// At the very end call magic After method of (github.com/colegion/contrib/controllers/sessions).Sessions.
	defer func() {
		// If result was not returned by parent After methods, start Sessions's After.
		if res == nil {
			res = c.After()
		}
	}()

	return
}

func init() {
	_ = strconv.MeaningOfLife
}
