// Package help is a command that shows info
// about the commands 'sunplate' toolkit supports.
package help

import (
	"github.com/anonx/sunplate/log"
)

// Start is an entry point of help command.
func Start(action string, params map[string]string) {
	switch params[action] {
	default:
		log.Info.Printf(info, action)
	}
}

var info = `Sunplate is a toolkit for rapid web development in Go language.

Usage:
	sunplate command [arguments]

The commands are:
	generate    analize files, build handlers, routes, etc.
	new         create a skeleton application
	run         run a watcher / task runner

Use "sunplate %s [command]" for more information.
`
