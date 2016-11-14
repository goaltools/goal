// Package create is used for copying skeleton app
// to a requested destination.
// Herewith, import path of the skeleton app is expected
// to be rewritten to a new one.
package create

import (
	"os"
	"path/filepath"

	"github.com/goaltools/goal/internal/log"
	"github.com/goaltools/goal/utils/tool"

	"github.com/goaltools/importpath"
)

// Handler is an instance of "new" subcommand (tool).
var Handler = tool.Handler{
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
	cli new github.com/goaltools/sample
	cli new ./sample
	cli new ../goaltools/sample
`,
}

// Main is an entry point of the subcommand (tool).
func main(hs []tool.Handler, i int, args tool.Data) {
	// The first argument in the list is a path.
	// If it's missing use an empty string instead.
	p := args.GetDefault(0, "")

	// Prepare source and destination directory paths.
	src, err := importpath.ToPath("github.com/goaltools/goal/internal/skeleton")
	if err != nil {
		log.Error.Panic(err)
	}
	destImp, err := importpath.ToImport(p)
	if err != nil {
		log.Error.Panic(err)
	}
	dest, err := importpath.ToPath(destImp)
	if err != nil {
		log.Error.Panic(err)
	}

	// Make sure the requested import path (dest) does not exist yet.
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		log.Error.Panicf(`Cannot use "%s", such import path already exists.`, destImp)
	}

	// Scan the skeleton directory and get a list of directories / files
	// to be copied / processed.
	res, err := walk(src)
	if err != nil {
		log.Error.Panic(err)
	}

	// Create the directories in destination path.
	for i := 0; i < len(res.dirs); i++ {
		err = os.MkdirAll(filepath.Join(dest, res.dirs[i]), 0755)
		if err != nil {
			log.Error.Panic(err)
		}
	}

	// Copy static files to the destination directories.
	for i := 0; i < len(res.files); i++ {
		copyFile(res.files[i].absolute, filepath.Join(dest, res.files[i].relative))
	}

	// Process source files and copy to the destination directories.
	for i := 0; i < len(res.srcs); i++ {
		copyModifiedFile(
			res.srcs[i].absolute, filepath.Join(dest, res.srcs[i].relative), [][][]byte{
				{
					[]byte("github.com/goaltools/goal/internal/skeleton"),
					[]byte(destImp),
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
