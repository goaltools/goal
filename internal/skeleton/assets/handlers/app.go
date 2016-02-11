// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/strconv"
)

// App is an insance of tApp that is automatically generated from App controller
// being found at "github.com/colegion/goal/internal/skeleton/controllers/app.go",
// and contains methods to be used as handler functions.
//
// App is a sample controller.
var App tApp

// tApp is a type with handler methods of App controller.
type tApp struct {
}

// New allocates (github.com/colegion/goal/internal/skeleton/controllers).App controller,
// initializes its parents; then returns the controller.
func (t tApp) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.App {
	c := &contr.App{}
	c.Controllers = Controllers.New(w, r, ctr, act)
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tApp) Before(c *contr.App, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if h := Controllers.Before(c.Controllers, w, r); h != nil {
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
func (t tApp) After(c *contr.App, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	// Execute magic After methods of embedded controllers.

	if h = Controllers.After(c.Controllers, w, r); h != nil {
		return h
	}

	return
}

// Index is a handler that was generated automatically.
// It calls Before, After methods, and Index action found at
// github.com/colegion/goal/internal/skeleton/controllers/app.go
// in appropriate order.
//
// Index is an action that renders a home page.
func (t tApp) Index(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := App.New(w, r, "App", "Index")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer App.After(c, w, r)
	if res := App.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Index(); res != nil {
		h = res
		return
	}
}

func initApp() {

	context.Add("App", "Index")

}

func init() {
	_ = strconv.MeaningOfLife
}
