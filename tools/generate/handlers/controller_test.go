package handlers

import (
	"path/filepath"
	r "reflect"
	"sort"
	"testing"

	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/reflect"
)

func TestControllerIgnoredArgs(t *testing.T) {
	c := controller{}
	a := ps["github.com/colegion/goal/tools/generate/handlers/testdata/controllers"].list[0].Actions[0]
	exp := ", _, _"
	if r := c.IgnoredArgs(&a); r != exp {
		t.Errorf(`Incorrect IgnoreArgs result. Expected "%s", got "%s".`, exp, r)
	}
}

func assertDeepEqualControllers(cs1, cs2 controllers) {
	if err := reflect.AssertEqualFunc(cs1.init, cs2.init); err != nil {
		log.Error.Panic(err)
	}
	if cs1.accessor != cs2.accessor {
		log.Error.Panicf(`Accessors are different: "%s" != "%s".`, cs1.accessor, cs2.accessor)
	}
	// Order of controller structs doesn't matter when comparing two controllers.
	// So, sort them.
	sort.Sort(ByName(cs1))
	sort.Sort(ByName(cs2))
	assertDeepEqualControllerSlices(cs1.list, cs2.list)
}

type ByName controllers

func (bn ByName) Len() int           { return len(bn.list) }
func (bn ByName) Swap(i, j int)      { bn.list[i], bn.list[j] = bn.list[j], bn.list[i] }
func (bn ByName) Less(i, j int) bool { return bn.list[i].Name < bn.list[j].Name }

func assertDeepEqualControllerSlices(cs1, cs2 []*controller) {
	if len(cs1) != len(cs2) {
		log.Error.Panicf("Controller slices are of different lengths: %d != %d.", len(cs1), len(cs2))
	}
	for i := range cs1 {
		assertDeepEqualController(cs1[i], cs2[i])
	}
}

func assertDeepEqualController(c1, c2 *controller) {
	if c1 == nil || c2 == nil {
		if c1 != c2 {
			log.Error.Panicf(
				"One of the controllers is equal to nil while another is not: %#v != %#v.", c1, c2,
			)
		}
		return
	}
	if filepath.Base(c1.File) != filepath.Base(c2.File) {
		log.Error.Panicf("Controllers are from different files: %s != %s.", c1.File, c2.File)
	}
	if !r.DeepEqual(c1.Comments, c2.Comments) {
		log.Error.Panicf("Controllers have different comments: %#v != %#v.", c1.Comments, c2.Comments)
	}
	assertDeepEqualParents(c1.Parents, c2.Parents)
	log.Trace.Println("Actions...")
	if err := reflect.AssertEqualFuncs(c1.Actions, c2.Actions); err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Println("Before...")
	if err := reflect.AssertEqualFunc(c1.Before, c2.Before); err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Println("After...")
	if err := reflect.AssertEqualFunc(c1.After, c2.After); err != nil {
		log.Error.Panic(err)
	}
	log.Trace.Println("Fields...")
	assertDeepEqualFields(c1.Fields, c2.Fields)
	log.Trace.Println("Routes...")
	assertDeepEqualRoutes(c1.Routes, c2.Routes)
}
