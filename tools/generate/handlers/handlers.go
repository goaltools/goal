// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"fmt"
	"os"
	"path/filepath"
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
	if err != nil {
		if !path.IsRelativePath(*output) {
			log.Error.Panic(err)
		}

		// Get rid of trailing slashes.
		handlersDir, err = filepath.Abs(*output)
		if err != nil {
			log.Error.Panic(err)
		}
	}

	log.Trace.Printf(`Processing "%s" package...`, controllersImport)
	ps.processPackage(controllersImport, routes.NewPrefixes())

	// Start generation of the handlers package.
	tplInit, err := path.ImportToAbsolute("github.com/colegion/goal/tools/generate/handlers/init.go.template")
	assertNil(err)
	tplHandlers, err := path.ImportToAbsolute("github.com/colegion/goal/tools/generate/handlers/handlers.go.template")
	assertNil(err)
	t := generation.NewType("", tplInit)
	t.Extension = ".go" // Save generated files as a .go source.

	// Create a new directory for the package (the old one should be removed).
	log.Trace.Printf(`Starting generation of "%s" package...`, *pkg)
	t.CreateDir(handlersDir)

	// Generate an init file.
	n := "Init"
	check := map[string]bool{
		n: true,
	}
	t.Package = strings.ToLower(n)
	t.Context = map[string]interface{}{
		"package": *pkg,
		"inits":   ps.AllInits(controllersImport),
		"routes":  ps.AllRoutes(),
	}
	t.Generate()

	// Iterate over all available controllers, generate a file for every of them.
	index := 0
	t = generation.NewType("", tplHandlers)
	t.Path = handlersDir
	t.Extension = ".go" // Save generated files as a .go source.
	for k := range ps {
		for i := range ps[k].list {
			// Use controller's name as a file name if it is unique.
			// Add an integer suffix otherwise.
			n = ps[k].list[i].Name
			if _, ok := check[n]; ok {
				n = fmt.Sprintf("%s%d", n, index)
			}
			check[n] = true

			// Set a name of the file. It is a lowercased controller name.
			t.Package = strings.ToLower(n)

			// Prepare parent controllers.
			pcs := ps[k].list[i].Parents.All(ps, "", newContext())

			// Set context variables.
			t.Context = map[string]interface{}{
				"package": *pkg,

				"controllers": ps[k].list,

				"controller": ps[k].list[i],
				"import":     k,

				"name":               n,
				"controllerFileName": filepath.Base(ps[k].list[i].File),

				"inits":             ps.AllInits(controllersImport),
				"parentControllers": pcs,

				"index":   index,
				"strconv": action.StrconvContext,
			}
			t.Generate()
			index++
		}
	}
}

// assertNil makes sure an error is nil. It panics otherwise.
func assertNil(err error) {
	if err != nil {
		log.Error.Panic(err)
	}
}
