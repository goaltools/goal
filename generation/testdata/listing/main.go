// Package main is used for testing of generated 'views' listing.
// There is no way to include a new import dynamically, thus
// we are running this test from generate_test.go
// as a new command using exec package.
package main

import (
	"../assets/views"

	"github.com/anonx/sunplate/log"
)

func main() {
	if l := len(views.Context); l != 2 {
		log.Error.Fatalf("Length of views.Context expected to be equal to 2, it is %d instead.", l)
	}

	//
	// Make sure templates are presented in the format we expect.
	//
	for k, v := range expectedValues {
		if views.Context[k] != v {
			log.Error.Fatalf("'%s' wasn't found in %#v.", k, views.Context)
		}
	}
}

var expectedValues = map[string]string{
	"testdata/views/test1.template": "testdata/views/test1.template",
	"testdata/views/test2.template": "testdata/views/test2.template",
}
