package path

import (
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

func TestSunplateDir(t *testing.T) {
	p := filepath.Join(build.Default.GOPATH, "src", spImp)

	if v := SunplateDir(); v != p {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, p, v)
	}

	hp := filepath.Join(p, "generation", "handlers")
	if v := SunplateDir("generation", "handlers"); v != hp {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, hp, v)
	}
}

func TestWorkingDir(t *testing.T) {
	p := filepath.Join(build.Default.GOPATH, "src")
	os.Chdir(p)

	if v := WorkingDir(); v != p {
		t.Errorf(`Incorrectly detected working directory. Expected "%s", got "%s"`, p, v)
	}
}

func TestAbsoluteImport_AbsImportArgument(t *testing.T) {
	if v := AbsoluteImport(spImp); v != "github.com/anonx/sunplate" {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, spImp, v)
	}
}

func TestAbsoluteImport(t *testing.T) {
	os.Chdir(filepath.Join(build.Default.GOPATH, "src", spImp))

	d := "./generation"
	exp := filepath.Join(spImp, d)
	if v := AbsoluteImport(d); v != exp {
		t.Errorf(`Incorrect result. Expected "%s", got "%s".`, exp, v)
	}
}

var spImp = "github.com/anonx/sunplate"
