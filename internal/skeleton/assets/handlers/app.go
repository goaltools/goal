// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/internal/skeleton/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// App is an insance of tApp that is automatically generated from App controller
// being found at "github.com/anonx/sunplate/internal/skeleton/controllers/app.go",
// and contains methods to be used as handler functions.
//
// App is a sample controller that is used for demonstration purposes.
var App tApp

// tApp is a type with handler methods of App controller.
type tApp struct {
}

// New allocates (github.com/anonx/sunplate/internal/skeleton/controllers).App controller,
// initializes its parents; then returns the controller.
func (t tApp) New() *contr.App {
	c := &contr.App{}
	c.Controller = Controller.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// calls (github.com/anonx/sunplate/internal/skeleton/controllers).App.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tApp) Before(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic Before actions of embedded controllers.
	if res := Controller.Before(c.Controller, w, r); res != nil {
		return res
	}
	// Call magic Before action of (github.com/anonx/sunplate/internal/skeleton/controllers).App.
	if res := c.Before( // "Binding" parameters.
		strconv.String(r.Form, "name"),
		strconv.Ints(r.Form, "pages"),
	); res != nil {
		return res
	}
	return nil
}

// After executes magic actions of embedded controllers, and
// calls (github.com/anonx/sunplate/internal/skeleton/controllers).App.After.
func (t tApp) After(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic After actions of embedded controllers.
	if res := Controller.After(c.Controller, w, r); res != nil {
		return res
	}
	// Call magic After action of (github.com/anonx/sunplate/internal/skeleton/controllers).App.
	if res := c.After( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tApp) Initially(c *contr.App, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Execute magic Initially methods of embedded controllers.
	if finish = Controller.Initially(c.Controller, w, r); finish {
		return finish
	}
	return
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tApp) Finally(c *contr.App, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Execute magic Finally methods of embedded controllers.
	if finish = Controller.Finally(c.Controller, w, r); finish {
		return finish
	}
	return
}

// Index is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Index action found at
// github.com/anonx/sunplate/internal/skeleton/controllers/app.go
// in appropriate order.
//
// Index is an action that is used for generation of a greeting form.
func (t tApp) Index(w http.ResponseWriter, r *http.Request) {
	c := App.New()
	defer App.Finally(c, w, r)
	if finish := App.Initially(c, w, r); finish {
		return
	}
	if res := App.Before(c, w, r); res != nil {
		res.Apply(w, r)
		return
	}
	if res := c.Index( // "Binding" parameters.
	); res != nil {
		res.Apply(w, r)
		return
	}
	if res := App.After(c, w, r); res != nil {
		res.Apply(w, r)
	}
}

// PostGreet is a handler that was generated automatically.
// It calls Before, After, Finally methods, and PostGreet action found at
// github.com/anonx/sunplate/internal/skeleton/controllers/app.go
// in appropriate order.
//
// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (t tApp) PostGreet(w http.ResponseWriter, r *http.Request) {
	c := App.New()
	defer App.Finally(c, w, r)
	if finish := App.Initially(c, w, r); finish {
		return
	}
	if res := App.Before(c, w, r); res != nil {
		res.Apply(w, r)
		return
	}
	if res := c.PostGreet( // "Binding" parameters.
		strconv.String(r.Form, "name"),
	); res != nil {
		res.Apply(w, r)
		return
	}
	if res := App.After(c, w, r); res != nil {
		res.Apply(w, r)
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
