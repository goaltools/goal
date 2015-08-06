// Package create is used for copying skeleton app
// to a requested destination.
// Herewith, import path of the skeleton app is expected
// to be rewritten to a new one.
package create

import (
	"os"
	"path/filepath"

	"github.com/anonx/sunplate/internal/command"
	p "github.com/anonx/sunplate/internal/path"
	"github.com/anonx/sunplate/internal/programs/generation/output"
	"github.com/anonx/sunplate/log"
)

// Handler is an instance of new subcommand.
var Handler = command.Handler{
	Name:  "new",
	Info:  "create a skeleton application",
	Usage: "new {path}",
	Desc: `New creates files and directories to get a new Sunplate toolkit
based application running quickly.
All files and directories will be put into the given import path.

The path must be a directory that does not exist yet, e.g.:
	./sample

or alternatively:
	github.com/anonx/sample

Moreover, it is required to be located inside $GOPATH.

Examples:
	sunplate new github.com/anonx/sample
	sunplate new ./sample
	sunplate new ../anonx/sample
`,

	Main: start,
}

// sourceFiles contains extensions of files that should be process
// rather than just copied, a replacement of import path is expected
// in them. As an example there should be github.com/user/project
// instead of github.com/anonx/sunplate/internal/skeleton.
var sourceFiles = map[string]bool{
	".go":  true,
	".yml": true,
}

// start is an entry point of the command.
var start = func(action string, params command.Data) {
	oldImport := p.SunplateImport("internal", "skeleton")
	newImport := p.AbsoluteImport(params.Default(action, "./"))

	inputDir := p.SunplateDir("internal", "skeleton")
	outputDir := p.PackageDir(newImport)

	// Make sure the output directory does not exist yet.
	if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
		log.Error.Panicf(`Abort: Import path "%s" already exists.`, newImport)
	}

	rs, fn := walkFunc(inputDir)
	filepath.Walk(inputDir, fn)

	for _, v := range rs.dirs {
		t := output.Type{}
		t.CreateDir(filepath.Join(outputDir, v))
	}

	for k, v := range rs.files {
		copyFile(k, filepath.Join(outputDir, v))
	}

	for k, v := range rs.srcs {
		copyModifiedFile(k, filepath.Join(outputDir, v), map[string]string{
			oldImport: newImport,
		})
	}

	log.Info.Printf(info, newImport, newImport)
}

// result represents objects found when scanning a skeleton directory.
// There are a few possible kind of them: directories, static files,
// go source files (that require additional processing of their content).
type result struct {
	dirs, files, srcs map[string]string // Keys are full paths, values are relative ones.
}

// walkFunc returns a result instance and a function that may be used for validation
// of found elements. Successfully validated ones are stored to the returned result.
func walkFunc(dir string) (result, func(string, os.FileInfo, error) error) {
	dir = p.Prefixless(p.Prefixless(dir, "./"), ".") // No "./" is allowed at the beginning.
	rs := result{
		dirs:  map[string]string{},
		files: map[string]string{},
		srcs:  map[string]string{},
	}

	return rs, func(path string, info os.FileInfo, err error) error {
		// Make sure there are no any errors.
		if err != nil {
			log.Warn.Printf(`An error occured while scanning a skeleton: "%s".`, err)
			return err
		}

		// Get filepath without the dir path at the beginning.
		// So, when we are scanning "controllers/app/init.go" our generated
		// result will be "app/init.go" instead.
		rel, _ := filepath.Rel(dir, path)

		// Check whether current element is a directory.
		if info.IsDir() {
			rs.dirs[path] = rel
			return err
		}

		// Find out whether it is a static file or a go / some other source.
		ext := filepath.Ext(path)
		if sourceFiles[ext] {
			rs.srcs[path] = rel
			return err
		}

		// If it is a static file, add it to the list.
		rs.files[path] = rel
		return err
	}
}

var info = `Your application "%s" is ready:
You can run it with:
	sunplate run %s
`
