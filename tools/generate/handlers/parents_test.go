package handlers

import (
	"testing"

	"github.com/colegion/goal/internal/log"
)

func TestParentControllerCalled(t *testing.T) {
	p := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"]
	c := p.list[0]
	pcs := c.Parents.All(ps, "", newContext())
	// Result (App): SubSubPackage, X, SubPackage, SubSubPackage, Controller
	exp := []bool{
		true,
		false,
		false,
		false,
		false,
	}
	if len(pcs) != len(exp) {
		t.Fail()
	}
	for i := range exp {
		if res := pcs[i].Called(); res != exp[i] {
			t.Errorf(`[%v] Expected: "%v", got "%v".`, i, exp[i], res)
		}
	}
}

func TestParentControllerAllocate(t *testing.T) {
	p := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"]
	c := p.list[0]
	pcs := c.Parents.All(ps, "", newContext())
	// Result (App): SubSubPackage, X, SubPackage, SubSubPackage, Controller
	exp := []string{
		"c.Controller.SubSubPackage",
		"&c1.X",
		"&c2.Controller",
		"&c0.SubSubPackage",
		"&customPackageName.Controller",
	}
	if len(pcs) != len(exp) {
		t.Fail()
	}
	for i := range exp {
		if res := pcs[i].Allocate("c", "customPackageName"); res != exp[i] {
			t.Errorf(`[%v]Incorrect allocation. Expected: "%s", got "%s".`, i, exp[i], res)
		}
	}
}

func TestParentControllersReverse(t *testing.T) {
	p := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"]
	c := p.list[0]
	pcs := c.Parents.All(ps, "", newContext()).Reverse()
	// Result (App): Controller, SubSubPackage, SubPackage, X, SubSubPackage
	exp := []string{
		"Controller",
		"Controller.SubSubPackage",
		"Controller.Controller",
		"Controller.Controller.X",
		"Controller.Controller.SubSubPackage",
	}
	if len(pcs) != len(exp) {
		t.Fail()
	}
	for i := range exp {
		if res := pcs[i].Access(); res != exp[i] {
			t.Errorf(`Incorrect access. Expected: "%s", got "%s".`, exp[i], res)
		}
	}
}

func TestParentControllersImports(t *testing.T) {
	p := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"]
	c := p.list[0]
	pcs := c.Parents.All(ps, "", newContext())

	exp_c0 := `c0 "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage"`
	exp_c1 := `c1 "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/x"`
	exp_c2 := `c2 "github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage"`
	exp := exp_c0 + "\n" + exp_c1 + "\n" + exp_c2 + "\n"

	if imp := pcs.Imports(); imp != exp {
		t.Errorf("Incorrect imports. Expected: `%s`, got `%s`.", exp, imp)
	}
}

func TestParentAll(t *testing.T) {
	p := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"]
	p1 := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage"]
	p2 := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/subsubpackage"]
	p3 := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers/subpackage/x"]
	c := p.list[0]
	pcs := c.Parents.All(ps, "", newContext())
	//	App {
	//		*Controller {
	//			*SubPackage {
	//				*SubSubPackage
	//				*X
	//			}
	//			*SubSubPackage
	//		}
	//	}
	//	// Result (App): SubSubPackage, X, SubPackage, SubSubPackage, Controller
	exp := parentControllers{
		{ // subsubpackage.SubSubPackage that embeds nothing.
			Accessor:   "c0",
			Prefix:     "Controller.Controller.", // The second Controller is of type subpackage.Controller.
			instance:   "Controller.SubSubPackage",
			Controller: p2.list[0],
		},
		{ // x.X that embeds nothing.
			Accessor:   "c1",
			Prefix:     "Controller.Controller.", // The second Controller is of type subpackage.Controller.
			instance:   "",
			Controller: p3.list[0],
		},
		{ // subpackage.Controller that embeds nothing.
			Accessor:   "c2",
			Prefix:     "Controller.",
			instance:   "",
			Controller: p1.list[0],
		},
		{ // subsubpackage.SubSubPackage that embeds nothing.
			Accessor:   "c0", // The same accessor as has the "subsubpackage" above.
			Prefix:     "Controller.",
			instance:   "", // The second Controller is of type subpackage.Controller.
			Controller: p2.list[0],
		},
		{ // Controller that embeds nothing.
			Accessor:   "",
			Prefix:     "",
			instance:   "",
			Controller: p.list[1],
		},
	}
	assertDeepEqualParentControllers(exp, pcs)
}

func assertDeepEqualParentControllers(pcs1, pcs2 parentControllers) {
	if len(pcs1) != len(pcs2) {
		log.Error.Panicf("Different number of parent controllers: %d != %d.", len(pcs1), len(pcs2))
	}
	for i := range pcs1 {
		log.Info.Printf(`[%v] Comparing "%s" and "%s".`, i, pcs1[i].Controller.Name, pcs2[i].Controller.Name)
		assertDeepEqualController(pcs1[i].Controller, pcs2[i].Controller)
		if pcs1[i].Accessor != pcs2[i].Accessor {
			log.Error.Panicf(`Expected accessor: "%s". Got: "%s".`, pcs1[i].Accessor, pcs2[i].Accessor)
		}
		if pcs1[i].Prefix != pcs2[i].Prefix {
			log.Error.Panicf(`Expected prefix: "%s". Got: "%s".`, pcs1[i].Prefix, pcs2[i].Prefix)
		}
		if pcs1[i].instance != pcs2[i].instance {
			log.Error.Panicf(`Expected instance: "%s". Got: "%s".`, pcs1[i].instance, pcs2[i].instance)
		}
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
