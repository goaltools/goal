// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"os"
	"strings"

	"github.com/colegion/goal/internal/action"
	"github.com/colegion/goal/internal/generation"
	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/routes"
	"github.com/colegion/goal/utils/path"
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
	controllersImport, err := path.CleanImport(*input)
	assertNil(err)
	handlersImport, err := path.CleanImport(*output)
	assertNil(err)
	handlersDir, err := path.ImportToAbsolute(handlersImport)
	assertNil(err)

	log.Trace.Printf(`Processing "%s" package...`, controllersImport)
	ps.processPackage(controllersImport, routes.NewPrefixes())

	// Start generation of the handlers package.
	tpl, err := path.ImportToAbsolute("github.com/colegion/goal/tools/generate/handlers/handlers.go.template")
	assertNil(err)
	t := generation.NewType("", tpl)
	t.Extension = ".go" // Save generated files as a .go source.

	// Create a new directory for the package (the old one should be removed).
	log.Trace.Printf(`Starting generation of "%s" package...`, *pkg)
	t.CreateDir(handlersDir)

	// Iterate over all available controllers, generate a file for every of them.
	for i := range ps[controllersImport].list {
		// Set a name of the file. It is a lowercased controller name.
		t.Package = strings.ToLower(ps[controllersImport].list[i].Name)

		// Prepare parent controllers.
		pcs := ps[controllersImport].list[i].Parents.All(ps, "", newContext())

		// Set context variables.
		t.Context = map[string]interface{}{
			"package":    *pkg,
			"controller": ps[controllersImport].list[i],

			"parentControllers": pcs,

			"index":   i,
			"strconv": action.StrconvContext,
		}
		t.Generate()
	}
}

// assertNil makes sure an error is nil. It panics otherwise.
func assertNil(err error) {
	if err != nil {
		log.Error.Panic(err)
	}
}
