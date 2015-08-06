// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/anonx/sunplate/internal/skeleton/assets/handlers/github.com/anonx/sunplate/controllers/results"
	contr "github.com/anonx/sunplate/internal/skeleton/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Controller is an insance of tController that is automatically generated from Controller controller
// being found at "github.com/anonx/sunplate/internal/skeleton/controllers/init.go",
// and contains methods to be used as handler functions.
//
// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
var Controller tController

// tController is a type with handler methods of Controller controller.
type tController struct {
}

// New allocates (github.com/anonx/sunplate/internal/skeleton/controllers).Controller controller,
// initializes its parents; then returns the controller.
func (t tController) New() *contr.Controller {
	c := &contr.Controller{}
	c.Template = c0.Template.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// calls (github.com/anonx/sunplate/internal/skeleton/controllers).Controller.Before.
func (t tController) Before(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic Before actions of embedded controllers.
	if res := c0.Template.Before(c.Template, w, r); res != nil {
		if res.Finish() {
			return res
		}
	}
	// Call magic Before action of (github.com/anonx/sunplate/internal/skeleton/controllers).Controller.
	if res := c.Before( // "Binding" parameters.
	); res != nil {
		if res.Finish() {
			return res
		}
	}
	return nil
}

// After executes magic actions of embedded controllers.
func (t tController) After(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic After actions of embedded controllers.
	if res := c0.Template.After(c.Template, w, r); res != nil {
		if res.Finish() {
			return res
		}
	}
	return nil
}

// Finally executes magic actions of embedded controllers.
func (t tController) Finally(c *contr.Controller, w http.ResponseWriter, r *http.Request) {
	// Execute magic Finally actions of embedded controllers.
	c0.Template.Finally(c.Template, w, r)
}

func init() {
	_ = strconv.MeaningOfLife
}
