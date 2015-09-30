// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/config"
	"github.com/colegion/goal/strconv"
)

// App is an insance of tApp that is automatically generated from App controller
// being found at "github.com/colegion/goal/internal/skeleton/controllers/app.go",
// and contains methods to be used as handler functions.
//
// App is a sample controller that is used for demonstration purposes.
var App tApp

// tApp is a type with handler methods of App controller.
type tApp struct {
}

// New allocates (github.com/colegion/goal/internal/skeleton/controllers).App controller,
// initializes its parents; then returns the controller.
func (t tApp) New() *contr.App {
	c := &contr.App{}
	c.Controllers = Controllers.New()
	return c
}

// Before executes magic actions of embedded controllers, and
// calls (github.com/colegion/goal/internal/skeleton/controllers).App.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tApp) Before(c *contr.App, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if res := Controllers.Before(c.Controllers, w, r); res != nil {
		return res
	}
	// Call magic Before action of (github.com/colegion/goal/internal/skeleton/controllers).App.
	if res := c.Before( // "Binding" parameters.
		strconv.String(r.Form, "name"),
		strconv.Ints(r.Form, "pages"),
	); res != nil {
		return res
	}
	return nil
}

// After executes magic actions of embedded controllers.
func (t tApp) After(c *contr.App, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic After actions of embedded controllers.
	if res := Controllers.After(c.Controllers, w, r); res != nil {
		return res
	}
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tApp) Initially(c *contr.App, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Execute magic Initially methods of embedded controllers.
	if finish = Controllers.Initially(c.Controllers, w, r); finish {
		return finish
	}
	return
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tApp) Finally(c *contr.App, w http.ResponseWriter, r *http.Request) (finish bool) {
	// Execute magic Finally methods of embedded controllers.
	if finish = Controllers.Finally(c.Controllers, w, r); finish {
		return finish
	}
	return
}

// Index is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Index action found at
// github.com/colegion/goal/internal/skeleton/controllers/app.go
// in appropriate order.
//
// Index is an action that is used for generation of a greeting form.
func (t tApp) Index(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := App.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer App.Finally(c, w, r)
	if finish := App.Initially(c, w, r); finish {
		return
	}
	if res := App.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Index( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := App.After(c, w, r); res != nil {
		h = res
	}
}

// PostGreet is a handler that was generated automatically.
// It calls Before, After, Finally methods, and PostGreet action found at
// github.com/colegion/goal/internal/skeleton/controllers/app.go
// in appropriate order.
//
// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (t tApp) PostGreet(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := App.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer App.Finally(c, w, r)
	if finish := App.Initially(c, w, r); finish {
		return
	}
	if res := App.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.PostGreet( // "Binding" parameters.
		strconv.String(r.Form, "name"),
	); res != nil {
		h = res
		return
	}
	if res := App.After(c, w, r); res != nil {
		h = res
	}
}

// Init is used to initialize controllers of "github.com/colegion/goal/internal/skeleton/controllers"
// and its parents.
func Init(g config.Getter) {
	initApp(g)
	initControllers(g)
}

func initApp(g config.Getter) {
}

func init() {
	_ = strconv.MeaningOfLife
}
