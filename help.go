package main

import (
	"fmt"

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
	// Make sure we can show help message about the requested
	// subcommand or it is not supported.
	if h, ok := Handlers[params[action]]; ok {
		log.Info.Printf(infoMsg, h.Usage, h.Desc)
		return
	}
	log.Info.Printf(helpMsg, showCommands()) // Show general message.
}

// showCommands returns a description of supported commands
// to be included in help message.
func showCommands() (s string) {
	for n := range Handlers {
		inf := Handlers[n].Info
		if inf != "" {
			s += fmt.Sprintf("\t%-12s%s\n", n, inf)
		}
	}
	return
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
%s
Use "sunplate help {command}" for more information.
`

var infoMsg = `Usage:
	sunplate %s

%s
`
