package views

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	Start(map[string]string{
		"--input":  "../testdata/views",
		"--output": "../testdata/assets/views",
	})

	// This function is tested on the level of "generation" package.

	// Remove the directory we have created.
	os.RemoveAll("../testdata/assets")
}

func TestWalkFunc(t *testing.T) {
	l, fn := walkFunc("")
	TestError := errors.New("this is a test error")
	if err := fn("", nil, TestError); err != TestError {
		t.Errorf(`walkFunc expected to return TestError, returned "%s".`, err)
	}

	err := fn("views/привет/myfile.txt", testFile{}, nil)
	if err == nil {
		t.Errorf("Incorrect directory name. Error expected, got nil.")
	}

	err = fn("views/myfile.data.txt", testFile{}, nil)
	if err == nil {
		t.Errorf("Incorrect directory name. Error expected, got nil.")
	}

	err = fn("views/index.html", testFile{}, nil)
	if err != nil {
		t.Errorf("Correct name. Nil expected, got error: %v.", err)
	}

	err = fn("./", testFile{}, nil)
	if err != nil {
		t.Errorf("Correct name. Nil expected, got error: %v.", err)
	}

	l, fn = walkFunc("views")
	fn("views/app/myfile.txt", testFile{}, nil)
	if len(l) == 0 || len(l["app"]) != 1 || l["app"][0].Name() != "myfileTXT" {
		t.Errorf("Failed to add a new path to the listing: %#v.", l)
	}

	l, fn = walkFunc("./views")
	fn("views/app", testFile{dir: true}, nil)
	if len(l) != 1 || len(l[""]) != 1 || l[""][0].Name() != "app" {
		t.Errorf("Failed to add a dir path to the listing: %#v.", l)
	}

	fn("views/app/myfile.txt", testFile{}, nil)
	if len(l) != 2 || len(l["app"]) != 1 || l["app"][0].Name() != "myfileTXT" {
		t.Errorf("Failed to add a file path to the listing: %#v.", l)
	}
}

type testFile struct {
	dir bool
}

func (t testFile) Name() string       { return "test" }
func (t testFile) Size() int64        { return 0 }
func (t testFile) Mode() os.FileMode  { return os.ModeTemporary }
func (t testFile) ModTime() time.Time { return time.Now() }
func (t testFile) IsDir() bool        { return t.dir }
func (t testFile) Sys() interface{}   { return nil }
