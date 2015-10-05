package subpackage

import (
	"net/http"
)

// Controller is some controller.
type Controller struct {
}

// Before is a magic function that is executed before any request.
func (c *Controller) Before() http.Handler {
	return nil
}

// Index is a sample action.
func (c Controller) Index(page int) http.Handler {
	return nil
}

// After is a magic function that is executed after any request.
func (c *Controller) After() http.Handler {
	return nil
}

// Finally is a magic function that is executed after any request
// no matter what.
func (c *Controller) Finally(http.ResponseWriter, *http.Request, []interface{}) bool {
	return false
}
