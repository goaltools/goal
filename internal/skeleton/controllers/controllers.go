package controllers

import (
	"net/http"

	"github.com/anonx/sunplate/internal/skeleton/assets/views"

	"github.com/anonx/sunplate/controllers/requests"
	"github.com/anonx/sunplate/controllers/sessions"
	"github.com/anonx/sunplate/controllers/templates"
)

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

// The line below tells golang's generate command you want
// it to generate a list of templates found in your ../views directory.
// Please, do not delete it unless you know what you are doing.
//
//go:generate sunplate generate views --input ../views --output ../assets/views

// Init gets a path to the root of this project
// and initializes parent controllers.
func Init(root string) {
	// Loading templates from the generated list.
	templates.Load(root, "views/", views.List)
}
