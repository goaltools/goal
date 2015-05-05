// Package scan is a ok toolkit command that scans
// requested directory and generates a .go file with
// a map of all found files.
package scan

import (
	"github.com/anonx/ok/log"
)

// Start is an entry point of scan command.
func Start(action string, params map[string]string) {
	if params[action] != "directory" {
		log.Warn.Printf(
			"I do not know how to scan '%s'.\nRun 'ok help %s' for more information.",
			params[action], action,
		)
	}
}
