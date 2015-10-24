// Package listing is a subcommand of generate that is
// used for scanning directories and composing a list
// of found files. It may be useful to statically check
// names of used templates.
// I.e. instead of using paths of templates in actions
// directely like `RenderTemplate("path/to/template.html")`,
// it is possible to do something like
// `RenderTemplate(views.Path.To.TemplateHTML)`.
package listing

import (
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/colegion/goal/internal/generation"
	"github.com/colegion/goal/utils/log"
	p "github.com/colegion/goal/utils/path"
	"github.com/colegion/goal/utils/tool"
)

var fileNamePattern = regexp.MustCompile(
	"^[A-Za-z]{1}\\w*[.]{0,1}\\w*$",
)

// Handler is an instance of "generate listing" subcommand (tool).
var Handler = tool.Handler{
	Run: main,

	Name:  "generate listing",
	Usage: "[flags]",
	Info:  "generate a listing of files and directories",
	Desc: `Tool "generate listing" scans the directory you request
and generates a package with file names that were found.
So, you can import and use the generated package as follows:

	listing.Paths.Some.Path.To.FileHTML

The variable above is a string containing a path to that file.
The advantage of this approach is type safety. If this tool is started
every time you modify your files it will garantee you are not using
a template that does not exist.
`,
}

var input, output, pkg *string

// main is an entry point of listing subcommand.
func main(hs []tool.Handler, i int, args tool.Data) {
	// Prepare a path to scan.
	path, err := p.ImportToAbsolute(*input)
	log.AssertNil(err)

	// Start search of files.
	fs, fn := walkFunc(path)
	filepath.Walk(path, fn)

	// Generate and save a new package.
	tpl, err := p.ImportToAbsolute("github.com/colegion/goal/tools/generate/listing/listing.go.template")
	log.AssertNil(err)
	t := generation.NewType(*pkg, tpl)
	t.CreateDir(*output)
	t.Extension = ".go" // Save generated file as a .go source.
	t.Context = map[string]interface{}{
		"listing": fs,
		"input":   filepath.ToSlash(*input),
	}
	t.Generate()

	// Generate and save ini config with file names.
	tpl, err = p.ImportToAbsolute("github.com/colegion/goal/tools/generate/listing/listing.ini.template")
	log.AssertNil(err)
	t = generation.NewType(*pkg, tpl)
	t.Path = *output
	t.Extension = ".ini"
	t.Context = map[string]interface{}{
		"listing": fs,
		"input":   filepath.ToSlash(*input),
	}
	t.Generate()
}

// walkFunc returns a files listing and a function that may be used for validation
// of found files. Successfully validated ones are stored to the listing variable.
func walkFunc(dir string) (listing, func(string, os.FileInfo, error) error) {
	l := listing{}

	return l, func(path string, info os.FileInfo, err error) error {
		// Make sure there are no any errors.
		if err != nil {
			log.Warn.Printf(`An error occured while creating a listing: "%s".`, err)
			return err
		}

		// Get filepath without the dir path at the beginning.
		// So, when we are scanning "views/app/index.html" our generated
		// result will be "app/index.html" instead.
		rel, _ := filepath.Rel(dir, path)

		// Make sure file name is of supported format.
		ss := strings.Split(rel, string(filepath.Separator))
		for _, s := range ss {
			// Ignore root directory.
			if s == "." {
				continue
			}

			// Check whether file / directory name is supported.
			if !fileNamePattern.MatchString(s) {
				log.Warn.Printf(`"%s" is ignored as "%s" is of unsupported format.`, rel, s)
				return fmt.Errorf(`"%s" is of unsupported type`, s)
			}

			// Notify user if file / directory's name is not exported.
			if !ast.IsExported(s) {
				log.Trace.Printf(`"%s" will not be exported as "%s" starts with a lower case letter.`, rel, s)
			}
		}

		// Add the directory to the list (if it is a dir).
		if info.IsDir() {
			l.addDir(rel)
			return nil
		}

		// Otherwise, register the file.
		l.addFile(rel)
		return nil
	}
}

func init() {
	input = Handler.Flags.String("input", "./views", "a path to directory with files to scan")
	output = Handler.Flags.String("output", "./assets/views", "a directory where generated package must be saved")
	pkg = Handler.Flags.String("package", "views", "name of the package to generate")
}
