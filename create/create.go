// Package create is used for copying skeleton app
// to a requested destination.
// Herewith, import path of the skeleton app is expected
// to be rewritten to a new one.
package create

import (
	"os"
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/log"
	p "github.com/anonx/sunplate/path"
)

// result represents objects found when scanning a skeleton directory.
// There are a few possible kind of them: directories, static files,
// go source files (that require additional processing of their content).
type result struct {
	dirs, files, srcs map[string]string // Keys are full paths, values are relative ones.
}

// Start is an entry point of the command.
func Start(action string, params command.Data) {
	oldImport := p.SunplateImport("skeleton")
	newImport := p.AbsoluteImport(params.Default(action, "./"))

	inputDir := p.SunplateDir("skeleton")
	outputDir := p.PackageDir(newImport)

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

// walkFunc returns a result instance and a function that may be used for validation
// of found elements. Successfully validated ones are stored to the returned result.
func walkFunc(dir string) (result, func(string, os.FileInfo, error) error) {
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

		relPath := p.Prefixless(path, dir)

		// Check whether current element is a directory.
		if info.IsDir() {
			rs.dirs[path] = relPath
			return err
		}

		// Find out whether it is a static file or a go source.
		if filepath.Ext(path) == ".go" {
			rs.srcs[path] = relPath
			return err
		}

		// If it is a static file, add it to the list.
		rs.files[path] = relPath
		return err
	}
}

var info = `Your application is ready:
	%s
You can run it with:
	sunplate run %s
`
