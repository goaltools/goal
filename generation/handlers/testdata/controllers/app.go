package controllers

import (
	"github.com/anonx/sunplate/action"
)

// App is a sample controller.
type App struct {
	Controller
}

// HelloWorld is a sample action.
func (c App) HelloWorld(page int) action.Result {
	return nil
}

// Finally is a magic method that is executed after every request.
func (c *Controller) Finally(name string) action.Result {
	return nil
}

// Smth is not an action as it returns nothing.
func (c App) Smth() {
}

func init() {
}
