// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/example/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// App is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/example/controllers/app.go".
//
// App is a sample controller that is used for demonstration purposes.
type App struct {
}

// New allocates (github.com/anonx/sunplate/example/controllers).App controller,
// initializes its parents; then returns the controller.
func (t App) New() *contr.App {
	c := &contr.App{}
	c.Controller = Controller{}.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// Before calls (github.com/anonx/sunplate/example/controllers).App.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t App) Before(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic Before action of (github.com/anonx/sunplate/example/controllers).App.
	if res := c.Before( // "Binding" parameters.
		strconv.String(r.Form, "name"),
		strconv.Ints(r.Form, "pages"),
	); res != nil {
		return res
	}
	return nil
}

// After executes magic actions of embedded controllers, and
// After calls (github.com/anonx/sunplate/example/controllers).App.After with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t App) After(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic After action of (github.com/anonx/sunplate/example/controllers).App.
	if res := c.After( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// Finally executes magic actions of embedded controllers, and
// Finally calls (github.com/anonx/sunplate/example/controllers).App.Finally with arguments
// that are extracted from r.Form and converted to appropriate types.
// The call is garanteed to be done no matter what happens during execution of parent actions.
func (t App) Finally(c *contr.App, w http.ResponseWriter, r *http.Request) {
	// Call magic Finally action of (github.com/anonx/sunplate/example/controllers).App.
	defer func() {
		if res := c.Finally( // "Binding" parameters.
			strconv.String(r.Form, "name"),
		); res != nil {
			res.Apply(w, r)
		}
	}()
	// Execute magic Finally actions of embedded controllers.
	Controller{}.Finally(c.Controller, w, r)
}

// Index is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Index action found at
// github.com/anonx/sunplate/example/controllers/app.go
// in appropriate order.
//
// Index is an action that is used for generation of a greeting form.
func (t App) Index(w http.ResponseWriter, r *http.Request) {
	c := App{}.New()
	defer App{}.Finally(c, w, r)
	if res := (App{}.Before(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := c.Index( // "Binding" parameters.
	); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := (App{}.After(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

// PostGreet is a handler that was generated automatically.
// It calls Before, After, Finally methods, and PostGreet action found at
// github.com/anonx/sunplate/example/controllers/app.go
// in appropriate order.
//
// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (t App) PostGreet(w http.ResponseWriter, r *http.Request) {
	c := App{}.New()
	defer App{}.Finally(c, w, r)
	if res := (App{}.Before(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := c.PostGreet( // "Binding" parameters.
		strconv.String(r.Form, "name"),
	); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := (App{}.After(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
