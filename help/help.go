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
		log.Info.Println(info)
	}
}

var info = `ok command argument [arguments]

The commands are:
	new         create a skeleton application
	run         run a watcher / task runner
	generate    analize files, build handlers, routes, etc.

Use "ok help command" for more information.
`
