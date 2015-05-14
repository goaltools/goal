// Package template provides functions for work
// with standard Go template engine.
// It should be embeded into Controller struct.
package template

import (
	"html/template"

	"github.com/anonx/sunplate/action"
)

var (
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
	Paths = map[string]string{}

	// Delims are action delimiters that are used for call to Parse.
	// Empty delimiters activate default: {{ and }}.
	Delims struct {
		Left, Right string
	}

	// Funcs are added to the template's function map.
	// Functions are expected to return just 1 argument or
	// 2 in case the second one is of error type.
	Funcs template.FuncMap
)

// Middleware is a main type that should be embeded into controller structs.
type Middleware struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}
}

// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t *Middleware) RenderTemplate(templatePath string) action.Result {
	return &HTML{}
}
