# Goal
Goal is a toolkit for high productivity web development in Go language.

Goal, being mostly inspired by [Revel Framework](https://github.com/revel/revel)
and its discussions, is built around the concept of
controllers and actions.
However, as opposed to Revel and other high level frameworks goal does not use runtime
reflection and does not require your app to import monolithic dependencies.

Instead goal is implemented in a form of independent utilities that
may be used with [`go generate`](https://blog.golang.org/generate).
That allows us to achieve type safety, minimalism of dependencies,
compatability with the standard library, and productivity for the end-developers.
At the same time goal is very customizable (you can bring your own router, template system,
and any other component). But that's without prejudice to the easiness and seamless of experience
thanks to good defaults.

# Getting Started

1. Install Goal toolkit:

        go get -u github.com/colegion/goal

2. Create a new skeleton application:

        goal new github.com/$username/$project

3. Start a watcher / task runner:

        goal run github.com/$username/$project

* [Documentation](http://colegion.github.io/goal)

# Status

**PoC** (not ready for use in the wild).

[![GoDoc](https://godoc.org/github.com/colegion/goal?status.svg)](https://godoc.org/github.com/colegion/goal)
[![Build Status](https://travis-ci.org/colegion/goal.svg?branch=master)](https://travis-ci.org/colegion/goal)
[![Coverage Status](https://coveralls.io/repos/colegion/goal/badge.svg?branch=master)](https://coveralls.io/r/colegion/goal?branch=master)
[![Go Report Card](http://goreportcard.com/badge/colegion/goal?t=3)](http:/goreportcard.com/report/colegion/goal)

### License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
