// Package templates provides abstractions for work
// with standard Go template engine.
package templates

import (
	"html/template"
	"net/http"
)

var (
	// BaseTemplate is a name of the file that will be loaded
	// with every template to make extends pattern possible.
	// So, if you have the following structure:
	//	./base.html
	//	./home.html
	//	./profile/base.html
	//	./profile/index.html
	// You will get pairs of (base.html, home.html) and
	// (profile/base.html, profile/index.html).
	// If no base template is found in ./profile/ directory,
	// one in a previous level (./) will be used.
	BaseTemplate = "Base.html"

	// TemplateName is name of the template that will be executed.
	// By-default, your base.html should have {%define "base"%}
	// that will be the entry point of every of your templates.
	TemplateName = "base"

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

// Templates is a controller that provides support of HTML result
// rendering to your application.
// Use SetTemplatePaths to register templates and
// call c.RenderTemplate from your action to render some.
type Templates struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}

	// Status is a status code that will be returned when rendering.
	Status int
}

// Before initializes Context that will be passed to template.
func (c *Templates) Before() http.Handler {
	c.Context = map[string]interface{}{}
	return nil
}

// RenderTemplate is an action that gets a path to template
// and renders it using data from Context.
func (c *Templates) RenderTemplate(templatePath string) http.Handler {
	return &Handler{
		context:  c.Context,
		template: templatePath,
	}
}

// RenderNotFound is an action that renders Error 404 page.
func (c *Templates) RenderNotFound() http.Handler {
	c.Status = http.StatusNotFound
	return c.RenderTemplate("Errors/NotFound.html")
}

func init() {
	// Use {% and %} instead of {{ and }} as default delimiters.
	if Delims.Left == "" || Delims.Right == "" {
		Delims.Left = "{%"
		Delims.Right = "%}"
	}
}
