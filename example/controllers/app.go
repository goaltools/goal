package controllers

// App is a sample controller that is used for demonstration purposes.
type App struct {
	Controller
}

// Before is a magic method that is executed before every request.
func (c *App) Before() {
}

// Index is an action that is used for generation of a greeting form.
func (c *App) Index() {
}

// PostGreet prints received user fullname. If it is not valid,
// user is redirected back to index page.
func (c *App) PostGreet(name string) {
}

// After is a magic method that is executed after every request.
func (c *App) After() {
}

// Finally is a magic method that is executed after every request
// no matter what.
func (c *App) Finally() {
}

// Init is a system method that will be called once during application's startup.
func (c *App) Init() {
}
