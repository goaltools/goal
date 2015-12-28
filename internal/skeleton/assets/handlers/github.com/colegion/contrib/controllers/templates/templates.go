// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/templates"

	"github.com/colegion/goal/strconv"
)

// Templates is an insance of tTemplates that is automatically generated from Templates controller
// being found at "github.com/colegion/contrib/controllers/templates/templates.go",
// and contains methods to be used as handler functions.
//
// Templates is a controller that provides support of HTML result
// rendering to your application.
// Use SetTemplatePaths to register templates and
// call c.RenderTemplate from your action to render some.
var Templates tTemplates

// tTemplates is a type with handler methods of Templates controller.
type tTemplates struct {
}

// New allocates (github.com/colegion/contrib/controllers/templates).Templates controller,
// then returns it.
func (t tTemplates) New() *contr.Templates {
	c := &contr.Templates{}
	return c
}

// Before is a dump method that always returns nil.
func (t tTemplates) Before(c *contr.Templates, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// After is a dump method that always returns nil.
func (t tTemplates) After(c *contr.Templates, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tTemplates) Initially(c *contr.Templates, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Call magic Initially method of (github.com/colegion/contrib/controllers/templates).Templates.
	return c.Initially(w, r, a)
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tTemplates) Finally(c *contr.Templates, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	return
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After, Finally methods, and RenderTemplate action found at
// github.com/colegion/contrib/controllers/templates/templates.go
// in appropriate order.
//
// RenderTemplate is an action that gets a path to template
// and renders it using data from Context.
func (t tTemplates) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "RenderTemplate"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.RenderTemplate( // "Binding" parameters.
		strconv.String(r.Form, "templatePath"),
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// Render is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Render action found at
// github.com/colegion/contrib/controllers/templates/templates.go
// in appropriate order.
//
// Render is an equivalent of the following:
//	RenderTemplate(CurrentController + "/" + CurrentAction + ".html")
// The default path pattern may be overriden by adding the following
// line to your configuration file:
//	[templates]
//	default.pattern = %s/%s.tpl
func (t tTemplates) Render(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "Render"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Render( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// Redirect is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Redirect action found at
// github.com/colegion/contrib/controllers/templates/templates.go
// in appropriate order.
//
// Redirect gets a URI or URN (e.g. "https://si.te/smt or "/users")
// and returns a handler for user's redirect using 303 status code.
func (t tTemplates) Redirect(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Templates", "Redirect"}
	defer Templates.Finally(c, w, r, a)
	if finish := Templates.Initially(c, w, r, a); finish {
		return
	}
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Redirect( // "Binding" parameters.
		strconv.String(r.Form, "urn"),
	); res != nil {
		h = res
		return
	}
	if res := Templates.After(c, w, r); res != nil {
		h = res
	}
}

// Init is used to initialize controllers of "github.com/colegion/contrib/controllers/templates"
// and its parents.
func Init() {
	initTemplates()
	contr.Init()
}

func initTemplates() {
}

func init() {
	_ = strconv.MeaningOfLife
}
