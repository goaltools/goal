// Package create is used for copying skeleton app
// to a requested destination.
// Herewith, import path of the skeleton app is expected
// to be rewritten to a new one.
package create

import (
	"os"
	"path/filepath"

	"github.com/colegion/goal/internal/command"
	"github.com/colegion/goal/internal/path"
	"github.com/colegion/goal/log"
)

// Handler is an instance of "new" subcommand (tool).
var Handler = command.Handler{
	Run: main,

	Name:  "new",
	Usage: "{path}",
	Info:  "create a skeleton application",
	Desc: `New creates files and directories to get a new app running quickly.
The created files and directories will be saved to the specified path.

The path must be a directory that does not exist yet, e.g:
	./sample

or alternatively:
	github.com/MyUsername/ProjectName

Moreover, it is required to be located inside "$GOPATH/src".

Examples:
	cli new github.com/colegion/sample
	cli new ./sample
	cli new ../colegion/sample
`,
}

// Main is an entry point of the subcommand (tool).
func main(hs []command.Handler, i int, args command.Data) {
	// The first argument in the list is a path.
	// If it's missing use an empty string instead.
	p := args.GetDefault(0, "")

	// Prepare source and destination directory paths.
	src, err := path.New("github.com/colegion/goal/internal/skeleton").Package()
	log.AssertNil(err)
	dest, err := path.New(p).Package()
	log.AssertNil(err)

	// Prepare an import path of the app to be created.
	destImp, err := dest.Import()
	log.AssertNil(err)

	// Make sure the requested import path (dest) does not exist yet.
	if _, err := os.Stat(dest.String()); !os.IsNotExist(err) {
		log.Error.Panicf(`Cannot use "%s", such import path already exists.`, destImp.String())
	}

	// Scan the skeleton directory and get a list of directories / files
	// to be copied / processed.
	res, err := walk(src.String())
	log.AssertNil(err)

	// Create the directories in destination path.
	for i := 0; i < len(res.dirs); i++ {
		err = os.MkdirAll(filepath.Join(dest.String(), res.dirs[i]), 0755)
		log.AssertNil(err)
	}

	// Copy static files to the destination directories.
	for i := 0; i < len(res.files); i++ {
		copyFile(res.files[i].absolute, filepath.Join(dest.String(), res.files[i].relative))
	}

	// Process source files and copy to the destination directories.
	for i := 0; i < len(res.srcs); i++ {
		copyModifiedFile(
			res.srcs[i].absolute, filepath.Join(dest.String(), res.srcs[i].relative), [][][]byte{
				{
					[]byte("github.com/colegion/goal/internal/skeleton"),
					[]byte(destImp.String()),
				},
			},
		)
	}

	log.Info.Printf(info, destImp)
}

// Arguments to format are:
//	[1]: destination app's import path.
var info = `Your application "%[1]s" is ready:
You can run it with:
	goal run %[1]s
`
