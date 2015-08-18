package controllers_test

import (
	"net/url"
	"testing"

	"github.com/anonx/sunplate/internal/skeleton/assets/handlers"
	"github.com/anonx/sunplate/internal/skeleton/routes"

	"github.com/anonx/sunplate/assert"
)

func TestAppIndex(t *testing.T) {
	a := assert.New(t)

	handlers.App.Index(a.Get("/").Args())
	a.Status(200)
	a.ContentType("text/html; charset=utf-8")
	a.Body.Contains("Hello, world!")
}

func TestAppPostGreet_Integration(t *testing.T) {
	a := assert.New(t)
	a.TryStartServer(routes.List.Build())
	defer a.StopServer()

	a.PostForm("/greet/James", url.Values{
		"message": {"Good day"},
	}).Do()
	a.StatusOK()
	a.Body.Contains("Good day")
	a.Body.Contains("James")
}
