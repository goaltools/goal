package main

import (
	"fmt"

	"github.com/colegion/goal/internal/command"
	"github.com/colegion/goal/log"
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
~ https://github.com/colegion/goal
~
`

var helpMsg = `goal is a toolkit for rapid web development in Go language.

Usage:
	goal {command} [arguments]

The commands are:
%s
Use "goal help {command}" for more information.
`

var infoMsg = `Usage:
	goal %s

%s
`
