// Package create is used for copying skeleton app
// to a requested destination.
// Herewith, import path of the skeleton app is expected
// to be rewritten to a new one.
package create

import (
	"errors"
	"fmt"
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
	Usage: "{import_path}",
	Info:  "create a skeleton application",
	Desc: `New creates files and directories to get a new app running quickly
in the requested import path directory.

The import path must be a directory that does not exist yet, e.g.:
	github.com/MyUsername/ProjectName
`,
}

// Main is an entry point of the subcommand (tool).
func main(hs []command.Handler, i int, args []string) error {
	// Make sure "import path" argument is present.
	if len(args) == 0 {
		return errors.New("import path argument is required")
	}

	// Prepare source and destination directory paths.
	src, err := path.New("github.com/colegion/goal/internal/path").Package()
	if err != nil {
		return err
	}
	dest, err := path.New(args[0]).Absolute()
	if err != nil {
		return err
	}

	// Prepare an import path of the app to be created.
	destImp, err := dest.Import()
	if err != nil {
		return err
	}

	// Make sure the requested import path (dest) does not exist yet.
	if _, err := os.Stat(dest.String()); !os.IsNotExist(err) {
		return fmt.Errorf(`cannot use "%s", such import path already exists`, dest)
	}

	// Scan the skeleton directory and get a list of directories / files
	// to be copied / processed.
	res, err := walk(src.String())
	if err != nil {
		return err
	}

	// Create the directories in destination path.
	for i := 0; i < len(res.dirs); i++ {
		err = os.MkdirAll(filepath.Join(dest.String(), res.dirs[i]), 0755)
		if err != nil {
			return err
		}
	}

	// Copy static files to the destination directories.
	for i := 0; i < len(res.files); i++ {
		err = copyFile(res.files[i].absolute, filepath.Join(dest.String(), res.files[i].relative))
		if err != nil {
			return err
		}
	}

	// Process source files and copy to the destination directories.
	for i := 0; i < len(res.srcs); i++ {
		err = copyModifiedFile(
			res.srcs[i].absolute, filepath.Join(dest.String(), res.srcs[i].relative), [][][]byte{
				[][]byte{
					[]byte("github.com/colegion/goal/internal/skeleton"),
					[]byte(destImp.String()),
				},
			},
		)
		if err != nil {
			return err
		}
	}

	log.Info.Printf(info, destImp)
	return nil
}

// Arguments to format are:
//	[1]: destination app's import path.
var info = `Your application "%[1]s" is ready:
You can run it with:
	goal run %[1]s
`
