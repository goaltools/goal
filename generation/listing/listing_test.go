package listing

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	Start(map[string]string{})

	// Remove the directory we have created.
	os.RemoveAll(filepath.Join(expectedParams["--output"], "../"))
}

func TestWalkFunc(t *testing.T) {
	files, fn := walkFunc("")
	TestError := errors.New("this is a test error")
	if err := fn("", nil, TestError); err != TestError {
		t.Errorf(`walkFunc expected to return TestError, returned "%s".`, err)
	}
	fn("/myfile", testFile{}, nil)
	if len(files) == 0 || files["/myfile"] != "/myfile" {
		t.Error("Failed to add path to files list.")
	}
	err := fn("", testFile{dir: true}, nil)
	if err != nil {
		t.Errorf(`Error expected to be nil, "%s" received instead.`, err)
	}
}

var expectedParams = map[string]string{
	"--input":   "./views",
	"--output":  "./assets/views/",
	"--package": "views",
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
