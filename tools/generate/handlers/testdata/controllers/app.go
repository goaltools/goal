package controllers

import (
	"net/http"
	"testing"
)

// App is a sample controller.
type App struct {
	*Controller
	*NotController
	*NotController1

	// Fields with incorrect types.
	a http.ResponseWriter `bind:"request"`
	b string              `bind:"response"`
	c int                 `bind:"action"`
	d int                 `bind:"controller"`
}

// NotController is not a controller as it doesn't have methods.
type NotController struct {
}

// NotController1 does have methods but it is still not a controller
// as there are not actions among those methods.
type NotController1 struct {
}

func (c NotController1) test() {
}

// HelloWorld is a sample action.
func (c App) HelloWorld(page int) (http.Handler, bool, error) {
	return nil, false, nil
}

// Initially is a magic method that is executed before every request.
func (c *Controller) Initially(w http.ResponseWriter, r *http.Request, a []string) bool {
	return false
}

// Finally is a magic method that is executed after every request.
func (c *Controller) Finally(w http.ResponseWriter, r *http.Request, a []string) bool {
	return false
}

// UnsupportedAction is not an action as it requires argument that is not
// of builtin type.
func (c Controller) UnsupportedAction(t *testing.T) http.Handler {
	return nil
}

// Smth is not an action as it returns nothing.
func (c App) Smth() {
}

func init() {
}
