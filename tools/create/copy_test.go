package create

import (
	"os"
	"testing"
)

func TestCopyFile_DoesNotExist(t *testing.T) {
	defer expectPanic("File does not exist, panic expected.")
	copyFile("fileThatDoesNotExist", "./xxx")
}

func TestCopyFile(t *testing.T) {
	src := "./testdata/skeleton/test.go"
	dst := "./testdata/test.go"
	copyFile(src, dst)

	sf, err := os.Open(src)
	if err != nil {
		t.Errorf(`Cannot open source file. Error: "%s".`, err)
		t.FailNow()
	}
	defer sf.Close()

	df, err := os.Open(dst)
	if err != nil {
		t.Errorf(`Cannot open destination file. Error: "%s".`, err)
		t.FailNow()
	}
	defer df.Close()

	srcInfo, _ := sf.Stat()
	dstInfo, _ := df.Stat()

	d := make([]byte, dstInfo.Size())
	_, err = df.Read(d)
	if err != nil {
		t.Errorf(`Cannot read from destination file. Error: "%s".`, err)
		t.FailNow()
	}

	exp := "package test"
	if r := string(d); r != (exp+"\n") && r != (exp+"\r\n") {
		t.Errorf("Incorrect content of the copied file. Expected `%s`, got `%s`.", exp, r)
	}

	expMod := srcInfo.Mode().String()
	rMod := dstInfo.Mode().String()
	if expMod != rMod {
		t.Errorf(`Incorrect mode of the created file. Expected "%s", got "%s".`, expMod, rMod)
	}

	// Remove the created file.
	os.Remove(dst)
}

func TestCopyModifiedFile_DoesNotExist(t *testing.T) {
	defer expectPanic("File does not exist, panic expected.")
	copyModifiedFile("fileThatDoesNotExist", "./xxx", [][][]byte{})
}

func TestCopyModifiedFile(t *testing.T) {
	src := "./testdata/skeleton/test.go"
	dst := "./testdata/test.go"
	copyModifiedFile(src, dst, [][][]byte{
		{
			[]byte("test"), []byte("somethingCool"),
		},
	})

	sf, err := os.Open(src)
	if err != nil {
		t.Errorf(`Cannot open source file. Error: "%s".`, err)
		t.FailNow()
	}
	defer sf.Close()

	df, err := os.Open(dst)
	if err != nil {
		t.Errorf(`Cannot open destination file. Error: "%s".`, err)
		t.FailNow()
	}
	defer df.Close()

	srcInfo, _ := sf.Stat()
	dstInfo, _ := df.Stat()

	d := make([]byte, dstInfo.Size())
	_, err = df.Read(d)
	if err != nil {
		t.Errorf(`Cannot read from destination file. Error: "%s".`, err)
		t.FailNow()
	}

	exp := "package somethingCool"
	if r := string(d); r != (exp+"\n") && r != (exp+"\r\n") {
		t.Errorf("Incorrect content of the copied file. Expected `%s`, got `%s`.", exp, r)
	}

	expMod := srcInfo.Mode().String()
	rMod := dstInfo.Mode().String()
	if expMod != rMod {
		t.Errorf(`Incorrect mode of the created file. Expected "%s", got "%s".`, expMod, rMod)
	}

	// Remove the created file.
	os.Remove(dst)
}
