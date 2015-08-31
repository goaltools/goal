package templates

import (
	"net/http"

	"github.com/anonx/sunplate/log"
)

// Handler is a template handler that implements http.Handler interface.
type Handler struct {
	context  map[string]interface{} // Variables to be passed to the template.
	template string                 // Path to the template to be rendered.
	status   int                    // Expected status code of the response.
}

// Apply writes to response the result received from action.
func (t *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set status of the response.
	if t.status == 0 {
		t.status = http.StatusOK
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// If required template exists, execute it.
	if tpl, ok := templates[t.template]; ok {
		w.WriteHeader(t.status)
		err := tpl.ExecuteTemplate(w, templateName, t.context)
		if err != nil {
			go log.Warn.Println(err)
		}
		return
	}

	// Otherwise, show internal server error.
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
	go log.Warn.Printf(`Template "%s" does not exist.`, t.template)
}
