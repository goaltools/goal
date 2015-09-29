package create

import (
	"os"
	"path/filepath"
)

// sourceFiles contains extensions of files that should be process
// rather than just copied.
// As an example, "github.com/colegion/goal/internal/skeleton"
// inside such files must be replaced by their new import path.
var sourceFiles = []string{".go", ".yml"}

// result represents objects found when scanning a skeleton directory.
// There are a few possible kind of them: directories, static files,
// source files (that require additional processing of their content).
type result struct {
	dirs        []string
	files, srcs []paths
}

// paths represents paths to an object in filesystem.
type paths struct {
	absolute, relative string
}

// walk scans a requested directory and returns found objects as
// three different categories: directories, files, and source files
// that require additional processing.
func walk(dest string) (*result, error) {
	res, fn := walkFunc(dest)
	err := filepath.Walk(dest, fn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// walkFunc returns an allocated result and a function that may be used for validation
// of found elements. Successfully validated ones are stored to the returned result.
func walkFunc(dest string) (*result, func(string, os.FileInfo, error) error) {
	res := &result{
		dirs:  []string{},
		files: []paths{},
		srcs:  []paths{},
	}

	return res, func(p string, info os.FileInfo, err error) error {
		// Make sure there are no any errors.
		if err != nil {
			return err
		}

		// Get filepath without the dir path at the beginning.
		// So, when we are scanning "/home/user/app/stuff"
		// we would get "app/stuff" instead.
		// Since we are always scanning inside "dest" an error
		// should never occur, so ignoring it.
		rel, _ := filepath.Rel(dest, p)

		// Check whether current element is a directory.
		if info.IsDir() {
			res.dirs = append(res.dirs, rel)
			return nil
		}

		// Find out whether it is a static file or a go / some other source.
		ext := filepath.Ext(p)
		for i := 0; i < len(sourceFiles); i++ {
			if ext == sourceFiles[i] {
				res.srcs = append(res.srcs, paths{
					absolute: p,
					relative: rel,
				})
				return nil
			}
		}

		// If it is a static file, add it to the list.
		res.files = append(res.files, paths{
			absolute: p,
			relative: rel,
		})
		return nil
	}
}
