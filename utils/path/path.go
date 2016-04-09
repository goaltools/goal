// Package path is a wrapper around standard path, path/filepath,
// os, and go/build packages for work with paths and
// import paths.
package path

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetCurrentGoPath() (string, error) {
	abs, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}

	// Check every $GOPATH whether some of them is a prefix of the input path.
	// That would mean the input path is within $GOPATH.
	gopaths := filepath.SplitList(build.Default.GOPATH)
	for i := 0; i < len(gopaths); i++ {
		// Getting a "$GOPATH/src".
		gopath := filepath.Clean(filepath.Join(gopaths[i], "src"))

		// Checking whether "$GOPATH/src" is a prefix of the input path.
		if strings.HasPrefix(abs, gopath) {
			return abs, nil
		}
	}

	// If no import path returned so far, requested path is not inside "$GOPATH/src".
	return "", fmt.Errorf(`current path "%s" is not inside "$GOPATH/src"`, abs)
}

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

	abs = filepath.Clean(abs)
	// Check every $GOPATH whether some of them is a prefix of the input path.
	// That would mean the input path is within $GOPATH.
	gopaths := filepath.SplitList(build.Default.GOPATH)
	for i := 0; i < len(gopaths); i++ {
		// Getting a "$GOPATH/src".
		gopath := filepath.Clean(filepath.Join(gopaths[i], "src"))

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
	// If the path is not relative, return it as is.
	if IsRelativePath(imp) {
		abs, err := filepath.Abs(imp)
		if nil != err {
			return "", err
		}
		// Get rid of trailing slashes.
		return filepath.Clean(abs), nil
	}

	// Replace the "/" by the platform specific separators.
	p := filepath.FromSlash(filepath.ToSlash(imp))

	// Make sure the path is not a valid absolute path.
	if filepath.IsAbs(p) {
		return p, nil
	}

	// Split $GOPATH list to use the first value.
	gopaths := filepath.SplitList(build.Default.GOPATH)

	for i := 0; i < len(gopaths); i++ {
		// Join input path with the "$GOPATH/src" and return.
		// Make sure $GOPATH is normalized (i.e. unix style delimiters are used).
		abs := filepath.Clean(path.Join(gopaths[i], "src", p))
		if _, e := os.Stat(abs); nil == e {
			return abs, nil
		}
	}

	// If no import path returned so far, requested path is not inside "$GOPATH/src".
	return "", fmt.Errorf(`package "%s" is not inside "$GOPATH/src"`, imp)
}

// CleanImport gets a package import path and returns it as is if it is absolute.
// Otherwise, it tryes to convert it to an absolute form.
func CleanImport(imp string) (string, error) {
	if "" == imp {
		return "", nil
	}
	// If the path is not relative, return it as is.
	if IsRelativePath(imp) {
		// Find a full absolute path to the requested import.
		abs, err := GetCurrentGoPath()
		if err != nil {
			return "", err
		}
		return AbsoluteToImport(filepath.Join(abs, imp))
	}

	if filepath.IsAbs(imp) {
		// Extract package's import from it.
		return AbsoluteToImport(imp)
	}

	abs, err := ImportToAbsolute(imp)
	if err != nil {
		return "", err
	}

	return AbsoluteToImport(abs)
}

func IsRelativePath(imp string) bool {
	impNorm := filepath.ToSlash(imp)
	return impNorm == "." || impNorm == ".." ||
		filepath.HasPrefix(impNorm, "./") ||
		filepath.HasPrefix(impNorm, "../")
}
