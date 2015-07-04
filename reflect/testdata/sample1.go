package sample

import (
	"fmt"

	l "github.com/anonx/sunplate/log"
	"github.com/anonx/sunplate/reflect"
)

// Test is a type.
type Test struct {
	Name string `tag:"name"`
}

// MapFunc is a test type.
type MapFunc map[string]reflect.Func

// Hello is a method.
func (t Test) Hello(names []string, args ...int) string {
	l.Trace.Println("Greeting returned...")
	return fmt.Sprintf("Hello, %s!", t.Name)
}
