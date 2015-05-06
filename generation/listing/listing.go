// Package listing is a generate subcommand that scans
// requested directory and generates a .go file with
// a list of all found files.
package listing

import (
	"os"
	"path/filepath"
)

var files []string

// Start is an entry point of listing subcommand.
func Start(params map[string]string) {
	// Initialize missed parameters.
	initDefaults(params)

	// Start search of files.
	findFiles(params["--path"])
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
		params["--output"] = "./assets/"
	}

	// Define the package name of the output.
	if params["--package"] == "" {
		params["--package"] = "views"
	}
}

// findFiles starts a search of files. The result will be stored
// to global files variable.
func findFiles(path string) {
	filepath.Walk(path, walkFunc)
}

// walkFunc is a function that is used by findFiles for validating
// found paths.
func walkFunc(path string, info os.FileInfo, err error) error {
	// Make sure there are no any errors.
	if err != nil {
		return err
	}

	// Assure it is not a directory.
	if info.IsDir() {
		return nil
	}

	// Add files to the global files varible.
	files = append(files, path)
	return nil
}
