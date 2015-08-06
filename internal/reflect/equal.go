package reflect

// This file contains a set of debugging functions that are used
// by tests of reflect and some other packages.

import (
	"fmt"
	r "reflect"
)

// AssertEqualType gets two *Type arguments and returns nil in case they
// are equal to each other and an error otherwise.
func AssertEqualType(t1, t2 *Type) error {
	if t1 == nil || t2 == nil {
		if t1 != t2 {
			return fmt.Errorf("one of the types is equal to nil while another is not: %#v != %#v", t1, t2)
		}
		return nil
	}
	if t1.String() != t2.String() {
		return fmt.Errorf("types are not equal: %s != %s", t1, t2)
	}
	return nil
}

// AssertEqualArg gets two *Arg parameters and returns nil in case they
// are equal and an error otherwise.
func AssertEqualArg(a1, a2 *Arg) error {
	if a1 == nil || a2 == nil {
		if a1 != a2 {
			return fmt.Errorf("one of the arguments is nil while another is not: %#v != %#v", a1, a2)
		}
		return nil
	}
	if a1.Name != a2.Name {
		return fmt.Errorf("arguments have different names: %s != %s", a1.Name, a2.Name)
	}
	if a1.Tag != a2.Tag {
		return fmt.Errorf("arguments have different tags: %s != %s", a1.Tag, a2.Tag)
	}
	return AssertEqualType(a1.Type, a2.Type)
}

// AssertEqualArgs gets two lists of arguments and returns nil if they are
// equal to each other and an error otherwise.
func AssertEqualArgs(as1, as2 Args) error {
	if len(as1) != len(as2) {
		return fmt.Errorf(
			"Argument slices %#v and %#v have different length: %d != %d.",
			as1, as2, len(as1), len(as2),
		)
	}
	for i := range as1 {
		if err := AssertEqualArg(&as1[i], &as2[i]); err != nil {
			return err
		}
	}
	return nil
}

// AssertEqualFunc returns nil if received functions are equal to each other
// and an error otherwise.
func AssertEqualFunc(f1, f2 *Func) error {
	if f1 == nil || f2 == nil {
		if f1 != f2 {
			return fmt.Errorf("one of the funcs is nil while another is not: %#v != %#v", f1, f2)
		}
		return nil
	}
	if f1.Name != f2.Name {
		return fmt.Errorf("functions have different names: %s != %s", f1.Name, f2.Name)
	}
	if f1.File != f2.File {
		return fmt.Errorf("functions are from different files: %s != %s", f1.File, f2.File)
	}
	if !r.DeepEqual(f1.Comments, f2.Comments) {
		return fmt.Errorf("Comments of funcs are not equal: %#v != %#v.", f1.Comments, f2.Comments)
	}
	if err := AssertEqualArg(f1.Recv, f2.Recv); err != nil {
		return err
	}
	if err := AssertEqualArgs(f1.Params, f2.Params); err != nil {
		return err
	}
	return AssertEqualArgs(f1.Results, f2.Results)
}

// AssertEqualFuncs returns nil if received function slices are equal to each other
// and an error otherwise.
func AssertEqualFuncs(fs1, fs2 Funcs) error {
	if len(fs1) != len(fs2) {
		return fmt.Errorf(
			"func slices %#v and %#v have different length: %d != %d",
			fs1, fs2, len(fs1), len(fs2),
		)
	}
	for i := range fs1 {
		if err := AssertEqualFunc(&fs1[i], &fs2[i]); err != nil {
			return err
		}
	}
	return nil
}

// AssertEqualStruct gets two *Struct arguments and makes sure they are equal.
// If they are a nil is returned. Otherwise, it will return an error.
func AssertEqualStruct(s1, s2 *Struct) error {
	if s1 == nil || s2 == nil {
		if s1 != s2 {
			return fmt.Errorf("one of the structs is nil while another is not: %#v != %#v", s1, s2)
		}
		return nil
	}
	if s1.Name != s2.Name {
		return fmt.Errorf("structures have different names: %s != %s", s1.Name, s2.Name)
	}
	if s1.File != s2.File {
		return fmt.Errorf("structures are from different files: %s != %s", s1.File, s2.File)
	}
	if !r.DeepEqual(s1.Comments, s2.Comments) {
		return fmt.Errorf("comments of structs are not equal: %#v != %#v", s1.Comments, s2.Comments)
	}
	return AssertEqualArgs(s1.Fields, s2.Fields)
}

// AssertEqualStructs checks whether two slices of Structs are equal.
// If they are nil is returned. Otherwise, it will return an error.
func AssertEqualStructs(ss1, ss2 Structs) error {
	if len(ss1) != len(ss2) {
		return fmt.Errorf(
			"struct slices %#v and %#v have different length: %d and %d",
			ss1, ss2, len(ss1), len(ss2),
		)
	}
	for i := range ss1 {
		if err := AssertEqualStruct(&ss1[i], &ss2[i]); err != nil {
			return err
		}
	}
	return nil
}

// AssertEqualMethods returns nil if received maps of methods are equal to each other
// and an error otherwise.
func AssertEqualMethods(ms1, ms2 Methods) error {
	if len(ms1) != len(ms2) {
		return fmt.Errorf(
			"methods maps %#v and %#v have different length: %d != %d",
			ms1, ms2, len(ms1), len(ms2),
		)
	}
	for k := range ms1 {
		if err := AssertEqualFuncs(ms1[k], ms2[k]); err != nil {
			return err
		}
	}
	return nil
}

// AssertEqualPkg makes sure two packages are equal to each other.
// If they are nil is returned. Otherwise, it returns an error.
func AssertEqualPkg(p1, p2 *Package) error {
	if p1 == nil || p2 == nil {
		if p1 != p2 {
			return fmt.Errorf("one of the packages is nil while another is not: %#v != %#v", p1, p2)
		}
		return nil
	}
	if p1.Name != p2.Name {
		return fmt.Errorf("packages have different names: %s != %s", p1.Name, p2.Name)
	}
	if !r.DeepEqual(p1.Imports, p2.Imports) {
		return fmt.Errorf("imports of packages are not equal: %#v != %#v", p1.Imports, p2.Imports)
	}
	if err := AssertEqualStructs(p1.Structs, p2.Structs); err != nil {
		return err
	}
	if err := AssertEqualFuncs(p1.Funcs, p2.Funcs); err != nil {
		return err
	}
	return AssertEqualMethods(p1.Methods, p2.Methods)
}
