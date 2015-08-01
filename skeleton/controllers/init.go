package controllers

import (
	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/controllers/results"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
	*results.Template
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controller) Before() a.Result {
	return nil
}
