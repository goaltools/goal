package handlers

import (
	"testing"
)

func TestParentPackage(t *testing.T) {
	p := parent{}
	s := p.Package()
	if s != "" {
		// E.g. if we are using it for generation of:
		//	uniquePkgName.Application.Index.
		// In case the Application is local (i.e. its import is empty) we need:
		//	Application.Index.
		// I.e. the method must return empty string.
		t.Errorf("Packages with empty imports must have no names.")
	}
	p = parent{
		ID:     1,
		Import: "net/http",
		Name:   "Request",
	}
	s = p.Package(".XXX")
	if s != "c1.XXX" {
		t.Errorf(`Incorrect package name. Expected "c1.XXX", got "%s".`, s)
	}
}
