package tests

import (
	"testing"

	ts "github.com/anonx/sunplate/testsuite"
)

func TestAppIndex(t *testing.T) {
	s.Start()
	defer s.Close()

	ts := ts.New(s.URL)
	ts.Get("/")
	ts.AssertOK()
}
