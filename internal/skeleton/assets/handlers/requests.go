// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/requests"

	"github.com/colegion/goal/strconv"
)

// Requests is an instance of tRequests that is automatically generated from
// Requests controller being found at "github.com/colegion/contrib/controllers/requests/requests.go",
// and contains methods to be used as handler functions.
//
// Requests is a controller that does two things:
// 1. Calls Request.ParseForm to parse GET / POST requests;
// 2. Makes Request available in your controller (use c.Request).
var Requests tRequests

// tRequests is a type with handler methods of Requests controller.
type tRequests struct {
}

// newC allocates (github.com/colegion/contrib/controllers/requests).Requests controller,
// initializes its parents and returns it.
func (t tRequests) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Requests {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Requests{

		Request: r,

		Response: w,
	}

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tRequests) before(c *contr.Requests, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call special Before action of (github.com/colegion/contrib/controllers/requests).Requests.
	if res := c.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tRequests) after(c *contr.Requests, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	return
}

func init() {
	_ = strconv.MeaningOfLife
}
