// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"path/filepath"
	"strings"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/path"
	"github.com/anonx/sunplate/strconv"
)

// Mappings of supported types and reflect functions.
var strconvContext = strconv.Context()

// Start is an entry point of the generate handlers command.
func Start(params command.Data) {
	inputDir := params.Default("--input", "./controllers")
	outputDir := params.Default("--output", "./assets/handlers")
	outPkg := params.Default("--package", "handlers")

	// Start processing of controllers.
	ps := packages{}
	absImport := path.AbsoluteImport(inputDir)
	absImportOut := path.AbsoluteImport(outputDir)
	ps.processPackage(absImport)

	// Start generation of handler packages.
	t := output.NewType(
		"", filepath.Join(path.SunplateDir("generation", "handlers"), "./handlers.go.template"),
	)
	t.Extension = ".go" // Save generated files as a .go source.

	// Iterate through all available packages and generate handlers for them.
	for imp := range ps {
		// Check whether current package is the main one
		// and should be stored at the root directory or it is a subpackage.
		//
		// I.e. if --input is "./controllers" and --output is "./assets/handlers",
		// we are saving processed "./controllers" package to "./assets/handlers"
		// and some "github.com/anonx/smth" to "./assets/handlers/github.com/anonx/smth".
		out := outputDir
		if imp != absImport {
			out = filepath.Join(out, imp)
		}
		t.CreateDir(out)

		// Iterate over all available controllers, generate handlers package on
		// every of them.
		for name := range ps[imp] {
			// Find parent controllers of this controller.
			cs := []parent{}
			for i, p := range ps[imp][name].Parents {
				// Make sure it is a controller rather than just some embedded struct.
				if _, ok := ps[p.Import]; ok {
					cs = append(cs, parent{
						ID:     i,
						Import: p.Import,
						Name:   p.Name,
					})
				}
			}

			// Initialize parameters and generate a package.
			t.Package = strings.ToLower(name)
			t.Context = map[string]interface{}{
				"after":   magicActionAfter,
				"before":  magicActionBefore,
				"finally": magicActionFinally,

				"controller":   ps[imp][name],
				"file":         filepath.Base(ps[imp][name].File),
				"import":       imp,
				"input":        inputDir,
				"name":         name,
				"outputImport": absImportOut,
				"output":       outputDir,
				"package":      outPkg,
				"parents":      cs,

				"strconv": strconvContext,
			}
			t.Generate()
		}
	}
}
