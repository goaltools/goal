// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goaltools/goal/internal/action"
	"github.com/goaltools/goal/internal/generation"
	"github.com/goaltools/goal/internal/log"
	"github.com/goaltools/goal/internal/routes"

	"github.com/goaltools/importpath"
)

// start is an entry point of the generate handlers command.
func start() {
	// Clean the out directory.
	log.Trace.Printf(`Removing "%s" directory if already exists...`, *output)
	err := os.RemoveAll(*output)
	if err != nil {
		log.Error.Panic(err)
	}

	// Start processing of controllers.
	ps := packages{}
	absImport, err := importpath.ToImport(*input)
	if err != nil {
		log.Error.Panic(err)
	}
	absImportOut, err := importpath.ToImport(*output)
	if err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Printf(`Processing "%s" package...`, absImport)
	ps.processPackage(absImport, routes.NewPrefixes())

	// Start generation of handler packages.
	tpl, err := importpath.ToPath("github.com/goaltools/goal/tools/generate/handlers/handlers.go.template")
	if err != nil {
		log.Error.Panic(err)
	}
	t := generation.NewType("", tpl)
	t.Extension = ".go" // Save generated files as a .go source.

	// Iterate through all available packages and generate handlers for them.
	// TODO: refactor this fragment. Consider use of fmt.Sprintf instead of html/template.
	log.Trace.Printf(`Starting generation of "%s" package...`, *pkg)
	for imp := range ps {
		// Check whether current package is the main one
		// and should be stored at the root directory or it is a subpackage.
		//
		// I.e. if --input is "./controllers" and --output is "./assets/handlers",
		// we are saving processed "./controllers" package to "./assets/handlers"
		// and some it imports "github.com/goaltools/smth" to "./assets/handlers/github.com/goaltools/smth".
		out := *output
		if imp != absImport {
			out = filepath.Join(out, filepath.FromSlash(imp))
		}
		t.CreateDir(out)

		// Iterate over all available controllers, generate handlers package on
		// every of them.
		n := 0
		for name := range ps[imp].data {
			// Find parent controllers of this controller.
			cs := []parent{}
			for i, p := range ps[imp].data[name].Parents {
				// Make sure it is a controller rather than just some embedded struct.
				check := p.Import
				if check == "" { // Embedded parent is a local structure.
					check = absImport
				}
				if _, ok := ps[check]; !ok { // Such package is not in the list of scanned ones.
					continue
				}
				if _, ok := ps[check].data[p.Name]; !ok { // There is no such controller.
					continue
				}

				// It is a valid parent controller, add it to the list.
				cs = append(cs, parent{
					ID:     i,
					Import: p.Import,
					Name:   p.Name,
				})
			}

			// Initialize parameters and generate a package.
			t.Package = strings.ToLower(name)
			t.Context = map[string]interface{}{
				"after":  action.MethodAfter,
				"before": action.MethodBefore,

				"controller":   ps[imp].data[name],
				"controllers":  ps[imp].data,
				"import":       imp,
				"input":        input,
				"name":         name,
				"outputImport": absImportOut,
				"output":       output,
				"package":      pkg,
				"parents":      cs,
				"initFunc":     ps[imp].init,
				"num":          n,

				"actionImport":    action.InterfaceImport,
				"actionInterface": action.Interface,
				"strconv":         action.StrconvContext,
			}
			t.Generate()
			n++
		}
	}
}
