package sample

import (
	"fmt"

	l "github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/internal/reflect"
)

// Test is a type.
type Test struct {
	Name string `tag:"name"`
}

// MapFunc is a test type.
type MapFunc map[string]reflect.Func

// Hello is a method.
func (t Test) Hello(names []string, x interface{}, args ...int) string {
	l.Trace.Println("Greeting returned...")
	return fmt.Sprintf("Hello, %s!", t.Name)
}
