package views

import (
	"fmt"
	"path/filepath"
	"strings"
)

// listing is a map of all found paths. Key is a string
// representation of directory path, values are files found in
// those directories.
type listing map[string][]path

// path represents a path to some file.
type path struct {
	Value []string // E.g. "path/to/smth" will be {"path", "to", "smth"}.
	Dir   bool     // Is it a path of a directory or a file.
}

// Name returns a name of the file / directory.
// File's name is the last segment of file with
// upper cased extension without a dot.
func (p path) Name() string {
	// Otherwise, get the last element of path (Base).
	n := p.Value[len(p.Value)-1]

	// Cut extension, translate it to upper case,
	// and concatenate again.
	ext := filepath.Ext(n)
	n = n[:len(n)-len(ext)]
	if len(ext) != 0 {
		ext = ext[1:] // Trim dot at the beginning
	}
	n += strings.ToUpper(ext)

	return n
}

// DirStructName returns a value that may be
// used for a name of a directory structure.
// Directory's name consists of segments
// concatenated without separators.
func (p path) DirStructName() string {
	if p.Dir {
		return strings.Join(p.Value, "")
	}
	return ""
}

// String returns a path of a file / directory as a string.
func (p path) String() string {
	return filepath.ToSlash(filepath.Join(p.Value...))
}

// addFile gets a path to some file as a string
// and registers it.
func (l listing) addFile(file string) {
	// Split path into segments and get a dir name (opposite of Base).
	ss := strings.Split(file, string(filepath.Separator))
	dir := filepath.ToSlash(filepath.Join(ss[:len(ss)-1]...)) // If path is "path/to/smth.txt", use "path/to".

	// Check whether such path already exists.
	if _, ok := l[dir]; !ok {
		// If not, initialize it.
		l[dir] = []path{}
	}

	// Add the file to listing.
	l[dir] = append(l[dir], path{
		Value: ss,
		Dir:   false,
	})
}

// addDir is identical to addFile except it is for
// directories rather than files.
func (l listing) addDir(dir string) {
	// Make sure path is not empty.
	if dir == "" || dir == "." {
		return
	}

	// Split path into segments and get a root dir name (opposite of Base).
	ss := strings.Split(dir, string(filepath.Separator))
	root := filepath.ToSlash(filepath.Join(ss[:len(ss)-1]...)) // If path is "path/to/dir", use "path/to".

	// Check whether such path already exists.
	if _, ok := l[root]; !ok {
		// If not, initialize it.
		l[root] = []path{}
	}

	// Add the directory to the list.
	l[root] = append(l[root], path{
		Value: ss,
		Dir:   true,
	})
}

// FilePaths returns a map of file name - file path pairs.
// Example of a name is "Path.To.FileHTML".
// It is used in template for generation of initialization
// function.
func (l listing) FilePaths() map[string]string {
	fs := map[string]string{}
	for _, p := range l.paths() {
		// Make sure the path does not belong to directory.
		if p.Dir {
			continue
		}

		// Prepare a name of the variable that will be
		// used for storing file's path.
		n := p.Name()
		if len(p.Value) > 1 {
			n = fmt.Sprintf(
				"%s.%s",
				strings.Join(p.Value[:len(p.Value)-1], "."),
				n,
			)
		}

		fs[n] = p.String()
	}
	return fs
}

// Dirs a map of directory name - a list of paths inside
// pairs. Examples of names are: "App", "AppProfile", etc.
// It is used in template to generate types for directories.
// Dirs does not return root directory ("./").
func (l listing) Dirs() map[string][]path {
	ds := map[string][]path{}
	for _, p := range l.paths() {
		// Ignore files.
		if !p.Dir {
			continue
		}

		// Add the current directory and its
		// children.
		ds[p.DirStructName()] = l[p.String()]
	}
	return ds
}

// RootAssets returns a list of directories and files
// that are found at the root of requested path.
// This method is used in template.
func (l listing) RootAssets() (ps []path) {
	return l[""]
}

// paths returns all paths of listing in a flat form:
// as a slice.
func (l listing) paths() (ps []path) {
	for _, paths := range l {
		ps = append(ps, paths...)
	}
	return
}
