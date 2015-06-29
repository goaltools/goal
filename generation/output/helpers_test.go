package output

import (
	"testing"
)

func TestExport(t *testing.T) {
	if r := export(""); r != "" {
		t.Errorf(`Empty input, the same string was expected as a result. Got "%s".`, r)
	}
	if r := export("a"); r != "A" {
		t.Errorf(`Incorrect result. Expected "A", got "%s".`, r)
	}
	if r := export("int"); r != "Int" {
		t.Errorf(`Incorrect result. Expected "Int", got "%s".`, r)
	}
}
