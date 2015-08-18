package controllers

import (
	"net/http"

	v "github.com/anonx/sunplate/internal/skeleton/assets/views"
)

// App is a sample controller that is used for demonstration purposes.
type App struct {
	*Controllers
}

// Before is a magic method that is executed before every request.
func (c *App) Before(name string, pages []int) http.Handler {
	return nil
}

// Index is an action that is used for generation of a greeting form.
func (c *App) Index() http.Handler {
	return c.RenderTemplate(v.Paths.App.IndexHTML)
}

// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (c *App) PostGreet(name string) http.Handler {
	c.Context["name"] = name
	c.Context["message"] = c.Request.FormValue("message")
	return c.RenderTemplate(v.Paths.App.GreetHTML)
}

// After is a magic method that is executed after every request.
func (c *App) After() http.Handler {
	return nil
}
