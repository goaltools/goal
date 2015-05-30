package controllers

import (
	"github.com/anonx/sunplate/example/assets/views"
	"github.com/anonx/sunplate/middleware/template"

	t "github.com/revel/revel/testing"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
	template.Middleware

	*t.TestSuite

	HeyWorld   TestType `tag:"smth_cool=xxx"`
	Bullshit   *string
	GPA, Grade float64
}

// TestType ...
type TestType struct {
}

// Yahooooooo is cool...
func (c *Controller) Yahooooooo(ctx1 *Controller, ctx2 Controller, name, lastname string, age int) bool {
	return true
}

func init() {
	// Define the templates that should be loaded.
	template.Paths = views.Context
}
