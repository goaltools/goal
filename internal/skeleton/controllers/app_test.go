package controllers_test

import (
	"net/url"
	"testing"

	"github.com/anonx/sunplate/internal/skeleton/assets/handlers"
	"github.com/anonx/sunplate/internal/skeleton/routes"

	"github.com/anonx/sunplate/assert"
)

func TestAppIndex(t *testing.T) {
	a := assert.New()

	handlers.App.Index(a.Get("/").Args())
	a.StatusOK().ContentType("text/html; charset=utf-8")
	a.Body.Contains("Hello, world!")
}

func TestAppPostGreet(t *testing.T) {
	a := assert.New().TryStartServer(routes.List.Build())
	defer a.StopServer()

	a.PostForm("/greet/James", url.Values{
		"message": {"Good day"},
	}).Do().StatusOK()

	a.Body.Contains("James")
	a.Body.Contains("Good day")
}
