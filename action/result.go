// Package action contains structs and interfaces that are
// imported from controller packages and used by actions.
package action

import (
	"net/http"
)

// Result is an interface that should be implemented by structs
// to be returned from actions.
type Result interface {
	// Apply will be called from handler function if Result is not nil.
	// Use it to write content to the response writer.
	Apply(w http.ResponseWriter, r *http.Request)
}
