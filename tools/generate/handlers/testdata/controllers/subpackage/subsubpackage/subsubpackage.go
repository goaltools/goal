package subsubpackage

import (
	"net/http"
)

// SubSubPackage is a controller.
type SubSubPackage struct {
}

// Before does nothing.
func (c *SubSubPackage) Before() http.Handler {
	return nil
}

// Init should be ignored as the number of arguments is incorrect.
func Init() {
}
