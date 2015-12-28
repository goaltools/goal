// Package path is a wrapper around standard path, path/filepath,
// os, and go/build packages for work with paths and
// import paths.
package path

import (
	"fmt"
	"go/build"
	"path"
	"path/filepath"
	"strings"
)

// AbsoluteToImport gets an absolute path and tryes to transform it into
// a valid package import path. E.g., if $GOPATH is "/home/user/go" then the path:
//	/home/user/go/src/github.com/colegion/goal
// must be transformed into:
//	github.com/colegion/goal
// The path must be within "$GOPATH/src", otherwise an error will be returned.
func AbsoluteToImport(abs string) (string, error) {
	// Make sure the input path is in fact absolute.
	if !filepath.IsAbs(abs) {
		return "", fmt.Errorf(`absolute path expected, got "%s"`, abs)
	}

	// Check every $GOPATH whether some of them is a prefix of the input path.
	// That would mean the input path is within $GOPATH.
	gopaths := filepath.SplitList(build.Default.GOPATH)
	for i := 0; i < len(gopaths); i++ {
		// Getting a "$GOPATH/src".
		gopath := filepath.Join(gopaths[i], "src")

		// Checking whether "$GOPATH/src" is a prefix of the input path.
		if res := strings.TrimPrefix(abs, gopath); res != abs {
			// Return the "$GOPATH/src"-less version of the path.
			// Make sure "/" are used as separators and there are no
			// leading or trailing slashes.
			return strings.Trim(filepath.ToSlash(res), "/"), nil
		}
	}

	// If no import path returned so far, requested path is not inside "$GOPATH/src".
	return "", fmt.Errorf(`path "%s" is not inside "$GOPATH/src"`, abs)
}

// ImportToAbsolute gets a valid package import path and tries to transform
// it into an absolute path. E.g., there is an input:
//	github.com/username/project
// It will output:
//	$GOPATH/src/github.com/username/project
// NOTE: The first value from the list of GOPATHs is always used.
func ImportToAbsolute(imp string) (string, error) {
	// Make sure the input import path is not relative.
	var err error
	imp, err = CleanImport(imp)
	if err != nil {
		return "", err
	}

	// Replace the "/" by the platform specific separators.
	p := filepath.FromSlash(imp)

	// Make sure the path is not a valid absolute path.
	if filepath.IsAbs(p) {
		return p, nil
	}

	// Split $GOPATH list to use the first value.
	gopaths := filepath.SplitList(build.Default.GOPATH)

	// Join input path with the "$GOPATH/src" and return.
	// Make sure $GOPATH is normalized (i.e. unix style delimiters are used).
	return path.Join(gopaths[0], "src", p), nil
}

// CleanImport gets a package import path and returns it as is if it is absolute.
// Otherwise, it tryes to convert it to an absolute form.
func CleanImport(imp string) (string, error) {
	// If the path is not relative, return it as is.
	if imp != "." && imp != ".." &&
		!filepath.HasPrefix(imp, "./") && !filepath.HasPrefix(imp, "../") {

		// Get rid of trailing slashes.
		return strings.TrimRight(imp, "/"), nil
	}

	// Find a full absolute path to the requested import.
	abs, err := filepath.Abs(imp)
	if err != nil {
		return "", err
	}

	// Extract package's import from it.
	return AbsoluteToImport(abs)
}
