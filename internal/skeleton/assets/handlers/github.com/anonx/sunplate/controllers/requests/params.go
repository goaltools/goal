// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/controllers/requests"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Params is an insance of tParams that is automatically generated from Params controller
// being found at "github.com/anonx/sunplate/controllers/requests/params.go",
// and contains methods to be used as handler functions.
//
// Params is used for parsing query from the URL
// and body of requests.
var Params tParams

// tParams is a type with handler methods of Params controller.
type tParams struct {
}

// New allocates (github.com/anonx/sunplate/controllers/requests).Params controller,
// then returns it.
func (t tParams) New() *contr.Params {
	c := &contr.Params{}
	return c
}

// Before is a dump method that always returns nil.
func (t tParams) Before(c *contr.Params, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// After is a dump method that always returns nil.
func (t tParams) After(c *contr.Params, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tParams) Initially(c *contr.Params, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Call magic Initially method of (github.com/anonx/sunplate/controllers/requests).Params.
	return c.Initially(w, r)
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tParams) Finally(c *contr.Params, w http.ResponseWriter, r *http.Request) (finish bool) {
	return
}

func init() {
	_ = strconv.MeaningOfLife
}
