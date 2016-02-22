// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/colegion/contrib/controllers/global"
	c1 "github.com/colegion/contrib/controllers/requests"
	contr "github.com/colegion/contrib/controllers/templates"

	"github.com/colegion/goal/strconv"
)

// Templates is an instance of tTemplates that is automatically generated from
// Templates controller being found at "github.com/colegion/contrib/controllers/templates/templates.go",
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

// newC allocates (github.com/colegion/contrib/controllers/templates).Templates controller,
// initializes its parents and returns it.
func (t tTemplates) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Templates {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Templates{}

	// Allocate its parents. Make sure controller of every type
	// is allocated just once, then reused.
	c.Requests = &c1.Requests{

		Request: r,

		Response: w,
	}
	c.Global = &c0.Global{

		CurrentAction: act,

		CurrentController: ctr,
	}

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tTemplates) before(c *contr.Templates, w http.ResponseWriter, r *http.Request) http.Handler {
	// Call special Before actions of the parent controllers.
	if res := c.Global.Before(); res != nil {
		return res
	}
	if res := c.Requests.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tTemplates) after(c *contr.Templates, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	// Execute magic After methods of embedded controllers.

	return
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, RenderTemplate, and After actions of a <no value> controller.
// RenderTemplate action may be found in the "github.com/colegion/contrib/controllers/templates/templates.go".
//
// RenderTemplate is an action that gets a path to template
// and returns an HTTP handler for its rendering.
func (t tTemplates) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.newC(w, r, "Templates", "RenderTemplate")
	defer func() {
		// If one of the actions (Before, After or RenderTemplate) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Templates.after(c, w, r) // Call this at the very end, but before applying result.
	if res := Templates.before(c, w, r); res != nil {
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
// It calls Before, Render, and After actions of a <no value> controller.
// Render action may be found in the "github.com/colegion/contrib/controllers/templates/templates.go".
//
// Render executes a te
func (t tTemplates) Render(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Templates.newC(w, r, "Templates", "Render")
	defer func() {
		// If one of the actions (Before, After or Render) returned
		// a handler, apply it.
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Templates.after(c, w, r) // Call this at the very end, but before applying result.
	if res := Templates.before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Render(); res != nil {
		h = res
		return
	}
}

func init() {
	_ = strconv.MeaningOfLife
}
