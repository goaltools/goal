package controllers

import (
	v "github.com/anonx/sunplate/skeleton/assets/views"

	a "github.com/anonx/sunplate/action"
)

// App is a sample controller that is used for demonstration purposes.
type App struct {
	*Controller
}

// Before is a magic method that is executed before every request.
func (c *App) Before(name string, pages []int) a.Result {
	return nil
}

// Index is an action that is used for generation of a greeting form.
func (c *App) Index() a.Result {
	return c.RenderTemplate(v.Paths.App.IndexHTML)
}

// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (c *App) PostGreet(name string) a.Result {
	c.Context["name"] = name
	return c.RenderTemplate(v.Paths.App.GreetHTML)
}

// After is a magic method that is executed after every request.
func (c *App) After() a.Result {
	return nil
}

// Finally is a magic method that is executed after every request
// no matter what.
func (c *App) Finally(name string) a.Result {
	return nil
}
