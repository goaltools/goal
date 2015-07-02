package controllers

import (
	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/generation/handlers/testdata/controllers/subpackage"

	"github.com/revel/revel/testing"
)

// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middleware controllers available.
type Controller struct {
	*subpackage.Controller

	testing.TestSuite
	Test testing.TestSuite
}

// Before is a magic method that is executed before every request.
func (c *Controller) Before(uid string) a.Result {
	return nil
}

// index is not an action as this method is not public.
func (c Controller) index(page int) a.Result {
	return nil
}

// Index is a sample action.
func (c *App) Index(page int) a.Result {
	return nil
}

// NotAction is not an action as this method doesn't return
// action.Result as its first argument.
func (c Controller) NotAction(page int) (bool, a.Result) {
	return false, nil
}

// After is a magic method that is executed after every request.
func (c *Controller) After(name string) a.Result {
	return nil
}
