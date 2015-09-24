package create

import (
	"os"
	"testing"

	"github.com/colegion/goal/log"
)

func TestCopyFile(t *testing.T) {
	src := "./testdata/skeleton/test.go"
	dst := "./testdata/test.go"
	copyFile(src, dst)

	sf, err := os.Open(src)
	log.AssertNil(err)
	defer sf.Close()

	df, err := os.Open(dst)
	log.AssertNil(err)
	defer df.Close()

	srcInfo, _ := sf.Stat()
	dstInfo, _ := df.Stat()

	d := make([]byte, dstInfo.Size())
	_, err = df.Read(d)
	log.AssertNil(err)

	exp := "package test\n"
	if r := string(d); r != exp {
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

func TestCopyModifiedFile(t *testing.T) {
	src := "./testdata/skeleton/test.go"
	dst := "./testdata/test.go"
	copyModifiedFile(src, dst, map[string]string{
		"test": "somethingCool",
	})

	sf, err := os.Open(src)
	log.AssertNil(err)
	defer sf.Close()

	df, err := os.Open(dst)
	log.AssertNil(err)
	defer df.Close()

	srcInfo, _ := sf.Stat()
	dstInfo, _ := df.Stat()

	d := make([]byte, dstInfo.Size())
	_, err = df.Read(d)
	log.AssertNil(err)

	exp := "package somethingCool\n"
	if r := string(d); r != exp {
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
