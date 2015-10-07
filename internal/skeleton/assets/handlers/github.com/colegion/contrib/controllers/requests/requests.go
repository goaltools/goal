// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/requests"

	"github.com/colegion/goal/config"
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

// tRequests is a type with handler methods of Requests controller.
type tRequests struct {
}

// New allocates (github.com/colegion/contrib/controllers/requests).Requests controller,
// then returns it.
func (t tRequests) New() *contr.Requests {
	c := &contr.Requests{}
	return c
}

// Before calls (github.com/colegion/contrib/controllers/requests).Requests.Before.
func (t tRequests) Before(c *contr.Requests, w http.ResponseWriter, r *http.Request) http.Handler {
	// Call magic Before action of (github.com/colegion/contrib/controllers/requests).Requests.
	if res := c.Before( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// After is a dump method that always returns nil.
func (t tRequests) After(c *contr.Requests, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tRequests) Initially(c *contr.Requests, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Call magic Initially method of (github.com/colegion/contrib/controllers/requests).Requests.
	return c.Initially(w, r, a)
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tRequests) Finally(c *contr.Requests, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	return
}

// Init is used to initialize controllers of "github.com/colegion/contrib/controllers/requests"
// and its parents.
func Init(g config.Getter) {
	initRequests(g)
}

func initRequests(g config.Getter) {
}

func init() {
	_ = strconv.MeaningOfLife
}
