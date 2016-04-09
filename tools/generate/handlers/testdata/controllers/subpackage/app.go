package subpackage

import (
	"net/http"
	"net/url"

	"github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage"
	"github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/x"
)

// Controller is some controller.
type Controller struct {
	*subsubpackage.SubSubPackage
	*x.X
}

// Before is a magic function that is executed before any request.
func (c *Controller) Before() http.Handler {
	return nil
}

// Index is a sample action.
//@post index someindexlabel
func (c Controller) Index(page int) http.Handler {
	return nil
}

// After is a magic function that is executed after any request.
func (c *Controller) After() http.Handler {
	return nil
}

// Finally is a magic function that is executed after any request
// no matter what.
func (c *Controller) Finally(w http.ResponseWriter, r *http.Request, a []string) bool {
	return false
}

// Init ...
func Init(ctx url.Values) {
}
