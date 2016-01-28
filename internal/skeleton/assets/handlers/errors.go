// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/strconv"
)

// Errors is an insance of tErrors that is automatically generated from Errors controller
// being found at "github.com/colegion/goal/internal/skeleton/controllers/errors.go",
// and contains methods to be used as handler functions.
//
// Errors is a controller with actions displaying error pages.
var Errors tErrors

// tErrors is a type with handler methods of Errors controller.
type tErrors struct {
}

// New allocates (github.com/colegion/goal/internal/skeleton/controllers).Errors controller,
// initializes its parents; then returns the controller.
func (t tErrors) New() *contr.Errors {
	c := &contr.Errors{}
	c.Controllers = Controllers.New()
	return c
}

// Before executes magic actions of embedded controllers.
func (t tErrors) Before(c *contr.Errors, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if res := Controllers.Before(c.Controllers, w, r); res != nil {
		return res
	}
	return nil
}

// After executes magic actions of embedded controllers.
func (t tErrors) After(c *contr.Errors, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic After actions of embedded controllers.
	if res := Controllers.After(c.Controllers, w, r); res != nil {
		return res
	}
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tErrors) Initially(c *contr.Errors, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Execute magic Initially methods of embedded controllers.
	if finish = Controllers.Initially(c.Controllers, w, r, a); finish {
		return finish
	}
	return
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tErrors) Finally(c *contr.Errors, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Execute magic Finally methods of embedded controllers.
	if finish = Controllers.Finally(c.Controllers, w, r, a); finish {
		return finish
	}
	return
}

// NotFound is a handler that was generated automatically.
// It calls Before, After, Finally methods, and NotFound action found at
// github.com/colegion/goal/internal/skeleton/controllers/errors.go
// in appropriate order.
//
// NotFound prints an error 404 "Page Not Found" message.
func (t tErrors) NotFound(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Errors", "NotFound"}
	defer Errors.Finally(c, w, r, a)
	if finish := Errors.Initially(c, w, r, a); finish {
		return
	}
	if res := Errors.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.NotFound( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := Errors.After(c, w, r); res != nil {
		h = res
	}
}

// InternalError is a handler that was generated automatically.
// It calls Before, After, Finally methods, and InternalError action found at
// github.com/colegion/goal/internal/skeleton/controllers/errors.go
// in appropriate order.
//
// InternalError displays an error 500 "Internal Server Error" message.
func (t tErrors) InternalError(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.New()
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	a := []string{"Errors", "InternalError"}
	defer Errors.Finally(c, w, r, a)
	if finish := Errors.Initially(c, w, r, a); finish {
		return
	}
	if res := Errors.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.InternalError( // "Binding" parameters.
	); res != nil {
		h = res
		return
	}
	if res := Errors.After(c, w, r); res != nil {
		h = res
	}
}

func initErrors() {
	context.Add("Errors", "NotFound")
	context.Add("Errors", "InternalError")
}

func init() {
	_ = strconv.MeaningOfLife
}
