# Goal
**Goal is a set of tools for high productivity web development in Go language.**

Goal, being mostly inspired by [Revel Framework](https://github.com/revel/revel)
and its discussions, is built around the concept of controllers and actions.
However, as opposed to Revel and other high level frameworks Goal does not use runtime
reflection and does not require your app to import monolithic dependencies.

Instead Goal is implemented in a form of independent tools that
may be used with [`go generate`](https://blog.golang.org/generate).
That allows us to achieve type safety, minimalism of dependencies,
compatability with the standard library, and productivity for the end-developers.
At the same time Goal is very customizable (you can bring your own router, template system,
and any other component). But that's without prejudice to the easiness and seamless of experience
thanks to good defaults.

### Getting Started

1. Install Goal:

        go get -u github.com/colegion/goal

2. Create a [new skeleton application](https://colegion.github.io/goal/manual/new/index.html):

        goal new github.com/$username/$project

3. Start a [watcher / task runner](https://colegion.github.io/goal/manual/run/index.html):

        goal run github.com/$username/$project

### Documentation

* **[colegion.github.io/goal](https://colegion.github.io/goal)**
  * [Introduction](https://colegion.github.io/goal/manual/index.html)
  * [`goal new`](https://colegion.github.io/goal/manual/new/index.html) - generate skeleton application
  * [`goal run`](https://colegion.github.io/goal/manual/run/index.html) - start watcher / task runner
  * [`goal generate handlers`](https://colegion.github.io/goal/manual/handlers/index.html) - generate HTTP handler functions
    * [Actions](https://colegion.github.io/goal/manual/handlers/actions.html)
    * [Controllers](https://colegion.github.io/goal/manual/handlers/controllers.html)
  * `goal generate listing` - generate a list of file paths

All `goal generate *` tools may be used with [`go generate`](https://blog.golang.org/generate).

### Overview

#### Controllers and Actions

```go
// Profiles is a sample controller.
type Profiles struct {
	*Controllers
}

// List is an action that renders a page with some profiles.
func (c *Profiles) List(page int) http.Handler {
	c.Context["profiles"] = models.PaginatedProfiles(page)
	return c.Render()
}
```
Value of `page` will be automatically binded from Form or URL.

#### Routes

```go
//@get /greet/:name
func (c *App) Greet(name string) http.Handler {
	c.Context["name"] = name
	return c.Render()
}
```

#### Magic Actions

Magic actions are the actions that are run automatically before or after requests.
```go
type ParentController struct {
	*Controllers
}

func (c *ParentController) Before() http.Handler {
	println("parent: Before")
	return nil
}

type AppController struct {
	*ParentController
}

func (c *AppController) Before() http.Handler {
	println("app: Before")
	return nil
}

func (c *AppController) Index() http.Handler {
	println("app: Index Action")
	return c.Render()
}

func (c *AppController) After() http.Handler {
	println("app: After")
	return nil
}
```
On every request the code above will print:
```
parent: Before
app: Before
app: Index Action
app: After
```

### Status

**Proof of Concept**: not ready for use in the wild.

[![GoDoc](https://godoc.org/github.com/colegion/goal?status.svg)](https://godoc.org/github.com/colegion/goal)
[![Build Status](https://travis-ci.org/colegion/goal.svg?branch=master)](https://travis-ci.org/colegion/goal)
[![Coverage Status](https://coveralls.io/repos/colegion/goal/badge.svg?branch=master)](https://coveralls.io/r/colegion/goal?branch=master)
[![Go Report Card](http://goreportcard.com/badge/colegion/goal?t=3)](http:/goreportcard.com/report/colegion/goal)

### License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
