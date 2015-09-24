// Package path is a set of helper functions for work with paths in goal package.
package path

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/colegion/goal/log"
)

// goalDir returns an absolute path of the directory where a package of goal
// toolkit is installed (it should be located at "$GOPATH/github.com/colegion/goal/$pkg").
//
// For example, to find out the path of goal's root directory use:
//	path.goalDir()
// To identify the path of goal/generation/handlers subpackage do the following:
//	path.goalDir("generation", "handlers")
func goalDir(pkgs ...string) string {
	p := filepath.Join(
		PackageDir(""), goalImport(pkgs...),
	)
	return filepath.ToSlash(p)
}

// goalImport is equivalent of the goalDir except it returns
// a go import path rather than a path to a directory.
func goalImport(pkgs ...string) string {
	p := "github.com/colegion/goal"
	for i := range pkgs {
		p = filepath.Join(p, pkgs[i])
	}
	return filepath.ToSlash(p)
}

// WorkingDir returns a path to the directory where goal program was run.
// It panics in case of error.
//
// So, if we are moving to some directory and starting goal program there:
//	cd /home/user/somedir
//	goal ...
// the WorkingDir() should return "/home/user/somedir"
//	path := WorkingDir() // Output: "/home/user/somedir"
func WorkingDir() string {
	p, err := os.Getwd()
	log.AssertNil(err)
	return filepath.ToSlash(p)
}

// AbsoluteImport gets an import path and returns its absolute representation.
// So, if it is already absolute or empty it is returned as is.
// Otherwise, it's assumed the path is relative to the current working directory.
// If something goes wrong AbsoluteImport panics.
func AbsoluteImport(path string) string {
	// If path is empty, do nothing.
	if path == "" {
		return path
	}

	// If it is an absolute path, remove the starting slashes, if any.
	path = filepath.ToSlash(path)
	if !strings.HasPrefix(path, ".") {
		return strings.TrimLeft(path, "/")
	}

	// Get absolute import path by removing "$GOPATH/src" at the beginning
	// of working dir + relative path.
	gopath := PackageDir("")
	pkgPath := filepath.ToSlash(filepath.Join(WorkingDir(), path))
	if !strings.HasPrefix(pkgPath, gopath) { // If there is no $GOPATH at the beginning.
		log.Error.Panicf("Your project must be located inside of $GOPATH.")
	}
	imp := Prefixless(pkgPath, gopath)

	// Import path never starts with a slash.
	return Prefixless(imp, "/")
}

// PackageDir gets a golang import path and returns its full path.
func PackageDir(imp string) string {
	gopaths := filepath.SplitList(build.Default.GOPATH)
	return filepath.ToSlash(
		filepath.Join(gopaths[0], "src", imp), // We are always using the first GOPATH in a list.
	)
}

// Prefixless cuts a prefix of a path and returns the result
// that is cleaned.
func Prefixless(path, prefix string) string {
	return filepath.ToSlash(
		filepath.Clean(strings.TrimPrefix(path, prefix)),
	)
}
