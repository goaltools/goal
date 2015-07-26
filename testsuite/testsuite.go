// Package testsuite is a library of functions that may be
// useful for handler's testing.
// It is highly inspired by Revel Framework's testing package.
//
// Some methods require URI while others URN as their input parameters.
// For reference, URI is URL + URN. E.g. "https://example.com/test"
// is a URI, "https://example.com" is URL, and "/test" is URN.
package testsuite

import (
	"io/ioutil"
	"net/http"

	"github.com/anonx/sunplate/log"
)

// Type represents a test suite.
type Type struct {
	Client       *http.Client
	Response     *http.Response
	ResponseBody []byte
	URL          string
}

// Request represents a client HTTP request to server.
// It is a wrapper on standard http.Request that includes
// testsuite information, too.
type Request struct {
	*http.Request
	ts *Type
}

// New allocates and returns a new test suite.
// Server's URL is expected as an input argument.
func New(url string) *Type {
	return &Type{
		URL: url,
	}
}

// Send issues a request and reads the response. If successfull, the caller
// may examine the Response and ResponseBody properties.
func (r *Request) Send() {
	var err error
	if r.ts.Response, err = r.ts.Client.Do(r.Request); err != nil {
		log.Error.Panic(err)
	}
	if r.ts.ResponseBody, err = ioutil.ReadAll(r.ts.Response.Body); err != nil {
		log.Error.Panic(err)
	}
}

// Request allocates and returns a new Request
// for the requested testsuite.
func (t *Type) Request(req *http.Request) *Request {
	return &Request{
		Request: req,
		ts:      t,
	}
}

// Get issues a GET request to the given URN of server's URL
// and stores the result in Response and ResponseBody.
func (t *Type) Get(urn string) {
	log.Trace.Printf(`GET "%s"...`, urn)
	t.GetCustom(t.URL + urn).Send()
	log.Trace.Println("\tDONE.")
}

// GetCustom returns a GET request to the given URI in
// a form of Request structure.
func (t *Type) GetCustom(uri string) *Request {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Error.Panic(err)
	}
	return t.Request(req)
}
