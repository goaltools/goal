package rendering

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/anonx/sunplate/log"
)

var templates = map[string]*template.Template{}

// SetTemplatePaths defines a map of templates.
// Template names are represented as keys and their paths as values.
// Initialize it from your init.go as follows:
//	import (
//		"github.com/user/project/assets/views"
//		"github.com/anonx/sunplate/controllers/rendering"
//	)
//
//	type Controller struct {
//		// ...
//		rendering.Template
//	}
//
//	func init() {
//		rendering.SetTemplatePaths(views.Context)
//	}
func SetTemplatePaths(templatePaths map[string]string) {
	log.Trace.Println("Loading templates...")

	// Iterating over all available template paths.
	for name, path := range templatePaths {
		// Find base for the current template
		// (either in the current dir or in one of the previous levels).
		var ok bool
		var base, cd, n string
		for {
			b := filepath.Base(name)
			dir := filepath.Join(name[:len(name)-len(b)], cd)
			cd += "../"

			// Check whether this template is a base. If so, do not load
			// any other bases.
			if b == BaseTemplate {
				break
			}

			// Check whether base template exists in the directory.
			n = filepath.Join(dir, BaseTemplate)
			base, ok = templatePaths[n]
			if ok {
				break
			}

			// Check whether we have unsuccessfully achieved the top level
			// of the path.
			if dir == "" {
				break
			}
		}

		log.Trace.Printf("\t%s (%s)", name, n)

		// If the base was found, use it. Otherwise, go without it.
		var err error
		t := template.New(name).Funcs(Funcs).Delims(Delims.Left, Delims.Right)
		if base != "" {
			templates[name], err = t.ParseFiles(base, path)
			showError(base, path, err)
			continue
		}
		templates[name], err = t.ParseFiles(path)
		showError(base, path, err)
	}
}

// showErrors writes an error to log.
func showError(base, path string, err error) {
	if err == nil {
		return
	}
	pwd, _ := os.Getwd()
	log.Error.Panicf(
		`Cannot parse "%s" with "%s" as a base template (pwd "%s"). Error: %v.`,
		path, base, pwd, err,
	)
}
