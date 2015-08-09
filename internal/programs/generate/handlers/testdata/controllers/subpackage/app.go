package subpackage

import (
	"net/http"

	"github.com/anonx/sunplate/action"
)

// Controller is some controller.
type Controller struct {
}

// Before is a magic function that is executed before any request.
func (c *Controller) Before() action.Result {
	return nil
}

// Index is a sample action.
func (c Controller) Index(page int) action.Result {
	return nil
}

// After is a magic function that is executed after any request.
func (c *Controller) After() action.Result {
	return nil
}

// Finally is a magic function that is executed after any request
// no matter what.
func (c *Controller) Finally(http.ResponseWriter, *http.Request) bool {
	return false
}
