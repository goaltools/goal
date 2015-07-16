package main

import (
	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/log"
)

// Show a help message about supported commands.
func showHelp(action string, params command.Data) {
	switch params[action] {
	default:
		log.Info.Printf(helpMsg, action)
	}
}

func catchPanic() {
	if err := recover(); err != nil {
		log.Error.Fatal(err)
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
