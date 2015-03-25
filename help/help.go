// Package help is a command that shows info
// about the commands 'ok' toolkit supports.
package help

import (
	"github.com/anonx/ok/log"
)

// Start is an entry point of help command.
func Start(action string, params map[string]string) {
	switch params[action] {
	default:
		log.Info.Printf(info, action)
	}
}

var info = `Ok is a toolkit for rapid web development in Go language.

Usage:
	ok command [arguments]

The commands are:
	generate    analize files, build handlers, routes, etc.
	new         create a skeleton application
	run         run a watcher / task runner

Use "ok %s [command]" for more information.
`
