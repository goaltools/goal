package handlers

import (
	"github.com/colegion/goal/internal/log"
)

func assertDeepEqualFields(f1, f2 []field) {
	if len(f1) != len(f2) {
		log.Error.Panicf("Different number of fields: %d != %d.", len(f1), len(f2))
	}
	for i := range f1 {
		if f1[i].Name != f2[i].Name || f2[i].Type != f2[i].Type {
			log.Error.Panicf("Different fields: %v != %v.", f1[i], f2[i])
		}
	}
}
