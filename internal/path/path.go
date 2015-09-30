// Package path is a wrapper around standard path, path/filepath,
// os, and go/build packages for work with paths and
// import paths.
package path

import (
	"errors"
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Path is a path in its normalized form
// (with unix style separators).
type Path struct {
	s   string // Normalized (unix like) form of a path.
	pkg bool   // Whether the path is an import path of a package.
}

// New allocates and returns a new Path.
func New(p string) *Path {
	return &Path{
		s: filepath.ToSlash(p),
	}
}

// String returns path as is if it is a package import path,
// or using a platform specific separators otherwise, e.g.
// as "path\\to\\smth" on Windows.
func (p *Path) String() string {
	if p.pkg {
		// Make sure path has no starting or ending slashes
		// (they are not allowed in package import paths).
		return strings.Trim(p.s, "/")
	}
	return filepath.FromSlash(p.s)
}

// Absolute returns path as an absolute one.
// It returns an error if current directory cannot
// be determined.
func (p *Path) Absolute() (*Path, error) {
	// If the path is already absolute, return it as is.
	if path.IsAbs(p.s) {
		return &Path{s: p.s}, nil
	}

	// Otherwise, join with the current directory.
	curr, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &Path{
		s: path.Join(filepath.ToSlash(curr), p.s),
	}, nil
}

// Import returns the path as a correct absolute package import path.
// For example, the path is:
//	../
// And current directory is "$GOPATH/src/user/project/x". Then the result is:
//	user/project
func (p *Path) Import() (*Path, error) {
	// If the path doesn't have "/", "./", or "../" at the beginning
	// it is already a valid go package path.
	relCurrPref := strings.HasPrefix(p.s, "./") || p.s == "."
	relBackPref := strings.HasPrefix(p.s, "../") || p.s == ".."
	absPref := strings.HasPrefix(p.s, "/")
	if !relCurrPref && !relBackPref && !absPref {
		return &Path{
			s:   p.s,
			pkg: true,
		}, nil
	}

	// Get an absolute form of the path.
	abs, err := p.Absolute()
	if err != nil {
		return nil, err
	}

	// Check every $GOPATH whether some of them
	// is a part of the path of absolute form.
	gopaths := filepath.SplitList(build.Default.GOPATH)
	for i := 0; i < len(gopaths); i++ {
		// Getting a normalized form (i.e. Unix style) of "$GOPATH/src".
		gopath := path.Join(filepath.ToSlash(gopaths[i]), "src")

		// Checking whether "$GOPATH/src" is a part of the absolute path.
		if res := strings.TrimPrefix(abs.s, gopath); res != abs.s {
			// Return the "$GOPATH/src"-less version of the path.
			return &Path{
				s:   res,
				pkg: true,
			}, nil
		}
	}

	// If no import path returned so far,
	// requested path is not inside "$GOPATH/src".
	return nil, fmt.Errorf(`path "%s" is not inside "$GOPATH/src"`, p.String())
}

// Package is an opposite of Import. It gets a valid package import path
// and returns its absolute path. E.g., there is input:
//	github.com/username/project
// It will output:
//	$GOPATH/src/github.com/username/project
// The first value in the list of GOPATHs is used.
func (p *Path) Package() (*Path, error) {
	imp, err := p.Import()
	if err != nil {
		return nil, err
	}

	// Make sure there is at least one $GOPATH value.
	gopaths := filepath.SplitList(build.Default.GOPATH)
	if len(gopaths) == 0 {
		return nil, errors.New("$GOPATH is not set")
	}

	// Join input path with the "$GOPATH/src" and return.
	// Make sure $GOPATH is normalized (i.e. unix style delimiters are used).
	return &Path{
		s: path.Join(filepath.ToSlash(gopaths[0]), "src", imp.String()),
	}, nil
}
