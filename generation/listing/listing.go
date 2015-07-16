// Package listing is a generate subcommand that scans
// requested directory and generates a .go file with
// a list of all found files.
package listing

import (
	"os"
	"path/filepath"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/log"
	p "github.com/anonx/sunplate/path"
)

// Start is an entry point of listing subcommand.
// It expects two parameters.
// basePath is where to find files necessary for generation of listing.
// params is a map with the following keys:
// --path defines what directory to analyze ("./views" by-default).
// --output is a path to directory where to create a new package ("./assets" by-default).
// --package is what package should be created as a result ("views" by-default).
func Start(params command.Data) {
	inputDir := params.Default("--input", "./views")
	outputDir := params.Default("--output", "./assets/views")
	outPkg := params.Default("--package", "views")

	// Start search of files.
	fs, fn := walkFunc(inputDir)
	filepath.Walk(inputDir, fn)

	// Generate and save a new package.
	t := output.NewType(
		outPkg, filepath.Join(
			p.SunplateDir("generation", "listing"), "./listing.go.template",
		),
	)
	t.CreateDir(outputDir)
	t.Extension = ".go" // Save generated file as a .go source.
	t.Context = map[string]interface{}{
		"files": fs,
		"input": params.Default("--path", "./views"),
	}
	t.Generate()
}

// walkFunc returns a map of files and a function that may be used for validation
// of found files. Successfully validated ones are stored to the files variable.
func walkFunc(dir string) (map[string]string, func(string, os.FileInfo, error) error) {
	files := map[string]string{}

	return files, func(path string, info os.FileInfo, err error) error {
		// Make sure there are no any errors.
		if err != nil {
			log.Warn.Printf(`An error occured while creating a listing: "%s".`, err)
			return err
		}

		// Assure it is not a directory.
		if info.IsDir() {
			return nil
		}

		// Add files to the list.
		log.Trace.Printf(`Path "%s" discovered.`, path)
		files[p.Prefixless(path, dir)] = path
		return nil
	}
}
