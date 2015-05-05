// Package template provides functions for work
// with standard Go template engine.
// It should be embeded into Controller struct.
package template

// TemplatePaths contains paths to directories that should be scanned
// for golang templates.
// By-default it includes local "./views" directory.
var TemplatePaths = []string{
	"./views",
}

// Middleware is a main type that should be embeded into controller structs.
type Middleware struct {
	// Context is used for passing variables to templates.
	Context map[string]interface{}
}
