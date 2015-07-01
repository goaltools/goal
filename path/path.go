// Package path is a set of helper functions for work with paths in sunplate package.
package path

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/anonx/sunplate/log"
)

// SunplateDir returns an absolute path of the directory where a package of sunplate
// toolkit is installed (it should be located at "$GOPATH/github.com/anonx/sunplate/$pkg").
//
// For example, to find out the path of sunplate's root directory use:
//	path.SunplateDir()
// To identify the path of sunplate/generation/handlers subpackage do the following:
//	path.SunplateDir("generation", "handlers")
func SunplateDir(pkgs ...string) string {
	p := filepath.Join(
		build.Default.GOPATH, "src/github.com/anonx/sunplate",
	)
	for i := range pkgs {
		p = filepath.Join(p, pkgs[i])
	}
	return p
}

// WorkingDir returns a path to the directory where sunplate program was run.
// It panics in case of error.
//
// So, if we are moving to some directory and starting sunplate program there:
//	cd /home/user/somedir
//	sunplate ...
// the WorkingDir() should return "/home/user/somedir"
//	path := WorkingDir() // Output: "/home/user/somedir"
func WorkingDir() string {
	p, err := os.Getwd()
	log.AssertNil(err)
	return p
}

// AbsoluteImport gets an import path and returns its absolute representation.
// So, if it is already absolute it is returned as is.
// Otherwise, it's assumed the path is relative to the current working directory.
// If something goes wrong AbsoluteImport panics.
func AbsoluteImport(path string) string {
	// If it is an absolute path, return it as is.
	if !strings.HasPrefix(path, "./") {
		return path
	}

	// Otherwise, try to convert the path to absolute.
	// Get rid of "$GOPATH/src" and "/" (that +1) part at the beginning.
	p := filepath.Join(WorkingDir(), path)
	return p[len(filepath.Join(build.Default.GOPATH, "src"))+1:]
}
