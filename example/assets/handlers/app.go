// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/example/controllers"

	"github.com/anonx/sunplate/strconv"
)

// App is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/example/controllers/app.go".
//
// App is a sample controller that is used for demonstration purposes.
type App struct {
}

// New allocates (github.com/anonx/sunplate/example/controllers).App controller,
// initializes its parents; then returns.
func (t App) New() *contr.App {
	c := &contr.App{}
	return c
}

// Before executes magic Before actions of parent controllers
// and calls Before action of (github.com/anonx/sunplate/example/controllers).App with arguments
// that are extracted from r.Form and converted to appropriate type.
func (t App) Before(c *contr.App, w http.ResponseWriter, r *http.Request) {
	// Execute magic Before actions of parent controllers.

	// Call magic Before action of (github.com/anonx/sunplate/example/controllers).App.
	c.Before(
		strconv.String(r.Form, "name"),
		strconv.Ints(r.Form, "pages"),
	)
}

func init() {
	_ = strconv.MeaningOfLife
}
