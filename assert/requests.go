package assert

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/anonx/sunplate/log"
)

// URL sets a url that will be used by other methods.
// It returns *Type so it can be used as t := New().URL("x")
func (t *Type) URL(s string) *Type {
	t.url = s
	return t
}

// Get creates a GET request to the given URN and saves it to
// the Request field of Type along with allocation
// of initialization of ResponseWriter.
func (t *Type) Get(urn string) *Type {
	req, err := http.NewRequest("GET", t.url+urn, nil)
	if err != nil {
		log.Error.Panic(err)
	}
	return t.NewRequest(req)
}

// Post creates a POST request to the given URN and saves it to
// the Request field of Type along with allocation
// of initialization of ResponseWriter.
func (t *Type) Post(urn string, contentType string, reader io.Reader) *Type {
	req, err := http.NewRequest("POST", t.url+urn, reader)
	if err != nil {
		log.Error.Panic(err)
	}
	req.Header.Set("Content-Type", contentType)
	return t.NewRequest(req)
}

// PostForm creates a POST request to the given URN that imitates a submission
// of "urlencoded" form and saves it to
// the Request field of Type along with allocation
// of initialization of ResponseWriter.
func (t *Type) PostForm(urn string, params url.Values) *Type {
	return t.Post(urn, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
}

// PostFile creates a POST request to the given URN that imitates a submission
// of "multipart" form and saves it to
// the Request field of Type along with allocation
// of initialization of ResponseWriter.
func (t *Type) PostFile(urn string, params url.Values, filePaths url.Values) *Type {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Prepare files to be added to request.
	for k, v := range filePaths {
		for i := range v {
			addFile(writer, k, v[i])
		}
	}

	// Make sure we have added non-file fields of form.
	for k, v := range params {
		for i := range v {
			err := writer.WriteField(k, v[i])
			if err != nil {
				log.Error.Panicf(`Cannot add value "%s" to field "%s", error: %v.`, v[i], k, err)
			}
		}
	}

	// Close the multipart writer.
	err := writer.Close()
	if err != nil {
		log.Error.Panicf("Failed to close a writer, error: %v.", err)
	}

	return t.Post(urn, writer.FormDataContentType(), body)
}

// NewRequest sets Request field of Type to the value that was given
// to it and (re)initializes ResponseWriter field.
func (t *Type) NewRequest(req *http.Request) *Type {
	log.Trace.Printf(`%s "%s"`, req.Method, req.URL)

	t.Request = req
	t.ResponseWriter = httptest.NewRecorder()

	// Create a new Assertion for Body so it can
	// be easily be tested by user.
	t.Body = t.NewAssertion(t.ResponseWriter.Body)
	return t
}

// addFile creates a part that contains content-disposition header,
// content-type, file name, and content of the specified file.
// Read RFC-1867 [7] (https://www.ietf.org/rfc/rfc1867.txt) for more info
// about uploading files.
// addFile panics in case of some error.
func addFile(writer *multipart.Writer, fieldname, filename string) {
	// Symbols that are expected to be escaped in file and field names.
	escaper := strings.NewReplacer("\\", "\\\\", "\"", "\\\"")

	// Try to open requested file.
	// Make sure it exists.
	file, err := os.Open(filename)
	if err != nil {
		log.Error.Panicf(`Cannot open "%s", error: %v.`, filename, err)
	}
	defer file.Close()

	// Create a new form-data header with the provided field name and file name.
	// Determine Content-Type of the file by its extension.
	h := textproto.MIMEHeader{}
	h.Set("Content-Disposition", fmt.Sprintf(
		`form-data; name="%s"; filename="%s"`,
		escaper.Replace(fieldname),
		escaper.Replace(filepath.Base(filename)),
	))
	h.Set("Content-Type", "application/octet-stream")
	if ct := mime.TypeByExtension(filepath.Ext(filename)); ct != "" {
		h.Set("Content-Type", ct)
	}
	part, err := writer.CreatePart(h)
	if err != nil {
		log.Error.Panicf("Cannot create a part, error: %v.", err)
	}

	// Copy content of the file we have just opened without reading
	// the whole file into memory.
	_, err = io.Copy(part, file)
	if err != nil {
		log.Error.Panicf("Cannot copy file, error: %v.", err)
	}
}
