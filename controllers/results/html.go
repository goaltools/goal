package results

import (
	"net/http"

	"github.com/anonx/sunplate/log"
)

// HTML is a result that is returned from actions by default.
// This struct implements action.Result interface.
type HTML struct {
	context  map[string]interface{} // Variables to be passed to the template.
	template string                 // Path to the template to be rendered.
}

// Apply writes to response the result received from action.
func (t *HTML) Apply(w http.ResponseWriter, r *http.Request) {
	// If required template exists, execute it.
	if tpl, ok := templates[t.template]; ok {
		err := tpl.ExecuteTemplate(w, TemplateName, t.context)
		if err != nil {
			go log.Warn.Println(err)
		}
		return
	}

	// Otherwise, show internal server error.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
	go log.Warn.Printf(`Template "%s" does not exist.`, t.template)
}

// Finish indicates that the next (magic) action should be started
// after this one.
func (t *HTML) Finish() bool {
	return false
}
