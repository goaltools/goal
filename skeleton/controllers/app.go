package controllers

import (
	v "github.com/anonx/sunplate/skeleton/assets/views"

	"github.com/anonx/sunplate/action"
)

// App is a sample controller that is used for demonstration purposes.
type App struct {
	*Controller
}

// Before is a magic method that is executed before every request.
func (c *App) Before(name string, pages []int) action.Result {
	return nil
}

// Index is an action that is used for generation of a greeting form.
func (c *App) Index() action.Result {
	return c.RenderTemplate(v.Paths.App.IndexHTML)
}

// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (c *App) PostGreet(name string) action.Result {
	c.Context["name"] = name
	return c.RenderTemplate(v.Paths.App.GreetHTML)
}

// After is a magic method that is executed after every request.
func (c *App) After() action.Result {
	return nil
}

// Finally is a magic method that is executed after every request
// no matter what.
func (c *App) Finally(name string) action.Result {
	return nil
}
