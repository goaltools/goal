package controllers

import (
	"github.com/anonx/sunplate/example/assets/views"
	"github.com/anonx/sunplate/middleware/template"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
	template.Middleware
}

func init() {
	// Define the templates that should be loaded.
	template.Paths = views.Context
}
