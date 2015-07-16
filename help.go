package main

import (
	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/log"
)

// helpHandler is an instance of a subcommand that is used
// for showing info about supported commands.
var helpHandler = command.Handler{
	Name: "help",
	Main: help,
}

// help is used for showing info about supported commands.
func help(action string, params command.Data) {
	switch params[action] {
	default:
		log.Info.Printf(helpMsg, action)
	}
}

var header = `~
~ https://github.com/anonx/sunplate
~
                        _       _
                       | |     | |
  ___ _   _ _ __  _ __ | | __ _| |_ ___
 / __| | | | '_ \| '_ \| |/ _' | __/ _ \
 \__ \ |_| | | | | |_) | | (_| | ||  __/
 |___/\__,_|_| |_| .__/|_|\__,_|\__\___|
                 | |
                 |_|
`

var helpMsg = `Sunplate is a toolkit for rapid web development in Go language.

Usage:
	sunplate {command} [arguments]

The commands are:
	new         create a skeleton application
	run         run a watcher / task runner

	generate    analize files, build handlers, routes, etc.

Use "sunplate %s [command]" for more information.
`
