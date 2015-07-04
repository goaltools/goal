// Package handlers is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/anonx/sunplate/example/controllers"

	a "github.com/anonx/sunplate/action"
	"github.com/anonx/sunplate/strconv"
)

// App is automatically generated from a controller
// that was found at "github.com/anonx/sunplate/example/controllers/app.go".
//
// App is a sample controller that is used for demonstration purposes.
type App struct {
}

// New allocates (github.com/anonx/sunplate/example/controllers).App controller,
// then returns it.
func (t App) New() *contr.App {
	c := &contr.App{}
	return c
}

// Before calls (github.com/anonx/sunplate/example/controllers).App.Before with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t App) Before(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic Before action of (github.com/anonx/sunplate/example/controllers).App.
	if r := c.Before( // Parameters should be binded.
		strconv.String(r.Form, "name"),
		strconv.Ints(r.Form, "pages"),
	); r != nil {
		return r
	}
	// Continue execution chain.
	return nil
}

// After calls (github.com/anonx/sunplate/example/controllers).App.After with arguments
// that are extracted from r.Form and converted to appropriate types.
func (t App) After(c *contr.App, w http.ResponseWriter, r *http.Request) a.Result {
	// Call magic After action of (github.com/anonx/sunplate/example/controllers).App.
	if r := c.After( // Parameters should be binded.
	); r != nil {
		return r
	}
	// Continue execution chain.
	return nil
}

func init() {
	_ = strconv.MeaningOfLife
}
