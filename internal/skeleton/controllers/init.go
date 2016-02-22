package controllers

import (
	"net/http"

	"github.com/colegion/contrib/controllers/errors"
	"github.com/colegion/contrib/controllers/global"
	"github.com/colegion/contrib/controllers/requests"
	"github.com/colegion/contrib/controllers/sessions"
	"github.com/colegion/contrib/controllers/static"
	"github.com/colegion/contrib/controllers/templates"
)

// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
type Controllers struct {
	*global.Global
	*requests.Requests
	*sessions.Sessions

	*static.Static `@get:"/"`
	*errors.Errors `@route:"/errors"`

	// Templates controller is responsible for rendering templates.
	// It must be located after all other controllers that may use
	// template rendering. When trying to render a template that doesn't
	// exist, a panic is thrown that must be caught by router.
	*templates.Templates
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controllers) Before() http.Handler {
	return nil
}
