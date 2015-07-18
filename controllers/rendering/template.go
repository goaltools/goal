// Package rendering provides abstractions for work
// with standard Go template engine.
package rendering

import (
	"html/template"

	"github.com/anonx/sunplate/action"
)

var (
	// TemplatePaths is a map of templates.
	// Template names are represented as keys and their paths as values.
	// Initialize it from your init.go as follows:
	//	import (
	//		"github.com/user/project/assets/views"
	//		"github.com/anonx/sunplate/controllers/rendering"
	//	)
	//
	//	type Controller struct {
	//		// ...
	//		rendering.Template
	//	}
	//
	//	func init() {
	//		rendering.TemplatePaths = views.Context
	//	}
	TemplatePaths = map[string]string{}

	// Delims are action delimiters that are used for call to Parse.
	// Empty delimiters activate default: {% and %}.
	Delims struct {
		Left, Right string
	}

	// Funcs are added to the template's function map.
	// Functions are expected to return just 1 argument or
	// 2 in case the second one is of error type.
	Funcs template.FuncMap
)

// Template is a main type that should be embeded into controller structs.
type Template struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}
}

// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t *Template) RenderTemplate(templatePath string) action.Result {
	return &HTML{}
}

func init() {
	// Use {% and %} instead of {{ and }} as default delimiters.
	if Delims.Left == "" || Delims.Right == "" {
		Delims.Left = "{%"
		Delims.Right = "%}"
	}
}
