// Package template provides functions for work
// with standard Go template engine.
// It should be embeded into Controller struct.
package template

import (
	"github.com/anonx/sunplate/action"
)

// Paths is a map of templates.
// Template names are represented as keys and their paths as values.
// Initialize it from your init.go as follows:
//	import (
//		"github.com/anonx/sunplate/assets/views"
//		"github.com/anonx/sunplate/middleware/template"
//	)
//
//	type Controller struct {
//		// ...
//		template.Middleware
//	}
//
//	func init() {
//		template.Paths = views.Context
//	}
var Paths = map[string]string{}

// Middleware is a main type that should be embeded into controller structs.
type Middleware struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}
}

// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t *Middleware) RenderTemplate(templatePath string) action.Result {
	return &HTML{}
}
