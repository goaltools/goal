package create

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/path"
)

func TestStart_ExistingDir(t *testing.T) {
	defer expectPanic("Creation of a project in an existing directory should cause a panic.")
	start("create", command.Data{
		"create": "./testdata/existingDir",
	})
}

func TestStart(t *testing.T) {
	dst := "./testdata/project"
	start("create", command.Data{
		"create": dst,
	})

	rs1, fn1 := walkFunc(dst)
	filepath.Walk(dst, fn1)

	p := path.SunplateDir("skeleton")
	rs2, fn2 := walkFunc(p)
	filepath.Walk(p, fn2)

	if len(rs1.dirs) != len(rs2.dirs) ||
		len(rs1.files) != len(rs2.files) ||
		len(rs1.srcs) != len(rs2.srcs) {

		t.Error("Looks like not all go files, static files, and/or directories are copied.")
	}

	// Remove the directory we have created.
	os.RemoveAll(dst)
}

func TestWalkFunc_Error(t *testing.T) {
	_, fn := walkFunc("")
	TestError := errors.New("this is a test error")
	if err := fn("", nil, TestError); err != TestError {
		t.Errorf(`walkFunc expected to return TestError, returned "%s".`, err)
	}
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
