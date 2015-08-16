package controllers

import (
	"net/http"

	"github.com/anonx/sunplate/controllers/requests"
	"github.com/anonx/sunplate/controllers/templates"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
	*requests.Requests
	*templates.Templates
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controller) Before() http.Handler {
	return nil
}
