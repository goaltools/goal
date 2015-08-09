package strconv

import (
	"net/url"
	"reflect"
	"testing"
)

/*
	Below are testing functions for booleans.
*/

func TestBool(t *testing.T) {
	k := "b"
	exp := []bool{
		true, false, false, true, false,
	}
	for i, v := range exp {
		if r := Bool(vs, k, i); r != v {
			t.Errorf(errMsg, v, r)
		}
	}
}

func TestBools(t *testing.T) {
	k := "b"
	exp := []bool{
		true, false, false, true, false,
	}
	rs := Bools(vs, k)
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

/*
	Below are testing functions for strings.
*/

func TestString(t *testing.T) {
	k := "s"
	exp := "z"
	if r := String(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestStrings(t *testing.T) {
	k := "s"
	rs := Strings(vs, k)
	if !reflect.DeepEqual(rs, vs[k]) {
		t.Errorf(errMsg, vs[k], rs)
	}
}

/*
	Below are testing functions for integers.
*/

func TestInt(t *testing.T) {
	k := "i"
	exp := 1
	if r := Int(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestInts(t *testing.T) {
	k := "i"
	rs := Ints(vs, k)
	exp := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestInt8(t *testing.T) {
	k := "i"
	exp := int8(1)
	if r := Int8(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestInt8s(t *testing.T) {
	k := "i"
	rs := Int8s(vs, k)
	exp := []int8{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestInt16(t *testing.T) {
	k := "i"
	exp := int16(1)
	if r := Int16(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestInt16s(t *testing.T) {
	k := "i"
	rs := Int16s(vs, k)
	exp := []int16{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestInt32(t *testing.T) {
	k := "i"
	exp := int32(1)
	if r := Int32(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestInt32s(t *testing.T) {
	k := "i"
	rs := Int32s(vs, k)
	exp := []int32{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestInt64(t *testing.T) {
	k := "i"
	exp := int64(1)
	if r := Int64(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestInt64s(t *testing.T) {
	k := "i"
	rs := Int64s(vs, k)
	exp := []int64{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

/*
	Below are testing functions for unsigned integers.
*/

func TestUint(t *testing.T) {
	k := "i"
	exp := uint(1)
	if r := Uint(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestUints(t *testing.T) {
	k := "i"
	rs := Uints(vs, k)
	exp := []uint{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestUint8(t *testing.T) {
	k := "i"
	exp := uint8(1)
	if r := Uint8(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestUint8s(t *testing.T) {
	k := "i"
	rs := Uint8s(vs, k)
	exp := []uint8{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestUint16(t *testing.T) {
	k := "i"
	exp := uint16(1)
	if r := Uint16(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestUint16s(t *testing.T) {
	k := "i"
	rs := Uint16s(vs, k)
	exp := []uint16{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestUint32(t *testing.T) {
	k := "i"
	exp := uint32(1)
	if r := Uint32(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestUint32s(t *testing.T) {
	k := "i"
	rs := Uint32s(vs, k)
	exp := []uint32{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestUint64(t *testing.T) {
	k := "i"
	exp := uint64(1)
	if r := Uint64(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestUint64s(t *testing.T) {
	k := "i"
	rs := Uint64s(vs, k)
	exp := []uint64{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

/*
	Bellow are testing functions for floating-point numbers.
*/

func TestFloat32(t *testing.T) {
	k := "f"
	exp := float32(2.2)
	if r := Float32(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestFloat32s(t *testing.T) {
	k := "f"
	rs := Float32s(vs, k)
	exp := []float32{1.1, 2.2}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

func TestFloat64(t *testing.T) {
	k := "f"
	exp := float64(2.2)
	if r := Float64(vs, k); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

func TestFloat64s(t *testing.T) {
	k := "f"
	rs := Float64s(vs, k)
	exp := []float64{1.1, 2.2}
	if !reflect.DeepEqual(rs, exp) {
		t.Errorf(errMsg, exp, rs)
	}
}

/*
	Below are tests for additional helper functions.
*/

func TestGet(t *testing.T) {
	k := "s"
	is := []int{2, -1, 10}

	exp := "z"
	if r := get(vs, k, is); r != exp {
		t.Errorf(errMsg, exp, r)
	}

	exp = ""
	if r := get(vs, "", is); r != exp {
		t.Errorf(errMsg, exp, r)
	}

	exp = ""
	if r := get(vs, k, []int{10}); r != exp {
		t.Errorf(errMsg, exp, r)
	}

	exp = ""
	if r := get(vs, "keyThatDoesNotExist", []int{}); r != exp {
		t.Errorf(errMsg, exp, r)
	}
}

var vs = url.Values{
	"b": {"t", "false", "0", "yes", "f"},
	"f": {"1.1", "2.2"},
	"i": {"5", "4", "3", "2", "1"},
	"s": {"x", "y", "z"},
}
var errMsg = `Incorrect result. Expected "%v", got "%v".`
