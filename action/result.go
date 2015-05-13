// Package action contains structs and interfaces that are
// imported from controller packages and used by actions.
package action

import (
	"net/http"
)

// Result is an interface that should be implemented by structs
// to be returned from actions.
type Result interface {
	// Apply writes the structure content to the response writer.
	Apply(w http.ResponseWriter, r *http.Request)

	// Finish is a method that returns true if it is expected
	// server should return result to the client.
	// No other results will be applied after a result that is finished.
	Finish() bool
}
