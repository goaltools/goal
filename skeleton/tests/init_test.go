package tests

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/anonx/sunplate/skeleton/assets/views"
	"github.com/anonx/sunplate/skeleton/routes"

	"github.com/anonx/sunplate/controllers/rendering"
	"github.com/anonx/sunplate/log"
)

var (
	s *httptest.Server
	h *http.Handler
)

func init() {
	// Server is expected to be started from the root directory
	// of the project.
	os.Chdir("../")

	// Initialize a list of templates to use.
	rendering.SetTemplatePaths(views.Context)

	// Build a handler and prepare a test server.
	h, err := routes.List.Build()
	log.AssertNil(err)
	s = httptest.NewUnstartedServer(h)
}
