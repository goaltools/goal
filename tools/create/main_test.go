package create

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/colegion/goal/internal/command"
	"github.com/colegion/goal/internal/path"
)

func TestMain_ExistingDir(t *testing.T) {
	defer expectPanic("Creation of a project in an existing directory should cause a panic.")
	main(handlers, 0, command.Data{"./testdata/existingDir"})
}

func TestStart(t *testing.T) {
	dst := "./testdata/project"
	main(handlers, 0, command.Data{dst})

	rs1, fn1 := walkFunc(dst)
	filepath.Walk(dst, fn1)

	p, err := path.New("github.com/colegion/goal/internal/skeleton").Package()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	rs2, err := walk(p.String())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

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

var handlers = []command.Handler{Handler}
