package controllers

import (
	"net/http"
)

// Errors is a controller with actions displaying error pages.
type Errors struct {
	*Controllers
}

// NotFound prints an error 404 "Page Not Found" message.
func (c *Errors) NotFound() http.Handler {
	return c.RenderNotFound()
}

// InternalError displays an error 500 "Internal Server Error" message.
func (c *Errors) InternalError() http.Handler {
	return c.RenderError(nil)
}
