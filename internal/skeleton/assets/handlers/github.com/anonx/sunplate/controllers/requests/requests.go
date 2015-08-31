// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/controllers/requests"

	"github.com/anonx/sunplate/config"
	"github.com/anonx/sunplate/strconv"
)

// Requests is an insance of tRequests that is automatically generated from Requests controller
// being found at "github.com/anonx/sunplate/controllers/requests/requests.go",
// and contains methods to be used as handler functions.
//
// Requests is a controller that does two things:
// 1. Calls Request.ParseForm to parse GET / POST requests;
// 2. Makes Request available in your controller (use c.Request).
var Requests tRequests

// tRequests is a type with handler methods of Requests controller.
type tRequests struct {
}

// New allocates (github.com/anonx/sunplate/controllers/requests).Requests controller,
// then returns it.
func (t tRequests) New() *contr.Requests {
	c := &contr.Requests{}
	return c
}

// Before is a dump method that always returns nil.
func (t tRequests) Before(c *contr.Requests, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// After is a dump method that always returns nil.
func (t tRequests) After(c *contr.Requests, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tRequests) Initially(c *contr.Requests, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Call magic Initially method of (github.com/anonx/sunplate/controllers/requests).Requests.
	return c.Initially(w, r)
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tRequests) Finally(c *contr.Requests, w http.ResponseWriter, r *http.Request) (finish bool) {
	return
}

// Init is used to initialize controllers of "github.com/anonx/sunplate/controllers/requests"
// and its parents.
func Init(g config.Getter) {
	initRequests(g)
}

func initRequests(g config.Getter) {
}

func init() {
	_ = strconv.MeaningOfLife
}
