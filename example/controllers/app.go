package controllers

import (
	"github.com/anonx/ok/action"
)

// App is a sample controller that is used for demonstration purposes.
type App struct {
	Controller
}

// Before is a magic method that is executed before every request.
func (c *App) Before() action.Result {
	return nil
}

// Index is an action that is used for generation of a greeting form.
func (c *App) Index() action.Result {
	return c.RenderTemplate("test.html")
}

// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (c *App) PostGreet(name string) action.Result {
	return nil
}

// After is a magic method that is executed after every request.
func (c *App) After() action.Result {
	return nil
}

// Finally is a magic method that is executed after every request
// no matter what.
func (c *App) Finally() action.Result {
	return nil
}

// Init is a system method that will be called once during application's startup.
func (c *App) Init() {
}
