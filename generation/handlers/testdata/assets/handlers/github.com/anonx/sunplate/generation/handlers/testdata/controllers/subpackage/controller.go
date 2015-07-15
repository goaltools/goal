// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// Controller is an instance of tController.
var Controller tController

// tController is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage/app.go".
//
// Controller is some controller.
type Controller struct {
}

// New allocates (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller controller,
// then returns it.
func (t tController) New() *contr.Controller {
	c := &contr.Controller{}
	return c
}

// Before calls (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tController) Before(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic Before action of (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller.
	if res := c.Before( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// After calls (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller.After with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tController) After(c *contr.Controller, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic After action of (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller.
	if res := c.After( // "Binding" parameters.
	); res != nil {
		return res
	}
	return nil
}

// Finally calls (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller.Finally with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t tController) Finally(c *contr.Controller, w http.ResponseWriter, r *http.Request) {
	// Call magic Finally action of (github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage).Controller.
	defer func() {
		if res := c.Finally( // "Binding" parameters.
			strconv.String(r.Form, "userID"),
		); res != nil {
			res.Apply(w, r)
		}
	}()
}

// Index is a handler that was generated automatically.
// It calls Before, After, Finally methods, and Index action found at
// github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage/app.go
// in appropriate order.
//
// Index is a sample action.
func (t Controller) Index(w http.ResponseWriter, r *http.Request) {
	c := Controller.New()
	defer Controller.Finally(c, w, r)
	if res := (Controller.Before(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := c.Index( // "Binding" parameters.
		strconv.Int(r.Form, "page"),
	); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
	if res := (Controller.After(c, w, r)); res != nil {
		res.Apply(w, r)
		if res.Finish() {
			return
		}
	}
}

func init() {
	_ = strconv.MeaningOfLife
}