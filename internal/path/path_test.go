package path

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathAbsolute_GetwdError(t *testing.T) {
	pushd(t)

	p, err := New("./something").Absolute()
	if p != nil || err == nil {
		t.Errorf(`Getwd failed. Error expected; got "%v", "%v".`, p, err)
	}

	popd(t)
}

func TestPathAbsolute(t *testing.T) {
	currDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Cannot determine current directory.")
	}
	for _, v := range []struct {
		p   *Path
		exp string
	}{
		{
			p:   New(""),
			exp: currDir,
		},
		{
			p:   New("/home/user/stuff"),
			exp: filepath.Clean("/home/user/stuff"),
		},
		{
			p:   New("../"),
			exp: filepath.Join(currDir, "../"),
		},
	} {
		if p, err := v.p.Absolute(); p.String() != v.exp || err != nil {
			t.Errorf(
				`Expected "%s", "nil". Got "%s", "%v".`,
				v.exp, p.String(), err,
			)
		}
	}
}

func TestPathImport(t *testing.T) {
	for _, v := range []struct {
		p   *Path
		exp string
	}{
		{
			p:   New(""),
			exp: "github.com/colegion/goal/internal/path",
		},
		{
			p:   New("github.com/revel/revel"),
			exp: "github.com/revel/revel",
		},
		{
			p:   New("../../"),
			exp: "github.com/colegion/goal",
		},
	} {

		if p, err := v.p.Import(); p.String() != v.exp || err != nil {
			t.Errorf(
				`Expected "%s", "nil". Got "%s", "%v".`,
				v.exp, p.String(), err,
			)
		}
	}
}

func TestPathImport_OutsideGOPATH(t *testing.T) {
	p, err := New("../../../../../../../").Import()
	if p != nil || err == nil {
		t.Errorf(`Imports outside of "$GOPATH/src" must not be allowed. Got "%v", "%v".`, p, err)
	}
}

func TestPathImport_GetwdError(t *testing.T) {
	pushd(t)

	p, err := New("../").Import()
	if p != nil || err == nil {
		t.Errorf(`Getwd failed. Error expected; got "%v", "%v".`, p, err)
	}

	popd(t)
}

func pushd(t *testing.T) {
	// Preparing a state for os.Getwd to fail.
	dir := "./someDirectoryThatDoesNotExistYet"
	err := os.Mkdir(dir, 0755)
	if err != nil {
		t.Error(err)
	}
	err = os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}
	err = os.Remove(filepath.Join("../", dir))
	if err != nil {
		t.Error(err)
	}
}

func popd(t *testing.T) {
	// Repairing old state for running tests from current dir.
	err := os.Chdir("../")
	if err != nil {
		t.Error(err)
	}
}
