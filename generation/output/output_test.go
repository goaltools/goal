// +build linux

package output

import (
	"os"
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
	defer expectPanic("when we are not able to read template file, a panic must occur")
	NewType("test", "./pathThatDoesNotExist")
}

func TestCreateDir_ExistingDirectory(t *testing.T) {
	typ := Type{}
	typ.CreateDir("./testdata")
	if typ.Path != "./testdata" {
		t.Error("Type.Path is expected to be initialized")
	}
}

func TestCreateDir_NoWritePrivileges(t *testing.T) {
	// Prepare a readonly directory.
	os.Chmod("./testdata/readonly", 0544)

	typ := Type{}
	defer expectPanic("we have no write privileges for './testdata/readonly', so panic expected")
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
	typ := NewType("test", "./testdata/incorrect_template.html")
	typ.Path = "./testdata"
	defer expectPanic("template has errors and thus an error expected")
	typ.Generate()
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
