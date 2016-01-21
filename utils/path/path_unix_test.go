//+build !windows

package path

import (
	"os"
	"path/filepath"
	"testing"
)

// TestImportToAbsolute_GetwdError simulates a failure of Getwd
// command. It relies on the behaviour in *nix where it is
// possible to move to a directory then delete it, causing Getwd's error.
func TestImportToAbsolute_GetwdError(t *testing.T) {
	pushd(t)

	p, err := ImportToAbsolute("../")
	if p != "" || err == nil {
		t.Errorf(`Getwd failed, error expected. Got "%s", "%v".`, p, err)
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
