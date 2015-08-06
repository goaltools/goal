// Package main is used for testing of generated 'views' listing.
// There is no way to include a new import dynamically, thus
// we are running this test from generate_test.go
// as a new command using exec package.
package main

import (
	"reflect"

	"github.com/anonx/sunplate/internal/programs/generation/testdata/assets/views"
	"github.com/anonx/sunplate/log"
)

func main() {
	if !reflect.DeepEqual(views.List, expectedOutput) {
		log.Error.Fatalf("Incorrect list. Expected %v, got %v.", expectedOutput, views.List)
	}

	if views.Paths.Subdir.SubSubdir.Test1TPL != expectedOutput[0] ||
		views.Paths.Subdir.Test1TPL != expectedOutput[1] ||
		views.Paths.Test1TPL != expectedOutput[2] ||
		views.Paths.Test2TPL != expectedOutput[3] {

		log.Error.Fatal("Values of Paths are not correct.")
	}
}

var expectedOutput = []string{
	"Subdir/SubSubdir/Test1.tpl",
	"Subdir/Test1.tpl",
	"Test1.tpl",
	"Test2.tpl",
	"z.tpl",
}
