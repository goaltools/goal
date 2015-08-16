package controllers

import (
	"net/http"

	"github.com/anonx/sunplate/controllers/requests"
	"github.com/anonx/sunplate/controllers/templates"
)

// The line below tells golang's generate command you want
// it to scan your controllers and generate handler functions
// from them using rules of sunplate toolkit.
// Please, do not delete it unless you know what you are doing.
//
//go:generate sunplate generate handlers --input ./ --output ../assets/handlers

// Controller is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
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
