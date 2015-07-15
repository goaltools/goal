// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/anonx/sunplate/generation/handlers/testdata/assets/handlers/github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage"
	contr "github.com/anonx/sunplate/generation/handlers/testdata/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Controller is an instance of tController.
var Controller tController

// tController is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/generation/handlers/testdata/controllers/init.go".
//
// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middleware controllers available.
type Controller struct {
}

// New allocates (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller controller,
// initializes its parents; then returns the controller.
func (t tController) New() *contr.Controller {
	c := &contr.Controller{}
	c.Controller = c0.Controller.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// Before calls (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tController) Before(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic Before action of (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller.
	if res := c.Before( // "Binding" parameters.
		strconv.String(r.Form, "uid"),
	); res != nil {
		return res
	}
	return nil
}

// After executes magic actions of embedded controllers, and
// After calls (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller.After with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tController) After(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic After action of (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller.
	if res := c.After( // "Binding" parameters.
		strconv.String(r.Form, "name"),
	); res != nil {
		return res
	}
	return nil
}

// Finally executes magic actions of embedded controllers, and
// Finally calls (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller.Finally with arguments
// that are extracted from r.Form and converted to appropriate types.
// The call is garanteed to be done no matter what happens during execution of parent actions.
func (t tController) Finally(c *contr.Controller, w http.ResponseWriter, r *http.Request) {
	// Call magic Finally action of (github.com/anonx/sunplate/generation/handlers/testdata/controllers).Controller.
	defer func() {
		if res := c.Finally( // "Binding" parameters.
			strconv.String(r.Form, "name"),
		); res != nil {
			res.Apply(w, r)
		}
	}()
	// Execute magic Finally actions of embedded controllers.
	c0.Controller.Finally(c.Controller, w, r)
}

func init() {
	_ = strconv.MeaningOfLife
}