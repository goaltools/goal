// Package importpath provides functions for manipulating
// Go package import paths and target operating system compatible
// filesystem paths.
package importpath

import (
	"errors"
	"fmt"
	"go/build"
	"path"
	"path/filepath"
	"strings"
)

// gopaths stores a splitted and normalized version of GOPATHs
// so they can be safely used even though some platforms may use a mix
// of different kinds of separators.
var gopaths = func() (lst []string) {
	lst = filepath.SplitList(build.Default.GOPATH)
	for i := 0; i < len(lst); i++ {
		lst[i] = filepath.ToSlash(lst[i])
	}
	return
}()

// ToImport gets a standard operating system compatible path
// and transforms it into a valid Go import path. E.g., if one of
// the $GOPATHs is "/home/user/go" and the input argument is:
//	/home/user/go/src/github.com/anonymous/repo/
// or "pwd" shows "/home/user/go/src/github.com/" and the argument is:
//	./anonymous/repo/
// The result will be:
//	github.com/anonymous/repo
// NB: This function accepts *nix style paths everywhere. But other platform
// specific ones (e.g. Windows' "D:\path\to" like) are supported on their
// respective platforms only.
func ToImport(path string) (string, error) {
	// Path cannot be empty to make sure behaviour on
	// *nix matches to the one on other platforms.
	if path == "" {
		return "", errors.New("path is not defined")
	}

	// Make sure the input path is transformed into an absolute form.
	abs, err := absolute(path)
	if err != nil {
		return "", err
	}

	// Try to split the input absolute path into "$GOPATH/src" part
	// and a project path itself.
	paths, err := extractGopath(abs)
	if err != nil {
		return "", err
	}

	// Return a "$GOPATH/src"-less version of the path.
	// Make sure "/" are used as separators and there are no
	// trailing slashes.
	return strings.Trim(paths[1], "/"), nil
}

// ToPath gets a valid Go package import path and returns an associated absolute path.
// It relies on a "getwd" to detect which of the GOPATH's values to use.
// For example, if "pwd" returns "$GOPATH[0]/src/github.com" and the input import path is:
//	github.com/goaltools/importpath
// or:
//	./goaltools/importpath
// The result on *nix systems will be:
//	$GOPATH[0]/src/github.com/goaltools/importpath
// Path separators of the result path are platform specific.
func ToPath(imp string) (string, error) {
	// Transform the input import path into a
	// valid absolute import path if necessary.
	abs, err := Clean(imp)
	if err != nil {
		return "", err
	}

	// Find the GOPATH value of the current directory.
	curr, err := absolute(".")
	if err != nil {
		return "", err
	}
	gopath, err := extractGopath(curr)
	if err != nil {
		return "", err
	}

	// Concatenate the "GOPATH/src" with the import path
	// and return it.
	return filepath.Join(gopath[0], abs), nil
}

// Clean gets anything that resembles a Go import path: a relative or absolute one,
// with slashes or other alternative platform specific separators, with or without
// trailing slashes; and returns a valid absolute Go package import path.
// E.g., if "pwd" returns "$GOPATH[0]/src/github.com/goaltools/importpath" both:
//	github.com/goaltools/importpath/
// and:
//	../importpath
// will be transformed into:
//	github.com/goaltools/importpath
func Clean(imp string) (string, error) {
	// Normalize the input import path to make sure right slashes
	// are used as delimiters. Get rid of trailing slashes at the end.
	norm := strings.TrimRight(filepath.ToSlash(imp), "/")

	// Make sure path is not an empty string.
	if norm == "" {
		norm = "."
	}

	// If the import path is not relative and looks valid, return it as is.
	isRel := norm == "." || norm == ".." || strings.HasPrefix(norm, "./") || strings.HasPrefix(norm, "../")
	if !isRel && !strings.HasPrefix(norm, "/") {
		return norm, nil
	}

	// Otherwise, try to convert it to the regular form.
	return ToImport(norm)
}

// extractGopath gets an absolute directory path in a platform specific format
// and splits it into 2 parts:
// 1. The $GOPATH part.
// 2. "$GOPATH/src"-less part of the path.
// If provided path does not contain "$GOPATH/src", an error is returned.
func extractGopath(abs string) ([2]string, error) {
	// Normalize the input path.
	abs = filepath.ToSlash(abs)

	// Check every $GOPATH whether any of them is a parent directory
	// of the input absolute path.
	for i := 0; i < len(gopaths); i++ {
		// Prepare a "$GOPATH/src".
		srcDir := path.Join(gopaths[i], "src")

		// Check whether the source dir is the root
		// of the specified absolute path.
		if parent(srcDir, abs) {
			return [2]string{srcDir, abs[len(srcDir):]}, nil
		}
	}

	// If no import path returned so far, requested path is not
	// inside "$GOPATH/src".
	return [2]string{}, fmt.Errorf(`path "%s" is not inside of any of %v`, abs, gopaths)
}

// absolute is a function that gets a path and returns its absolute form.
// It is implemented as a variable rather than a direct inline call of the function
// for a better testability of this package.
// NB: The function automatically calls filepath.Clean on the input.
var absolute = filepath.Abs
