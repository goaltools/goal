// Package listing is a generate subcommand that scans
// requested directory and generates a .go file with
// a list of all found files.
package listing

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/log"
)

var (
	rootPath string
	files    map[string]string
)

// Start is an entry point of listing subcommand.
// It expects two parameters.
// basePath is where to find files necessary for generation of listing.
// params is a map with the following keys:
// --path defines what directory to analyze ("./views" by-default).
// --output is a path to directory where to create a new package ("./assets" by-default).
// --package is what package should be created as a result ("views" by-default).
func Start(basePath string, params map[string]string) {
	// Initialize missed parameters.
	initDefaults(params)

	// Start search of files.
	findFiles(params["--path"])

	// Generate and save a new package.
	t := output.NewType(params["--package"], filepath.Join(basePath, "./views.go.template"))
	t.CreateDir(params["--output"])
	t.Extension = ".go" // Save generated file as a .go source.
	t.Context = map[string]interface{}{
		"files":    files,
		"rootPath": params["--path"],
	}
	t.Generate()
}

// initDefaults makes sure required parameters are not empty.
// And if they are default ones are used instead.
func initDefaults(params map[string]string) {
	// Check path to be scanned.
	if params["--path"] == "" {
		params["--path"] = "./views"
	}

	// Identify the dir where output will be stored.
	if params["--output"] == "" {
		params["--output"] = "./assets/views/"
	}

	// Define the package name of the output.
	if params["--package"] == "" {
		params["--package"] = "views"
	}
}

// findFiles starts a search of files. The result will be stored
// to global files variable.
func findFiles(path string) {
	rootPath = path
	filepath.Walk(path, walkFunc)
}

// walkFunc is a function that is used by findFiles for validating
// found paths.
func walkFunc(path string, info os.FileInfo, err error) error {
	// Make sure there are no any errors.
	if err != nil {
		log.Warn.Printf("An error occured while creating a listing: '%s'.", err)
		return err
	}

	// Assure it is not a directory.
	if info.IsDir() {
		return nil
	}

	// Get a file path without root path. So, path like "./views/accounts/index.html"
	// will be transformed into "accounts/index.html" if our root path is "./views".
	relPath := strings.TrimPrefix(path, rootPath)
	relPath = filepath.Clean(relPath)

	// Add files to the global files varible.
	log.Trace.Printf("Path '%s' discovered.", path)
	files[relPath] = path
	return nil
}

func init() {
	files = map[string]string{}
}
