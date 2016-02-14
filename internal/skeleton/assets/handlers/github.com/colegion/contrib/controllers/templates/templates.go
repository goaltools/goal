// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

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

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tTemplates is a type with handler methods of Templates controller.
type tTemplates struct {
}

// New allocates (github.com/colegion/contrib/controllers/templates).Templates controller,
// then returns it.
func (t tTemplates) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Templates {
	c := &contr.Templates{

		Action: act,

		Controller: ctr,
	}
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tTemplates) Before(c *contr.Templates, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call magic Before action of (github.com/colegion/contrib/controllers/templates).Before.
	if h := c.Before(); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tTemplates) After(c *contr.Templates, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	return
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After methods, and RenderTemplate action found at
// github.com/colegion/contrib/controllers/templates/templates.go
// in appropriate order.
//
// RenderTemplate is an action that gets a path to template
// and renders it using data from Context.
func (t tTemplates) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New(w, r, "Templates", "RenderTemplate")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Templates.After(c, w, r)
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.RenderTemplate(
		strconv.String(r.Form, "templatePath"),
	); res != nil {
		h = res
		return
	}
}

// Render is a handler that was generated automatically.
// It calls Before, After methods, and Render action found at
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
	c := Templates.New(w, r, "Templates", "Render")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Templates.After(c, w, r)
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Render(); res != nil {
		h = res
		return
	}
}

// Redirect is a handler that was generated automatically.
// It calls Before, After methods, and Redirect action found at
// github.com/colegion/contrib/controllers/templates/templates.go
// in appropriate order.
//
// Redirect gets a URI or URN (e.g. "https://si.te/smt or "/users")
// and returns a handler for user's redirect using 303 status code.
func (t tTemplates) Redirect(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.New(w, r, "Templates", "Redirect")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Templates.After(c, w, r)
	if res := Templates.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Redirect(
		strconv.String(r.Form, "urn"),
	); res != nil {
		h = res
		return
	}
}

// Init initializes controllers of "github.com/colegion/contrib/controllers/templates",
// its parents, and returns a list of routes along
// with handler functions associated with them.
func Init() (routes []struct {
	Method, Pattern string
	Handler         http.HandlerFunc
}) {

	routes = append(routes, initTemplates()...)

	contr.Init(context)

	return
}

func initTemplates() (rs []struct {
	Method, Pattern string
	Handler         http.HandlerFunc
}) {
	context.Add("Templates", "RenderTemplate")

	context.Add("Templates", "Render")

	context.Add("Templates", "Redirect")

	return
}

func init() {
	_ = strconv.MeaningOfLife
}
