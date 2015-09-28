package path

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathAbsolute_GetwdError(t *testing.T) {
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

	// Testing Absolute().
	res, err := New("./something").Absolute()
	if res != nil || err == nil {
		t.Errorf(`Getwd failed. Error expected; got "%v", "%v".`, res, err)
	}

	// Repairing old state for running tests from current dir.
	err = os.Chdir("../")
	if err != nil {
		t.Error(err)
	}
}

func TestPathAbsolute(t *testing.T) {
	currDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Cannot determine current directory.")
	}
	for _, v := range []struct {
		p *Path

		exp string
		err bool
	}{
		{
			p: New(""),

			exp: currDir,
		},
		{
			p: New("/home/user/stuff"),

			exp: "/home/user/stuff",
		},
		{
			p: New("../"),

			exp: filepath.Join(currDir, "../"),
		},
	} {
		if s, err := v.p.Absolute(); s.String() != v.exp ||
			(v.err && err == nil) || (!v.err && err != nil) {

			t.Errorf(
				`Expected "%s", err == nil -> "%v". Got "%s", "%v".`,
				v.exp, v.err, s.String(), err,
			)
		}
	}
}
