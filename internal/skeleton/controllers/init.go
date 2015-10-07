package controllers

import (
	"net/http"

	"github.com/colegion/contrib/controllers/requests"
	"github.com/colegion/contrib/controllers/sessions"
	"github.com/colegion/contrib/controllers/templates"
)

// The line below tells golang's generate command you want
// it to scan your controllers and generate handler functions
// from them using rules of goal toolkit.
// Please, do not delete it unless you know what you are doing.
//
//go:generate goal generate handlers --input ./ --output ../assets/handlers

// The line below tells golang's generate command you want
// it to scan your views and generate a listing.
// Please, do not delete it unless you know what you are doing.
//
//go:generate goal generate listing --input ../views --output ../assets/views

// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
type Controllers struct {
	*requests.Requests
	*templates.Templates
	*sessions.Sessions
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controllers) Before() http.Handler {
	return nil
}
