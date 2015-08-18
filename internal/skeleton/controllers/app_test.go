package controllers_test

import (
	"testing"

	"github.com/anonx/sunplate/internal/skeleton/assets/handlers"

	"github.com/anonx/sunplate/assert"
)

func TestAppIndex(t *testing.T) {
	a := assert.New(t)

	a.Get("/")
	handlers.App.Index(a.Args())
	a.Status(200)
	a.ContentType("text/html; charset=utf-8")
	a.Body.Contains("Hello, world!")
}
