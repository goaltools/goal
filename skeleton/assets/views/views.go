// Package views is generated automatically by Sunplate toolkit.
// Please, do not edit it manually.
package views

// Root is a directory where templates are located.
var Root = "./views"

// List is a list of files that were found at "./views"
// in a form of slice of strings.
var List []string

// Paths stores information about all files that
// were found at "./views".
var Paths tPaths

// tPaths represents a root directory with files.
type tPaths struct {
	//
	// Below are the assets of root directory.
	//

	// App is a "App" directory.
	App tPathApp
	// BaseHTML is a "Base.html" file.
	BaseHTML string
}

// tApp is a type that represents a directory.
type tPathApp struct {
	//
	// Below are the assets of this directory.
	//

	// GreetHTML is a "App/Greet.html" file.
	GreetHTML string
	// IndexHTML is a "App/Index.html" file.
	IndexHTML string
}

func init() {
	Paths.App.GreetHTML = "App/Greet.html"
	Paths.App.IndexHTML = "App/Index.html"
	Paths.BaseHTML = "Base.html"
	List = []string{ // Make file paths available in a form of slice of strings.
		Paths.App.GreetHTML,
		Paths.App.IndexHTML,
		Paths.BaseHTML,
	}
}
