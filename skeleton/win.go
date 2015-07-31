// Graceful restarts and shutdowns are not supported
// on windows, so using usual http.Serve instead.
//
// +build windows

package main

import (
	"net/http"
)

// serve is a wrapper on standard http.Serve method.
func serve(s *http.Server) error {
	return http.Serve(s)
}
