// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/anonx/sunplate/internal/skeleton/assets/handlers/github.com/anonx/sunplate/controllers/requests"
	c1 "github.com/anonx/sunplate/internal/skeleton/assets/handlers/github.com/anonx/sunplate/controllers/results"
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
	c.Params = c0.Params.New()
	c.Template = c1.Template.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// calls (github.com/anonx/sunplate/internal/skeleton/controllers).Controller.Before.
func (t tController) Before(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic Before actions of embedded controllers.
	if res := c0.Params.Before(c.Params, w, r); res != nil {
		return res
	}
	if res := c1.Template.Before(c.Template, w, r); res != nil {
		return res
	}
	// Call magic Before action of (github.com/anonx/sunplate/internal/skeleton/controllers).Controller.
	if res := c.Before( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// After executes magic actions of embedded controllers.
func (t tController) After(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Execute magic After actions of embedded controllers.
	if res := c0.Params.After(c.Params, w, r); res != nil {
		return res
	}
	if res := c1.Template.After(c.Template, w, r); res != nil {
		return res
	}
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tController) Initially(c *contr.Controller, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Execute magic Initially methods of embedded controllers.
	if finish = c0.Params.Initially(c.Params, w, r); finish {
		return finish
	}
	if finish = c1.Template.Initially(c.Template, w, r); finish {
		return finish
	}
	return
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tController) Finally(c *contr.Controller, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Execute magic Finally methods of embedded controllers.
	if finish = c0.Params.Finally(c.Params, w, r); finish {
		return finish
	}
	if finish = c1.Template.Finally(c.Template, w, r); finish {
		return finish
	}
	return
}

func init() {
	_ = strconv.MeaningOfLife
}
