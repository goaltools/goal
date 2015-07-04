// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	c0 "github.com/anonx/sunplate/example/assets/handlers/github.com/anonx/sunplate/middleware/template"
	contr "github.com/anonx/sunplate/example/controllers"

	"github.com/anonx/sunplate/strconv"
)

// Controller is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/example/controllers/init.go".
//
// Controller is a struct that should be embedded into every controller
// of your app to make methods provided by middlewares available.
type Controller struct {
}

// New allocates (github.com/anonx/sunplate/example/controllers).Controller controller,
// initializes its parents; then returns.
func (t Controller) New() *contr.Controller {
	c := &contr.Controller{}
	c.Middleware = c0.Middleware{}.New()
	return c
}

// Before executes magic Before actions of parent controllers
// and calls Before action of (github.com/anonx/sunplate/example/controllers).Controller with arguments
// that are extracted from r.Form and converted to appropriate type.
func (t Controller) Before(c *contr.Controller, w http.ResponseWriter, r *http.Request) {
	// Execute magic Before actions of parent controllers.
	c0.Middleware{}.Before(c.Middleware, w, r)

	// Call magic Before action of (github.com/anonx/sunplate/example/controllers).Controller.
	c.Before()
}

func init() {
	_ = strconv.MeaningOfLife
}
