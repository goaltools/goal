# sunplate
Toolkit for high productivity web development in Go language.

Sunplate, being mostly inspired by [Revel Framework](https://github.com/revel/revel)
and its discussions, is built around the concept of
controllers and actions.
However, as opposed to Revel and other high level frameworks Sunplate does not use runtime
reflection and does not require your app to import monolithic dependencies.

Instead Sunplate is implemented in a form of independent utilities that
may be used with [`go generate`](https://blog.golang.org/generate).
That allows us to achieve type safety, minimalism of dependencies,
compatability with the standard library, and productivity for the end-developers.
At the same time Sunplate is very customizable (you can bring your own router, template system,
and any other component). But that's without prejudice to the easiness and seamless of experience
thanks to good defaults.

Status of the project: **PoC** (not ready for use in the wild).

[![GoDoc](https://godoc.org/github.com/anonx/sunplate?status.svg)](https://godoc.org/github.com/anonx/sunplate)
[![Build Status](https://travis-ci.org/anonx/sunplate.svg?branch=master)](https://travis-ci.org/anonx/sunplate)
[![Coverage Status](https://coveralls.io/repos/anonx/sunplate/badge.svg?branch=master)](https://coveralls.io/r/anonx/sunplate?branch=master)
[![Go Report Card](http://goreportcard.com/badge/anonx/sunplate?t=3)](http:/goreportcard.com/report/anonx/sunplate)

### Learn More
* https://sunplate.club/

### License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
