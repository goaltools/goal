package sample

import (
	"fmt"

	l "github.com/anonx/sunplate/log"
)

// Test is a type.
type Test struct {
	Name string `tag:"name"`
}

// Hello is a method.
func (t Test) Hello() string {
	l.Trace.Println("Greeting returned...")
	return fmt.Sprintf("Hello, %s!", t.Name)
}
