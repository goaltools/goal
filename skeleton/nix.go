// When not on windows using graceful restarts and shutdowns.
//
// +build !windows

package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

// serve is a wrapper on gracehttp.Serve. As opposed to
// the standard http library, this one may be terminated
// and/or restarted without dropping any connections.
func serve(s *http.Server) error {
	return gracehttp.Serve(s)
}
