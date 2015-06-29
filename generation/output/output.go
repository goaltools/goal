// Package output is used for generation and saving files.
// For example, handlers, routes, etc.
package output

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/anonx/sunplate/log"
)

// Type is a context that stores information that is used for generation
// and saving files (mostly go packages).
type Type struct {
	// Path is a directory where generated file will be stored.
	// Example: ./assets/routes/
	// It is expected to be relative to the root of a project.
	Path string

	// Package is a name of a generated file without extension.
	// It may be used as a package name in case of generating go files.
	Package string

	// Extension is added to the end of name of a generated file.
	// So file will be saved to:
	//	filepath.Join(Path, Package+Extension)
	Extension string

	// Name of the template that is used by this Type.
	TemplateName string

	// Template is a skeleton of file that has to be generated.
	Template *template.Template

	// Context is used for passing data to the Template.
	Context map[string]interface{}
}

// NewType reads the requested template and returns an output.Type
// with initialized Template field.
// << and >> are used as delimiters.
func NewType(pkg, templatePath string) Type {
	n := filepath.Base(templatePath)
	t, err := template.New(n).Delims("<@", ">"). // Use <@ and > as delimiters.
							Funcs(funcs).ParseFiles(templatePath)
	if err != nil {
		log.Error.Panicf("Didn't manage to open template '%s', error: '%s'.", templatePath, err)
	}
	return Type{
		Package:      pkg,
		TemplateName: n,
		Template:     t,
	}
}

// CreateDir initializes output.Type.Path with the requested path
// and tries to create it in filesystem if it doesn't exist yet.
// It panics in case of error.
func (t *Type) CreateDir(path string) {
	t.Path = path

	// Check whether directory already exists.
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return
	}

	// If not, try to create it.
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Error.Panicf("Cannot create directory '%s', error: '%s'.", path, err)
	}
}

// Generate creates a file with a name specified in Type.Package and Type.Extension
// in the location defined in Type.Path and with the content defined by Type.Template.
// The output directory should be created in advance. It's possible to do it using:
//	CreateDir("./path/to/output/")
func (t *Type) Generate() {
	// Generate a template file.
	var buffer bytes.Buffer
	err := t.Template.ExecuteTemplate(&buffer, t.TemplateName, map[string]interface{}{
		"context":   t.Context,
		"extension": t.Extension,
		"package":   t.Package,
		"path":      t.Path,
	})
	if err != nil {
		log.Error.Panicf("Didn't manage to execute a template, error: '%s'.", err)
	}

	// Print debugging information.
	path := filepath.Join(t.Path, t.Package+t.Extension)
	log.Info.Printf("Saving generated '%s' file to '%s'.", t.Package, path)

	// Write result to the file.
	err = ioutil.WriteFile(path, buffer.Bytes(), 0644)
	if err != nil {
		log.Error.Panicf("Failed to save generated file, error: '%s'.", err)
	}
}
