# sunplate
Toolkit for high productivity web development in Go language.

Sunplate, being mostly inspired by [Revel Framework](https://github.com/revel/revel)
and its discussions, is built around the concept of
controllers and actions ([more about the concept](https://github.com/anonx/concept#concept)).
However, as opposed to Revel and other high level frameworks Sunplate does not use runtime
reflection and does not require your app to import monolithic dependencies.

Instead Sunplate uses code generation and `go generate` mechanism.
That allows us to achieve type safety, minimalism of dependencies,
compatability with the standard libraries, and productivity for the end-developers.
At the same time Sunplate is very customizable (you can bring your own router, template system,
and any other component). But without prejudice to the easiness and seamless of experience
thanks to good defaults.

Status of the project: **PoC** (not ready for use in the wild).

[![GoDoc](https://godoc.org/github.com/anonx/sunplate?status.svg)](https://godoc.org/github.com/anonx/sunplate)
[![Build Status](https://travis-ci.org/anonx/sunplate.svg?branch=master)](https://travis-ci.org/anonx/sunplate)
[![Coverage Status](https://coveralls.io/repos/anonx/sunplate/badge.svg?branch=master)](https://coveralls.io/r/anonx/sunplate?branch=master)
[![Go Report Card](http://goreportcard.com/badge/anonx/sunplate?t=3)](http:/goreportcard.com/report/anonx/sunplate)

## Installation
```sh
go install github.com/anonx/sunplate
```

## Getting started
**Step 1**: Create a skeleton application
```sh
sunplate new {path}    # e.g. "github.com/anonx/sample" or "./sample"
```

**Step 2**: Start a task runner
```sh
sunplate run {path}
```

**Step 3**: Start making changes you need to the generated skeleton app.

**Step 4**: Use `sunplate help` to get more information about supported commands
and `sunplate help {command}` to find out more about a specific command.

## License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
