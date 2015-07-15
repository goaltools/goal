// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/generation/handlers/testdata/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// App is an instance of tApp.
var App tApp

// tApp is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/generation/handlers/testdata/controllers/app.go".
//
// App is a sample controller.
type App struct {
}

// New allocates (github.com/anonx/sunplate/generation/handlers/testdata/controllers).App controller,
// initializes its parents; then returns the controller.
func (t tApp) New() *contr.App {
	c := &contr.App{}
	c.Controller = Controller.New()
	return c
}

// Before executes magic actions of embedded controllers.
func (t tApp) Before(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// After executes magic actions of embedded controllers.
func (t tApp) After(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// Finally executes magic actions of embedded controllers.
func (t tApp) Finally(c *contr.App, w http.ResponseWriter, r *http.Request) {
	// Execute magic Finally actions of embedded controllers.
	Controller.Finally(c.Controller, w, r)
}

// HelloWorld is a handler that was generated automatically.
// It calls Before, After, Finally methods, and HelloWorld action found at
// github.com/anonx/sunplate/generation/handlers/testdata/controllers/app.go
// in appropriate order.
//
// HelloWorld is a sample action.
func (t App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	c := App.New()
	defer App.Finally(c, w, r)
	if res := (App.Before(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := c.HelloWorld( // "Binding" parameters.
		strconv.Int(r.Form, "page"),
	); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := (App.After(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

// Index is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Index action found at
// github.com/anonx/sunplate/generation/handlers/testdata/controllers/init.go
// in appropriate order.
//
// Index is a sample action.
func (t App) Index(w http.ResponseWriter, r *http.Request) {
	c := App.New()
	defer App.Finally(c, w, r)
	if res := (App.Before(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := c.Index( // "Binding" parameters.
		strconv.Int(r.Form, "page"),
	); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := (App.After(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

func init() {
	_ = strconv.MeaningOfLife
}