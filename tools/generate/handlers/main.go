// Package handlers scans your controllers and generates
// handler function package from them.
package handlers

import (
	"github.com/colegion/goal/internal/command"
)

// Handler is an instance of "new" subcommand (tool).
var Handler = command.Handler{
	Run: main,

	Name:  "generate handlers",
	Usage: "[flags]",
	Info:  "generate handler functions from controllers",
	Desc: `Tool "generate handlers" scans your controllers and generates
a standard handler function for every of your action.
So, you can use the generated package with any router you want.
`,
}

var input, output, pkg *string

func main(hs []command.Handler, i int, args command.Data) {
	start()
}

func init() {
	input = Handler.Flags.String("input", "./controllers", "a path to directory with controllers to scan")
	output = Handler.Flags.String("output", "./assets/handlers", "a directory where generated package must be saved")
	pkg = Handler.Flags.String("package", "handlers", "name of the package to generate")
}
