// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	contr "github.com/colegion/contrib/controllers/static"

	"github.com/colegion/goal/strconv"
)

// Static is an insance of tStatic that is automatically generated from Static controller
// being found at "github.com/colegion/contrib/controllers/static/static.go",
// and contains methods to be used as handler functions.
//
// Static is a controller that brings static
// assets' serving functionality to your app.
var Static tStatic

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tStatic is a type with handler methods of Static controller.
type tStatic struct {
}

// New allocates (github.com/colegion/contrib/controllers/static).Static controller,
// then returns it.
func (t tStatic) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Static {
	c := &contr.Static{}
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tStatic) Before(c *contr.Static, w http.ResponseWriter, r *http.Request) http.Handler {

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tStatic) After(c *contr.Static, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	return
}

// Serve is a handler that was generated automatically.
// It calls Before, After methods, and Serve action found at
// github.com/colegion/contrib/controllers/static/static.go
// in appropriate order.
//
// Serve is a wrapper around Go's standard FileServer
// and StripPrefix HTTP handlers.
//@get /*filepath
func (t tStatic) Serve(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Static.New(w, r, "Static", "Serve")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Static.After(c, w, r)
	if res := Static.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Serve(
		strconv.String(r.Form, "filepath"),
	); res != nil {
		h = res
		return
	}
}

// Init initializes controllers of "github.com/colegion/contrib/controllers/static",
// its parents, and returns a list of routes along
// with handler functions associated with them.
func Init() (routes []struct {
	Method, Pattern, Label string
	Handler                http.HandlerFunc
}) {

	routes = append(routes, initStatic()...)

	return
}

func initStatic() (rs []struct {
	Method, Pattern, Label string
	Handler                http.HandlerFunc
}) {
	context.Add("Static", "Serve")
	rs = append(rs, []struct {
		Method, Pattern, Label string
		Handler                http.HandlerFunc
	}{
		{
			Method:  "GET",
			Pattern: "/*filepath",
			Label:   "",
			Handler: Static.Serve,
		},
	}...)
	return
}

func init() {
	_ = strconv.MeaningOfLife
}
