package controllers

import (
	"github.com/anonx/sunplate/skeleton/assets/views"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/controllers/rendering"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
	*rendering.Template
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controller) Before() a.Result {
	return nil
}

// The line below tells golang's generate command you want
// it to generate a list of views (views.Context) for you.
// Please, do not delete it unless you know what you are doing.
//
//go:generate sunplate generate listing --input ../views --output ../assets/views

func init() {
	// Define the templates that should be loaded.
	rendering.SetTemplatePaths(views.Context)
}
