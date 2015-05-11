// +build linux

package output

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestNewType(t *testing.T) {
	// We are making sure there are no any panics.
	typ := NewType("test", "./output.go")
	if typ.Package != "test" {
		t.Error("package name was not initialized")
	}
}

func TestNewType_IncorrectPath(t *testing.T) {
	defer expectPanic("When we are not able to read template file, a panic must occur.")
	NewType("test", "./pathThatDoesNotExist")
}

func TestCreateDir_ExistingDirectory(t *testing.T) {
	typ := Type{}
	typ.CreateDir("./testdata")
	if typ.Path != "./testdata" {
		t.Error("Type.Path is expected to be initialized.")
	}
}

func TestCreateDir_NoWritePrivileges(t *testing.T) {
	// Prepare a readonly directory.
	os.Chmod("./testdata/readonly", 0544)

	typ := Type{}
	defer expectPanic("We have no write privileges for './testdata/readonly', so panic expected.")
	typ.CreateDir("./testdata/readonly/something")
}

func TestCreateDir(t *testing.T) {
	typ := Type{}
	typ.CreateDir("./testdata/assets/something")

	// Make sure the directory exists.
	if _, err := os.Stat("./testdata/assets/something"); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("directory was not created, error: '%s'", err)
		}
	}

	// Remove the directories we have created.
	os.RemoveAll("./testdata/assets")
}

func TestGenerate_IncorrectTemplate(t *testing.T) {
	typ := NewType("test", "./testdata/incorrect.template")
	typ.Path = "./testdata"
	defer expectPanic("Template has errors and thus an error expected.")
	typ.Generate()
}

func TestGenerate_NoWritePrivileges(t *testing.T) {
	// Prepare a readonly directory.
	os.Chmod("./testdata/readonly", 0544)

	typ := NewType("test", "./output.go")
	typ.Path = "./testdata/readonly"
	defer expectPanic("We do not have write access to the directory, thus panic expected.")
	typ.Generate()
}

func TestGenerate(t *testing.T) {
	// Generate a new "test" package using "./testdata/test.template" template
	// and save it to "./testdata/result/test.go".
	typ := NewType("test", "./testdata/test.template")
	typ.CreateDir("./testdata/result/")
	typ.Extension = ".go"
	typ.Generate()

	// Read the file, make sure its content is valid.
	c, err := ioutil.ReadFile("./testdata/result/test.go")
	if err != nil {
		t.Errorf("Cannot read a generated file, error: '%s'.", err)
	}
	if res := strings.Trim(string(c), " \r\n\t"); res != "package test" {
		t.Errorf("Generated file expected to contain 'package test', instead it is '%s'.", res)
	}

	// Remove the directories we have created.
	os.RemoveAll("./testdata/result")
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
