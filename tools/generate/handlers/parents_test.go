package handlers

import (
	"testing"

	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/reflect"
)

func TestParentPackage(t *testing.T) {
	p := parent{
		Import: "github.com/colegion/goal/controllers",
	}
	s := p.Package("github.com/colegion/goal/controllers")
	if s != "" {
		t.Errorf("No accessor is needed for accessing a type from the same package. Got: %v.", s)
	}
	p = parent{
		ID:     1,
		Import: "net/http",
		Name:   "Request",
	}
	s = p.Package("github.com/colegion/goal/controllers", ".", "XXX")
	if s != "c1.XXX" {
		t.Errorf(`Incorrect package name. Expected "c1.XXX", got "%s".`, s)
	}
}

func TestParentAll(t *testing.T) {
	p := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"]
	p1 := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage"]
	p2 := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage"]
	c := p.list[0]
	res := c.Parents.All(ps)
	expCs := []*controller{
		p2.list[0],
		p1.list[0],
		p2.list[0],
		p.list[1],
	}
	assertDeepEqualControllerSlices(expCs, res.list)
	expInits := []reflect.Func{
		*p1.init,
		*p.init,
	}
	if err := reflect.AssertEqualFuncs(expInits, res.inits); err != nil {
		t.Error(err)
	}
}

func assertDeepEqualParents(p1, p2 parents) {
	if len(p1.list) != len(p2.list) {
		log.Error.Panicf("Different number of parents %d != %d.", len(p1.list), len(p2.list))
	}
	for i := range p1.list {
		if p1.list[i] != p2.list[i] {
			log.Error.Panicf("Different parents: %v != %v.", p1.list[i], p2.list[i])
		}
	}
	if p1.childImport != p2.childImport {
		log.Error.Panicf(`Different child imports: "%s" != "%s".`, p1.childImport, p2.childImport)
	}
}
