package main

import (
	"github.com/anonx/sunplate/internal/skeleton/assets/views"

	"github.com/anonx/sunplate/controllers/templates"
)

// The line below tells golang's generate command you want
// it to generate a list of templates found in your ../views directory.
// Please, do not delete it unless you know what you are doing.
//
//go:generate sunplate generate views

func init() {
	// Define the templates that should be loaded.
	// Views directory's path should be relative to the root of your project.
	templates.SetTemplatePaths(views.Root, views.List)
}
