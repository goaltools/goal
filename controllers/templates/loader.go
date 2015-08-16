package templates

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/anonx/sunplate/log"
)

var templates = map[string]*template.Template{}

// SetTemplatePaths gets a path to templates directory
// and a list of files, and then loads them.
// You may use it in your main.go as follows:
//	package main
//
//	import (
//		"github.com/user/project/assets/views"
//
//		"github.com/anonx/sunplate/controllers/results"
//	)
//
//	// ...
//
//	func init() {
//		results.SetTemplatePaths(views.Root, views.List)
//	}
func SetTemplatePaths(root string, templatePaths []string) {
	log.Trace.Println("Loading templates...")

	// Iterating over all available template paths.
	for _, path := range templatePaths {
		// Find base for the current template
		// (either in the current dir or in one of the previous levels).
		var base, cd string
		for {
			b := filepath.Base(path)
			dir := filepath.Join(path[:len(path)-len(b)], cd)
			cd += "../"

			// Check whether this template is a base. If so, do not load
			// any other bases.
			if b == BaseTemplate {
				break
			}

			// Check whether base template exists in the directory.
			base = filepath.Join(dir, BaseTemplate)
			if _, ok := templates[base]; ok || contains(templatePaths, base) {
				break
			}
			base = ""

			// Check whether we have unsuccessfully achieved the top level
			// of the path.
			if strings.HasPrefix(dir, "../") {
				break
			}
		}

		log.Trace.Printf("\t%s (%s)", path, base)

		// If the base was found, use it. Otherwise, go without it.
		var err error
		t := template.New(path).Funcs(Funcs).Delims(Delims.Left, Delims.Right)
		if base != "" {
			templates[path], err = t.ParseFiles(
				filepath.Join(root, base),
				filepath.Join(root, path),
			)
			showError(root, base, path, err)
			continue
		}
		templates[path], err = t.ParseFiles(filepath.Join(root, path))
		showError(root, base, path, err)
	}
}

// contains returns true if a requested value found
// in the requested slice of strings.
func contains(arr []string, value string) bool {
	for i := range arr {
		if arr[i] == value {
			return true
		}
	}
	return false
}

// showErrors writes an error to log.
func showError(root, base, path string, err error) {
	if err == nil {
		return
	}
	pwd, _ := os.Getwd()
	log.Error.Panicf(
		`Cannot parse "%s" with "%s" as a base template (pwd "%s"). Error: %v.`,
		filepath.Join(root, path), filepath.Join(root, base), pwd, err,
	)
}
