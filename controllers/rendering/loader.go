package rendering

import (
	"html/template"
	"path/filepath"

	"github.com/anonx/sunplate/log"
	p "github.com/anonx/sunplate/path"
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
	go log.Trace.Println("Loading templates...")

	// Iterating over all available template paths.
	for name, path := range templatePaths {
		// Find base for the current template
		// (either in the current dir or in one of the previous levels).
		var ok bool
		var base, cd string
		for {
			dir := p.Prefixless(filepath.Join(filepath.Base(name), cd), ".")
			cd += "../"

			// Check whether base template exists in the directory.
			n := filepath.Join(dir, BaseTemplate)
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

		// If the base was found, use it. Otherwise, go without it.
		var err error
		t := template.New(name).Funcs(Funcs).Delims(Delims.Left, Delims.Right)
		if base != "" {
			templates[name], err = t.ParseFiles(base, path)
			log.AssertNil(err)
			return
		}
		templates[name], err = t.ParseFiles(path)
		log.AssertNil(err)
	}
}
