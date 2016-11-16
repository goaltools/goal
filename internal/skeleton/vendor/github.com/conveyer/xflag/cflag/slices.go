package cflag

import (
	"flag"
	"time"

	"github.com/conveyer/xflag/cflag/types"
)

// Strings is an equivalent of flag.String but for []string value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a string
// slice variable that stores the value of the flag.
func Strings(name string, value []string, usage string) *[]string {
	p := &types.Strings{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Ints is an equivalent of flag.Int but for []int value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of an int
// slice variable that stores the value of the flag.
func Ints(name string, value []int, usage string) *[]int {
	p := &types.Ints{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Int64s is an equivalent of flag.Int64 but for []int64 value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of an int64
// slice variable that stores the value of the flag.
func Int64s(name string, value []int64, usage string) *[]int64 {
	p := &types.Int64s{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Uints is an equivalent of flag.Uint but for []uint value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a uint
// slice variable that stores the value of the flag.
func Uints(name string, value []uint, usage string) *[]uint {
	p := &types.Uints{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Uint64s is an equivalent of flag.Uint64 but for []uint64 value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a uint64
// slice variable that stores the value of the flag.
func Uint64s(name string, value []uint64, usage string) *[]uint64 {
	p := &types.Uint64s{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Float64s is an equivalent of flag.Float64 but for []float64 value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a float64
// slice variable that stores the value of the flag.
func Float64s(name string, value []float64, usage string) *[]float64 {
	p := &types.Float64s{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Bools is an equivalent of flag.Bool but for []bool value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a bool
// slice variable that stores the value of the flag.
func Bools(name string, value []bool, usage string) *[]bool {
	p := &types.Bools{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}

// Durations is an equivalent of flag.Duration but for []time.Duration value.
// It defines a slice flag with the specified name, default value,
// and usage string. The returned value is the address of a time.Duration
// slice variable that stores the value of the flag.
func Durations(name string, value []time.Duration, usage string) *[]time.Duration {
	p := &types.Durations{Value: value}
	flag.Var(p, name, usage)
	return &p.Value
}
