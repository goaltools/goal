package main

import (
	"github.com/anonx/sunplate/skeleton/assets/views"

	"github.com/anonx/sunplate/controllers/results"
)

// The line below tells golang's generate command you want
// it to generate a list of views (views.Context) for you.
// Please, do not delete it unless you know what you are doing.
//
//go:generate sunplate generate views

func init() {
	// Define the templates that should be loaded.
	results.SetTemplatePaths(views.Root, views.List)
}
