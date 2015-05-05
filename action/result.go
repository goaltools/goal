// Package action contains structs and interfaces that are
// imported from controller packages and used by actions.
package action

import (
	"net/http"
)

// Result is an interface that should be implemented by structs
// to be returned from actions.
type Result interface {
	// BasePath initializes structure with the path of an app.
	// It will be called automatically right before Apply.
	BasePath(s string)

	// Apply writes the structure content to the response writer.
	Apply(w http.ResponseWriter, r *http.Request)
}
