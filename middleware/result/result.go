// Package result is a middleware for ok toolkit applications.
// It should be embeded into Controller struct.
// Result provides functions for result rendering such as
// Render, RenderJSON, etc.
package result

// Middleware is a main type that should be embeded into controller structs.
type Middleware struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}
}
