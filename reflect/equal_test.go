package reflect

import (
	"testing"
)

func TestAssertEqualType(t *testing.T) {
	t1 := &Type{
		Name:    "Int",
		Package: "types",
		Star:    true,
	}
	t2 := &Type{
		Name:    "String",
		Package: "types",
		Star:    true,
	}
	if err := AssertEqualType(t1, t1); err != nil {
		t.Errorf("Nil is expected when input parameters are identical, got error instead: %s.", err)
	}
	if err := AssertEqualType(t1, nil); err == nil {
		t.Errorf("Error is expected when one of the types is nil while another is not. Got nil.")
	}
	if err := AssertEqualType(t1, t2); err == nil {
		t.Errorf("Error expected as %#v != %#v. Got nil.", t1, t2)
	}
}

func TestAssertEqualArg(t *testing.T) {
	if err := AssertEqualArg(&Arg{}, nil); err == nil {
		t.Errorf("One of the arguments if nil, while another is not. Error expected, got nil.")
	}

	if err := AssertEqualArg(&Arg{Name: "arg1"}, &Arg{Name: "arg2"}); err == nil {
		t.Errorf("Arguments have different names. Error expected, got nil.")
	}

	if err := AssertEqualArg(&Arg{Name: "arg", Tag: "1"}, &Arg{Name: "arg", Tag: "2"}); err == nil {
		t.Errorf("Arguments have different tags. Error expected, got nil.")
	}

	a := &Arg{
		Name: "page",
		Type: &Type{
			Name: "int",
		},
	}
	if err := AssertEqualArg(a, a); err != nil {
		t.Errorf("Equal arguments, nil expected. Got %s.", err)
	}
}

func TestAssertEqualArgs(t *testing.T) {
	a1 := Args{
		{
			Name: "arg1",
			Type: &Type{
				Name: "type1",
			},
		},
		{
			Name: "arg2",
			Type: &Type{
				Name: "type2",
			},
		},
	}
	a2 := Args{
		{
			Name: "arg1",
			Type: &Type{
				Name: "type1",
			},
		},
	}
	if err := AssertEqualArgs(a1, a2); err == nil {
		t.Errorf("Argument slices have different length. Error expected, got nil.")
	}
	if err := AssertEqualArgs(a1, a1); err != nil {
		t.Errorf("Nil expected, argument slices are equal. Got error: %s.", err)
	}
}

func TestAssertEqualFunc(t *testing.T) {
	if err := AssertEqualFunc(&Func{}, nil); err == nil {
		t.Errorf("One of the funcs is nil while another is not. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{Name: "f1"}, &Func{Name: "f2"}); err == nil {
		t.Errorf("Functions have different names. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{Name: "f", File: "x1"}, &Func{Name: "f", File: "x2"}); err == nil {
		t.Errorf("Functions are from different files. Error expected, got nil.")
	}
}

func TestAssertEqualFuncs(t *testing.T) {
	fs1 := Funcs{
		{
			Name: "f1",
			File: "test.go",
		},
		{
			Name: "f1",
			File: "test.go",
		},
	}
	fs2 := Funcs{
		{
			Name: "f1",
			File: "test.go",
		},
	}
	if err := AssertEqualFuncs(fs1, fs2); err == nil {
		t.Errorf("Functions have different length. Error expected, got nil.")
	}
	if err := AssertEqualFuncs(fs1, fs1); err != nil {
		t.Errorf("Function lists are identical. Nil expected, got %s.", err)
	}
}

func TestAssertEqualStruct(t *testing.T) {
	if err := AssertEqualStruct(&Struct{}, nil); err == nil {
		t.Errorf("One of the structures is nil while another is not. Error expected, got nil.")
	}
	if err := AssertEqualStruct(&Struct{Name: "s1"}, &Struct{Name: "s2"}); err == nil {
		t.Errorf("Structs have different names. Error expected, got nil.")
	}
	if err := AssertEqualStruct(&Struct{Name: "s", File: "s1"}, &Struct{Name: "s", File: "s2"}); err == nil {
		t.Errorf("Structs are from different files. Error expected, got nil.")
	}
}

func TestAssertEqualStructs(t *testing.T) {
	ss1 := Structs{
		{
			Name: "s1",
			File: "f1",
		},
		{
			Name: "s2",
			File: "f2",
		},
	}
	ss2 := Structs{
		{
			Name: "s1",
			File: "f1",
		},
	}
	if err := AssertEqualStructs(ss1, ss2); err == nil {
		t.Errorf("Lists of structs have different lengths. Error expected, got nil.")
	}
	if err := AssertEqualStructs(ss1, ss1); err != nil {
		t.Errorf("Lists of structs are identical. Nil expected, got error: %s.", err)
	}
}

func TestAssertEqualMethods(t *testing.T) {
	ms1 := Methods{
		"App": Funcs{
			{
				Name: "f1",
				File: "app.go",
			},
		},
		"Controller": Funcs{
			{
				Name: "f1",
				File: "app.go",
			},
		},
	}
	ms2 := Methods{
		"App": Funcs{
			{
				Name: "f1",
				File: "app.go",
			},
		},
	}
	if err := AssertEqualMethods(ms1, ms2); err == nil {
		t.Errorf("Methods groups have different length. Error expected, got nil.")
	}
	if err := AssertEqualMethods(ms1, ms1); err != nil {
		t.Errorf("Methods are identical. Nil expected, got error: %s.", err)
	}
}

func TestAssertEqualPkg(t *testing.T) {
	if err := AssertEqualPkg(&Package{}, nil); err == nil {
		t.Errorf("One of the packages is nil while another is not. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p1"}, &Package{Name: "p2"}); err == nil {
		t.Errorf("Packages have different names. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p", Imports: Imports{"x": map[string]string{}}}, &Package{Name: "p"}); err == nil {
		t.Errorf("Packages have imports. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p"}, &Package{Name: "p"}); err != nil {
		t.Errorf("Packages are identical. Nil expected, got error: %s.", err)
	}
}
