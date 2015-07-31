// Package echo is used for printing text.
// It is not possible to use "echo" command in "sunplate.yml"
// as that would not work on Windows.
// Thus, "sunplate echo" is a way around.
package echo

import (
	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/log"
)

// Handler is an instance of echo subcommand.
var Handler = command.Handler{
	Name:  "echo",
	Info:  "print a line of text",
	Usage: "echo [text]",
	Desc: `Sometimes it is required to add "echo 'smth'" to "sunplate.yml" config.
However, the command works only on *Nix platforms as Windows requires
"cmd /c echo" rather than just "echo". "sunplate echo" is a workaround.
This command uses a regular print to show the requested text.
`,

	Main: start,
}

// start is an entry point of the command.
var start = func(action string, params command.Data) {
	log.Info.Print(params.Default(action, ""))
}
