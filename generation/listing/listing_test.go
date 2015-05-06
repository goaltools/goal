package listing

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	Start(map[string]string{})
}

func TestInitDefaults(t *testing.T) {
	params := map[string]string{}
	initDefaults(params)

	for k, v := range expectedPaths {
		if params[k] != v {
			t.Errorf("default parameter '%s' is not set, expected '%s' got '%s'", k, v, params[k])
		}
	}
}

func TestWalkFunc(t *testing.T) {
	TestError := errors.New("this is a test error")
	if err := walkFunc("", nil, TestError); err != TestError {
		t.Errorf("walkFunc expected to return TestError, returned '%s'", err)
	}
	walkFunc("/myfile", testFile{}, nil)
	if len(files) == 0 || files[0] != "/myfile" {
		t.Error("failed to add path to files list")
	}
	err := walkFunc("", testFile{dir: true}, nil)
	if err != nil {
		t.Errorf("error expected to be nil, '%s' received instead", err)
	}
}

var expectedPaths = map[string]string{
	"--path":    "./views",
	"--output":  "./assets/",
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
