package controllers

import (
	"github.com/anonx/ok/middleware/result"
)

// Controller is a struct that should be embedded into every controller
// to make such methods as Render, RenderJSON, etc. available.
type Controller struct {
	result.Middleware
}
