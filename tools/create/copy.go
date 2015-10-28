package create

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/colegion/goal/utils/log"
)

// copyFile reads a source file and copies it to the destination.
// It doesn't check whether input source parameter is a regular file
// rather than a directory.
func copyFile(src, dst string) {
	// Open source file.
	sf, err := os.Open(src)
	log.AssertNil(err)
	defer sf.Close() // Make sure the file will be closed.

	// Get the meta info of the source file.
	info, err := sf.Stat()
	log.AssertNil(err)

	// Create a destination file.
	df, err := os.Create(dst)
	log.AssertNil(err)
	defer df.Close() // Make sure it will be closed at the end.

	// Copy the content of source to destination.
	_, err = io.Copy(df, sf)
	log.AssertNil(err)

	// Set the chmod of source to destination file.
	err = os.Chmod(dst, info.Mode())
	log.AssertNil(err)
}

// copyModifiedFile is similar to copyFile except it takes changes of type
// [][][]byte as the third argument.
// Example:
//	[][][]byte{
//		[][]byte{
//			[]byte("key"), []byte("value"),
//		}
//	}
// Keys are what should be replaced and their values are the replacements.
func copyModifiedFile(src, dst string, changes [][][]byte) {
	// Open source file.
	sf, err := os.Open(src)
	log.AssertNil(err)
	defer sf.Close() // Make sure the file will be closed.

	// Get the meta info of the source file.
	info, err := sf.Stat()
	log.AssertNil(err)

	// Read its content.
	d := make([]byte, info.Size())
	_, err = sf.Read(d)
	log.AssertNil(err)

	// Make the required changes.
	for i := 0; i < len(changes); i++ {
		d = bytes.Replace(d, changes[i][0], changes[i][1], -1)
	}

	// Write the content to the destination file.
	err = ioutil.WriteFile(dst, d, info.Mode())
	log.AssertNil(err)
}
