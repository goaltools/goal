package testsuite

import (
	"net/http"

	"github.com/anonx/sunplate/log"
)

// AssertOK makes sure that the last time a request was made
// 200 OK was returned.
func (t *Type) AssertOK() {
	t.AssertStatus(http.StatusOK)
}

// AssertNotFound is used to ensure and error 404 not found
// was returned at the time of making a request.
func (t *Type) AssertNotFound() {
	t.AssertStatus(http.StatusNotFound)
}

// AssertStatus validates the response contains a specific
// status code.
func (t *Type) AssertStatus(status int) {
	if t.Response.StatusCode != status {
		log.Error.Panicf("Status: (expected) %d != %d (actual)", status, t.Response.StatusCode)
	}
}
