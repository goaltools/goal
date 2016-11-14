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

        go get -u github.com/goaltools/goal

2. Create a [new skeleton application](https://goaltools.github.io/goal/manual/new/index.html):

        goal new github.com/$username/$project

3. Start a [watcher / task runner](https://goaltools.github.io/goal/manual/run/index.html):

        goal run github.com/$username/$project

### Documentation

* **[goaltools.github.io/goal](https://goaltools.github.io/goal)**
  * [Introduction](https://goaltools.github.io/goal/manual/index.html)
  * [`goal new`](https://goaltools.github.io/goal/manual/new/index.html) - generate skeleton application
  * [`goal run`](https://goaltools.github.io/goal/manual/run/index.html) - start watcher / task runner
  * [`goal generate handlers`](https://goaltools.github.io/goal/manual/handlers/index.html) - generate HTTP handler functions
    * [Actions](https://goaltools.github.io/goal/manual/handlers/actions.html)
    * [Controllers](https://goaltools.github.io/goal/manual/handlers/controllers.html)
    * [Routes](https://goaltools.github.io/goal/manual/handlers/routes.html) - a slice of routes and handler functions associated with them
  * `goal generate listing` - ~~generate a list of file paths~~ (deprecated)

All `goal generate *` tools may be used with [`go generate`](https://blog.golang.org/generate).

### Status

**Proof of Concept**: not ready for use in the wild. **Working on**: splitting the project into
independent repos.

[![GoDoc](https://godoc.org/github.com/goaltools/goal?status.svg)](https://godoc.org/github.com/goaltools/goal)
[![Build Status](https://travis-ci.org/goaltools/goal.svg?branch=master)](https://travis-ci.org/goaltools/goal)
[![Coverage Status](https://coveralls.io/repos/goaltools/goal/badge.svg?branch=master)](https://coveralls.io/r/goaltools/goal?branch=master)
[![Go Report Card](http://goreportcard.com/badge/goaltools/goal?t=3)](http:/goreportcard.com/report/goaltools/goal)

### License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
