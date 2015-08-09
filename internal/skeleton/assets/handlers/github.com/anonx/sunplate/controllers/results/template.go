// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/controllers/results"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Template is an insance of tTemplate that is automatically generated from Template controller
// being found at "github.com/anonx/sunplate/controllers/results/template.go",
// and contains methods to be used as handler functions.
//
// Template is a main type that should be embeded into controller structs.
var Template tTemplate

// tTemplate is a type with handler methods of Template controller.
type tTemplate struct {
}

// New allocates (github.com/anonx/sunplate/controllers/results).Template controller,
// then returns it.
func (t tTemplate) New() *contr.Template {
	c := &contr.Template{}
	return c
}

// Before calls (github.com/anonx/sunplate/controllers/results).Template.Before.
func (t tTemplate) Before(c *contr.Template, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic Before action of (github.com/anonx/sunplate/controllers/results).Template.
	if res := c.Before( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// After is a dump method that always returns nil.
func (t tTemplate) After(c *contr.Template, w http.ResponseWriter, r *http.Request) a.Result {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tTemplate) Initially(c *contr.Template, w http.ResponseWriter, r *http.Request) (finish bool) {
	return
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tTemplate) Finally(c *contr.Template, w http.ResponseWriter, r *http.Request) (finish bool) {
	return
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderTemplate action found at
// github.com/anonx/sunplate/controllers/results/template.go
// in appropriate order.
//
// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t tTemplate) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	c := Template.New()
	defer Template.Finally(c, w, r)
	if finish := Template.Initially(c, w, r); finish {
		return
	}
	if res := Template.Before(c, w, r); res != nil {
		res.Apply(w, r)
		return
	}
	if res := c.RenderTemplate( // "Binding" parameters.
		strconv.String(r.Form, "templatePath"),
	); res != nil {
		res.Apply(w, r)
		return
	}
	if res := Template.After(c, w, r); res != nil {
		res.Apply(w, r)
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
