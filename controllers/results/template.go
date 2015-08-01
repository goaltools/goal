// Package results provides abstractions for work
// with standard Go template engine.
package results

import (
	"html/template"

	"github.com/anonx/sunplate/action"
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

// Template is a main type that should be embeded into controller structs.
type Template struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}
}

// Before initializes Context that will be passed to template.
func (t *Template) Before() action.Result {
	t.Context = map[string]interface{}{}
	return nil
}

// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t *Template) RenderTemplate(templatePath string) action.Result {
	return &HTML{
		context:  t.Context,
		template: templatePath,
	}
}

func init() {
	// Use {% and %} instead of {{ and }} as default delimiters.
	if Delims.Left == "" || Delims.Right == "" {
		Delims.Left = "{%"
		Delims.Right = "%}"
	}
}
