package listing

import (
	"reflect"
	"testing"
)

func TestPathName(t *testing.T) {
	inp := []path{
		{Value: []string{"path", "to", "some", "file.txt"}},
		{Value: []string{"file"}},
		{Value: []string{"app"}, Dir: true},
		{Value: []string{"Path", "To", "Dir"}, Dir: true},
	}
	exp := []string{"fileTXT", "file", "app", "Dir"}
	for i, p := range inp {
		if r := p.Name(); r != exp[i] {
			t.Errorf(`Incorrect result with path "%s". Expected "%s", got "%s".`, p, exp[i], r)
		}
	}
}

func TestPathDirStructName(t *testing.T) {
	inp := []path{
		{Value: []string{"file"}},
		{Value: []string{"app"}, Dir: true},
		{Value: []string{"Path", "To", "Dir"}, Dir: true},
	}
	exp := []string{"", "app", "PathToDir"}
	for i, p := range inp {
		if r := p.DirStructName(); r != exp[i] {
			t.Errorf(`Incorrect result with path "%s". Expected "%s", got "%s".`, p, exp[i], r)
		}
	}
}

func TestListingAddFile(t *testing.T) {
	f := "path/to/some/dir/file.txt"
	f1 := "path/to/some/file.txt"
	f2 := "root.txt"
	f3 := "root1.txt"
	f4 := "path/to/some/file1.txt"

	l := listing{}
	l.addFile(f1)
	l.addFile(f2)
	l.addFile(f3)
	l.addFile(f4)
	l.addFile(f)

	r, ok := l["path/to/some/dir"]
	if !ok {
		t.Errorf(`File "%s" was not added: %#v.`, f, l)
	}

	if len(r) != 1 || len(r[0].Value) != 5 || r[0].Name() != "fileTXT" {
		t.Errorf(`File "%s" ("%s") was added incorrectly: %#v.`, f, r[0].Name(), l)
	}

	r, ok = l[""]
	if !ok {
		t.Errorf(`File "%s" was not added: %#v.`, f3, l)
	}

	if len(r) != 2 || r[0].Name() != "rootTXT" || r[1].Name() != "root1TXT" {
		t.Errorf(`Files "%s" and "%s" were added incorrectly: %#v.`, f2, f3, l)
	}

	r, ok = l["path/to/some"]
	if !ok {
		t.Errorf(`File "%s" was not added: %#v.`, f4, l)
	}

	if len(r) != 2 || r[0].Name() != "fileTXT" || r[1].Name() != "file1TXT" {
		t.Errorf(`Files "%s" and "%s" were added incorrectly: %#v.`, f1, f4, l)
	}
}

func TestPathString(t *testing.T) {
	f := path{Value: []string{"path", "to", "something"}, Dir: true}
	exp := "path/to/something"
	if r := f.String(); r != exp {
		t.Errorf(`Incorrect path to the file. Expected "%s", got "%s".`, exp, r)
	}
}

func TestListingFilePaths(t *testing.T) {
	l := listing{}
	l.addDir(".")
	l.addFile("Base.html")
	l.addDir("Path/To/Dir")
	l.addFile("Path/To/Dir/Index.html")
	l.addDir("Path/To/Dir/Subdir")
	l.addFile("Path/To/Dir/Subdir/Index.html")

	exp := map[string]string{
		"BaseHTML":                     "Base.html",
		"Path.To.Dir.IndexHTML":        "Path/To/Dir/Index.html",
		"Path.To.Dir.Subdir.IndexHTML": "Path/To/Dir/Subdir/Index.html",
	}
	if r := l.FilePaths(); !reflect.DeepEqual(r, exp) {
		t.Errorf(`Incorrect values of paths. Expected "%s", got "%s".`, exp, r)
	}
}

func TestListingDirs(t *testing.T) {
	l := listing{}
	l.addDir(".")
	l.addFile("Base.html")
	l.addDir("Path")
	l.addDir("Path/To")
	l.addDir("Path/To/Dir")
	l.addFile("Path/To/Dir/Index.html")
	l.addFile("Path/To/Dir/Index2.html")
	l.addDir("Path/To/Dir/Subdir")
	l.addFile("Path/To/Dir/Subdir/Index.html")

	exp := map[string][]path{
		"Path": {
			{Value: []string{"Path", "To"}, Dir: true},
		},
		"PathTo": {
			{Value: []string{"Path", "To", "Dir"}, Dir: true},
		},
		"PathToDir": {
			{Value: []string{"Path", "To", "Dir", "Index.html"}, Dir: false},
			{Value: []string{"Path", "To", "Dir", "Index2.html"}, Dir: false},
			{Value: []string{"Path", "To", "Dir", "Subdir"}, Dir: true},
		},
		"PathToDirSubdir": {
			{Value: []string{"Path", "To", "Dir", "Subdir", "Index.html"}, Dir: false},
		},
	}
	if r := l.Dirs(); !reflect.DeepEqual(r, exp) {
		t.Errorf(`Incorrect values of dirs. Expected "%s", got "%s".`, exp, r)
	}
}

func TestListingRootAssets(t *testing.T) {
	l := listing{}
	l.addDir(".")
	l.addFile("Base.html")
	l.addDir("Path")
	l.addDir("Path/To")
	l.addDir("Path/To/Dir")
	l.addFile("Path/To/Dir/Index.html")
	l.addFile("Path/To/Dir/Index2.html")
	l.addDir("Path/To/Dir/Subdir")
	l.addFile("Path/To/Dir/Subdir/Index.html")

	exp := []path{
		{Value: []string{"Base.html"}, Dir: false},
		{Value: []string{"Path"}, Dir: true},
	}
	if r := l.RootAssets(); !reflect.DeepEqual(r, exp) {
		t.Errorf(`Incorrect values of root directory assets. Expected "%s", got "%s".`, exp, r)
	}
}
