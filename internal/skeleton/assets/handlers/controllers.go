// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/colegion/goal/internal/skeleton/assets/handlers/github.com/colegion/contrib/controllers/requests"
	c2 "github.com/colegion/goal/internal/skeleton/assets/handlers/github.com/colegion/contrib/controllers/sessions"
	c1 "github.com/colegion/goal/internal/skeleton/assets/handlers/github.com/colegion/contrib/controllers/templates"
	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/strconv"
)

// Controllers is an insance of tControllers that is automatically generated from Controllers controller
// being found at "github.com/colegion/goal/internal/skeleton/controllers/init.go",
// and contains methods to be used as handler functions.
//
// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
var Controllers tControllers

// tControllers is a type with handler methods of Controllers controller.
type tControllers struct {
}

// New allocates (github.com/colegion/goal/internal/skeleton/controllers).Controllers controller,
// initializes its parents; then returns the controller.
func (t tControllers) New() *contr.Controllers {
	c := &contr.Controllers{}
	c.Requests = c0.Requests.New()
	c.Templates = c1.Templates.New()
	c.Sessions = c2.Sessions.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// calls (github.com/colegion/goal/internal/skeleton/controllers).Controllers.Before.
func (t tControllers) Before(c *contr.Controllers, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if res := c0.Requests.Before(c.Requests, w, r); res != nil {
		return res
	}

	if res := c1.Templates.Before(c.Templates, w, r); res != nil {
		return res
	}

	if res := c2.Sessions.Before(c.Sessions, w, r); res != nil {
		return res
	}

	// Call magic Before action of (github.com/colegion/goal/internal/skeleton/controllers).Controllers.
	if res := c.Before(); res != nil {
		return res
	}

	return nil
}

// After executes magic actions of embedded controllers.
func (t tControllers) After(c *contr.Controllers, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic After actions of embedded controllers.
	if res := c0.Requests.After(c.Requests, w, r); res != nil {
		return res
	}

	if res := c1.Templates.After(c.Templates, w, r); res != nil {
		return res
	}

	if res := c2.Sessions.After(c.Sessions, w, r); res != nil {
		return res
	}

	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tControllers) Initially(c *contr.Controllers, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {

	// Execute magic Initially methods of embedded controllers.

	if finish = c0.Requests.Initially(c.Requests, w, r, a); finish {
		return finish
	}

	if finish = c1.Templates.Initially(c.Templates, w, r, a); finish {
		return finish
	}

	if finish = c2.Sessions.Initially(c.Sessions, w, r, a); finish {
		return finish
	}

	return

}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tControllers) Finally(c *contr.Controllers, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {

	// Execute magic Finally methods of embedded controllers.

	if finish = c0.Requests.Finally(c.Requests, w, r, a); finish {
		return finish
	}

	if finish = c1.Templates.Finally(c.Templates, w, r, a); finish {
		return finish
	}

	if finish = c2.Sessions.Finally(c.Sessions, w, r, a); finish {
		return finish
	}

	return
}

func initControllers() {

	c0.Init()

	c1.Init()

	c2.Init()

}

func init() {
	_ = strconv.MeaningOfLife
}
