package requests

import (
	"net/http"

	"github.com/colegion/goal/log"
)

// Requests is a controller that does two things:
// 1. Calls Request.ParseForm to parse GET / POST requests;
// 2. Makes Request available in your controller (use c.Request).
type Requests struct {
	Request *http.Request
}

// Before is a magic action of Requests controller.
func (c *Requests) Before() http.Handler {
	return nil
}

// Initially calls ParseForm on the request and saves it to c.Request.
// At the same time, if used with a standard goal routing package,
// parameters extracted from URN are saved to the Form field of the Request.
func (c *Requests) Initially(w http.ResponseWriter, r *http.Request, a []interface{}) bool {
	// Save the old value of Form, "github.com/colegion/goal/routing"
	// stores parameters extracted from URN there.
	t := r.Form

	// Set r.Form to nil, otherwise ParseForm will not work.
	r.Form = nil

	// Parse request and make sure there are no errors.
	err := r.ParseForm()
	if err != nil {
		go log.Warn.Printf("Error parsing request body: %v.", err)
	}

	// Add the old values from router to the new r.Form.
	for k, vs := range t {
		r.Form.Add(k, vs[0])
	}

	// Save the request, so it may be used from child controllers.
	c.Request = r
	return false
}
