package path

import (
	"go/build"
	"os"
	"path/filepath"
	"testing"
)

func TestAbsoluteToImport(t *testing.T) {
	for _, v := range []struct {
		abs string
		exp string
		exE bool // indicates whether an error expected.
	}{
		{
			abs: "",
			exE: true,
		},
		{
			abs: "./stuff",
			exE: true,
		},
		{
			abs: filepath.Join(build.Default.GOPATH, "src/github.com/revel/revel"),
			exp: "github.com/revel/revel",
		},
		{
			abs: filepath.Join(build.Default.GOPATH, "src/github.com/revel/revel/"),
			exp: "github.com/revel/revel",
		},
		{
			abs: filepath.Join(build.Default.GOPATH, "src"),
			exp: "",
		},
		{
			abs: filepath.Clean("/some/path/outside/GOPATH"),
			exE: true,
		},
	} {

		if p, err := AbsoluteToImport(v.abs); p != v.exp || v.exE && err == nil || !v.exE && err != nil {
			t.Errorf(
				`"%s": expected "%s", (err == nil -> "%v"). Got "%s", ("%v").`,
				v.abs, v.exp, v.exE, p, err,
			)
		}
	}
}

func TestImportToAbsolute(t *testing.T) {
	for _, v := range []struct {
		imp string
		exp string
		exE bool // indicates whether an error expected.
	}{
		{
			imp: "",
			exp: filepath.Join(build.Default.GOPATH, "src"),
		},
		{
			imp: value(os.Getwd()).(string), // use already valid abs path as input.
			exp: value(os.Getwd()).(string),
		},
		{
			imp: "github.com/colegion/goal",
			exp: filepath.Join(build.Default.GOPATH, "src", "github.com/colegion/goal"),
		},
		{
			imp: "..",
			exp: filepath.Join(value(os.Getwd()).(string), "../"),
		},
		{
			imp: "../",
			exp: filepath.Join(value(os.Getwd()).(string), "../"),
		},
		{
			imp: "../../../../../../",
			exE: true,
		},
		{
			imp: ".",
			exp: value(os.Getwd()).(string),
		},
		{
			imp: "./",
			exp: value(os.Getwd()).(string),
		},
	} {

		if p, err := ImportToAbsolute(v.imp); p != v.exp || v.exE && err == nil || !v.exE && err != nil {
			t.Errorf(
				`"%s": expected "%s", (err == nil -> "%v"). Got "%s", ("%v").`,
				v.imp, v.exp, v.exE, p, err,
			)
		}
	}
}

func TestImportToAbsolute_GetwdError(t *testing.T) {
	pushd(t)

	p, err := ImportToAbsolute("../")
	if p != "" || err == nil {
		t.Errorf(`Getwd failed, error expected. Got "%s", "%v".`, p, err)
	}

	popd(t)
}

func TestCleanImport(t *testing.T) {
	for _, v := range []struct {
		imp string
		exp string
		exE bool // indicates whether an error expected.
	}{
		{
			imp: "",
			exp: "",
		},
		{
			imp: "github.com/revel/revel",
			exp: "github.com/revel/revel",
		},
		{
			imp: "github.com/revel/revel/",
			exp: "github.com/revel/revel",
		},
		{
			imp: "./",
			exp: "github.com/colegion/goal/utils/path",
		},
	} {

		if p, err := CleanImport(v.imp); p != v.exp || v.exE && err == nil || !v.exE && err != nil {
			t.Errorf(
				`"%s": expected "%s", (err == nil -> "%v"). Got "%s", ("%v").`,
				v.imp, v.exp, v.exE, p, err,
			)
		}
	}
}

func value(v interface{}, err error) interface{} {
	return v
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
