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
			Name: "arg2",
			Type: &Type{
				Name: "type2",
			},
		},
		{
			Name: "arg1",
			Type: &Type{
				Name: "type1",
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
	if err := AssertEqualArgs(Args{a1[0]}, a2); err == nil {
		t.Errorf("Arguments are different. Error expected, got nil.")
	}
	if err := AssertEqualArgs(a1, a1); err != nil {
		t.Errorf("Nil expected, argument slices are equal. Got error: %s.", err)
	}
}

func TestAssertEqualFunc(t *testing.T) {
	if err := AssertEqualFunc(&Func{}, nil); err == nil {
		t.Error("One of the funcs is nil while another is not. Error expected, got nil.")
	}
	if err := AssertEqualFunc(nil, nil); err != nil {
		t.Errorf("Both functions are nil. Nil expected, got error: %s.", err)
	}
	if err := AssertEqualFunc(&Func{Name: "f1"}, &Func{Name: "f2"}); err == nil {
		t.Error("Functions have different names. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{Name: "f", File: "x1"}, &Func{Name: "f", File: "x2"}); err == nil {
		t.Error("Functions are from different files. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{Comments: Comments{}}, &Func{}); err == nil {
		t.Error("Functions have different comments. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{Recv: &Arg{}}, &Func{}); err == nil {
		t.Error("Functions have different receivers. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{Params: Args{{}}}, &Func{}); err == nil {
		t.Error("Functions have different parameters. Error expected, got nil.")
	}
	if err := AssertEqualFunc(&Func{}, &Func{}); err != nil {
		t.Errorf("Functions are identical. Nil expected, got error: %s.", err)
	}
}

func TestAssertEqualFuncs(t *testing.T) {
	fs1 := Funcs{
		{
			Name: "f2",
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
		t.Error("Functions have different length. Error expected, got nil.")
	}
	if err := AssertEqualFuncs(Funcs{fs1[0]}, fs2); err == nil {
		t.Error("Functions are different. Error expected, got nil.")
	}
	if err := AssertEqualFuncs(fs1, fs1); err != nil {
		t.Errorf("Function lists are identical. Nil expected, got %s.", err)
	}
}

func TestAssertEqualStruct(t *testing.T) {
	if err := AssertEqualStruct(&Struct{}, nil); err == nil {
		t.Error("One of the structures is nil while another is not. Error expected, got nil.")
	}
	if err := AssertEqualStruct(nil, nil); err != nil {
		t.Errorf("Both structures are nil. Nil expected, got error: %s.", err)
	}
	if err := AssertEqualStruct(&Struct{Name: "s1"}, &Struct{Name: "s2"}); err == nil {
		t.Error("Structs have different names. Error expected, got nil.")
	}
	if err := AssertEqualStruct(&Struct{Name: "s", File: "s1"}, &Struct{Name: "s", File: "s2"}); err == nil {
		t.Error("Structs are from different files. Error expected, got nil.")
	}
	if err := AssertEqualStruct(&Struct{Comments: Comments{}}, &Struct{}); err == nil {
		t.Error("Structs have different comments. Error expected, got nil.")
	}
	if err := AssertEqualStruct(&Struct{Name: "f"}, &Struct{Name: "f"}); err != nil {
		t.Errorf("Structs are equal to each other. Nil expected, got error: %s.", err)
	}
}

func TestAssertEqualStructs(t *testing.T) {
	ss1 := Structs{
		{
			Name: "s2",
			File: "f2",
		},
		{
			Name: "s1",
			File: "f1",
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
	if err := AssertEqualStructs(Structs{ss1[0]}, ss2); err == nil {
		t.Errorf("Lists are different. Error expected, got nil.")
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
	if err := AssertEqualMethods(Methods{"Controller": ms1["Controller"]}, ms2); err == nil {
		t.Errorf("Methods are different. Error expected, got nil.")
	}
	if err := AssertEqualMethods(ms1, ms1); err != nil {
		t.Errorf("Methods are identical. Nil expected, got error: %s.", err)
	}
}

func TestAssertEqualPkg(t *testing.T) {
	if err := AssertEqualPkg(&Package{}, nil); err == nil {
		t.Error("One of the packages is nil while another is not. Error expected, got nil.")
	}
	if err := AssertEqualPkg(nil, nil); err != nil {
		t.Errorf("Both packages are nil. Nil expected, got error: %s.", err)
	}
	if err := AssertEqualPkg(&Package{Name: "p1"}, &Package{Name: "p2"}); err == nil {
		t.Error("Packages have different names. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p", Imports: Imports{"x": map[string]string{}}}, &Package{Name: "p"}); err == nil {
		t.Error("Packages have imports. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p", Structs: Structs{{}}}, &Package{Name: "p"}); err == nil {
		t.Error("Packages have different structs. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p", Funcs: Funcs{{}}}, &Package{Name: "p"}); err == nil {
		t.Error("Packages have different functions. Error expected, got nil.")
	}
	if err := AssertEqualPkg(&Package{Name: "p"}, &Package{Name: "p"}); err != nil {
		t.Errorf("Packages are identical. Nil expected, got error: %s.", err)
	}
}
