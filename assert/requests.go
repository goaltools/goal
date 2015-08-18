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

// StartServer starts a test server to use with integration tests.
func (t *Type) StartServer(h http.Handler) *Type {
	t.server = httptest.NewServer(h)
	t.URL(t.server.URL)
	return t
}

// TryStartServer is shortcut of StartServer for standard routing package
// that has a Build() (http.Handler, error).
// It makes sure error is nil and only then starts a server.
func (t *Type) TryStartServer(h http.Handler, err error) *Type {
	t.NewAssertion(err).Nil()
	return t.StartServer(h)
}

// StopServer stops a started test server.
func (t *Type) StopServer() {
	t.server.Close()
}

// Get creates a GET request to the given URN and saves it to
// the Request field of Type along with allocation
// of initialization of ResponseWriter.
func (t *Type) Get(urn string) *Type {
	req, err := http.NewRequest("GET", t.url+urn, nil)
	t.NewAssertion(err).Nil()
	return t.NewRequest(req)
}

// Post creates a POST request to the given URN and saves it to
// the Request field of Type along with allocation
// of initialization of ResponseWriter.
func (t *Type) Post(urn string, contentType string, reader io.Reader) *Type {
	req, err := http.NewRequest("POST", t.url+urn, reader)
	t.NewAssertion(err).Nil()
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
			t.addFile(writer, k, v[i])
		}
	}

	// Make sure we have added non-file fields of form.
	for k, v := range params {
		for i := range v {
			err := writer.WriteField(k, v[i])
			t.NewAssertion(err).Nil()
		}
	}

	// Close the multipart writer.
	err := writer.Close()
	t.NewAssertion(err).Nil()

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

// Do sends a t.request request to the t.server.
func (t *Type) Do() *Type {
	res, err := t.client.Do(t.Request)
	t.NewAssertion(err).Nil()
	t.ResponseWriter = &httptest.ResponseRecorder{
		Body:      new(bytes.Buffer),
		Code:      res.StatusCode,
		HeaderMap: res.Header,
	}
	t.ResponseWriter.Body.ReadFrom(res.Body)
	t.Body = t.NewAssertion(t.ResponseWriter.Body)
	t.Response = res
	return t
}

// addFile creates a part that contains content-disposition header,
// content-type, file name, and content of the specified file.
// Read RFC-1867 [7] (https://www.ietf.org/rfc/rfc1867.txt) for more info
// about uploading files.
// addFile panics in case of some error.
func (t *Type) addFile(writer *multipart.Writer, fieldname, filename string) {
	// Symbols that are expected to be escaped in file and field names.
	escaper := strings.NewReplacer("\\", "\\\\", "\"", "\\\"")

	// Try to open requested file.
	// Make sure it exists.
	file, err := os.Open(filename)
	t.NewAssertion(err).Nil()
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
	t.NewAssertion(err).Nil()

	// Copy content of the file we have just opened without reading
	// the whole file into memory.
	_, err = io.Copy(part, file)
	t.NewAssertion(err).Nil()
}
