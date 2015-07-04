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

Status of the project: **WIP** (it does't work yet).

[![GoDoc](https://godoc.org/github.com/anonx/sunplate?status.svg)](https://godoc.org/github.com/anonx/sunplate)
[![Build Status](https://travis-ci.org/anonx/sunplate.svg?branch=master)](https://travis-ci.org/anonx/sunplate)
[![Coverage Status](https://coveralls.io/repos/anonx/sunplate/badge.svg?branch=master)](https://coveralls.io/r/anonx/sunplate?branch=master)
[![Go Report Card](http://goreportcard.com/badge/anonx/sunplate?t=3)](http:/goreportcard.com/report/anonx/sunplate)

## Commands
The list of currently planned commands include:

#### General commands
##### - `sunplate run path/to/app` - Run a tool that watches the app files and rebuilds if necessary. It can be used as a task runner, too (Planned).
##### - `sunplate help` - Show information about the tool and supported commands.

#### Code generation
These commands are not implied to be run manually but rather using `go generate` mechanism. To do so, include in your `.go` file the following line:
```go
//go:generate sunplate generate {command} [arguments]
```

##### - `sunplate generate handlers`- Scan controllers, generate handlers (WIP).
Default parameters are:
* `--input ./controllers`
* `--output ./assets/handlers`
* `--package handlers`

Read more about controllers and actions [here](https://github.com/anonx/concept/blob/master/basics.md#basics).
The idea is to have automatically generated handlers (just usual golang handlers) from Revel framework like controllers.
Currently, `Before` and `After` magic actions are implemented.
**TODO**: support of actions, `Finally` magic method.

##### - `sunplate generate listing`- Scan a directory, generate a mapping of file names to their paths.
Default parameters are:
* `--input ./views`
* `--output ./assets/views`
* `--package views`

This command generates a listing of all found files in a requested directory.
To to advantage of the generated package import it and use the `Context map[string]string` variable.
Its format is:
```go
map[string]string{
	"main.html":          "./views/main.html",
	"accounts/info.html": "./views/accounts/info.html",
}
```

##### - `sunplate generate routes`- Scan handlers, generate routes (Planned, router is ready).

##### - `sunplate generate autoforms`- Scan models, generate a package for easy validation, binding, and rendering of forms.

## License
Distributed under the BSD 2-clause "Simplified" License unless otherwise noted.
