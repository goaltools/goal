package controllers

import (
	"net/http"

	"github.com/goaltools/contrib/controllers/requests"
	"github.com/goaltools/contrib/controllers/sessions"
	"github.com/goaltools/contrib/controllers/static"
	"github.com/goaltools/contrib/controllers/templates"
)

// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
type Controllers struct {
	*requests.Requests
	*templates.Templates
	*sessions.Sessions
	*static.Static `@get:"/"`
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controllers) Before() http.Handler {
	return nil
}
