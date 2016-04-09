package x

import (
	"net/http"
	"net/url"
)

// X ...
type X struct {
}

// Before ...
func (c *X) Before() http.Handler {
	return nil
}

// Init ...
func Init(url.Values) {
}
