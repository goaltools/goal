package run

import (
	"fmt"
	"testing"
)

func TestExecute_IncorrectTask(t *testing.T) {
	ts := []string{
		"gxgbhsjdjdduuhdhsh", // This is expected to cause a panic.
	}
	defer expectPanic(fmt.Sprintf(`Command "%s" does not exist. Panic expected.`, ts[0]))
	execute(ts)
}

func TestExecute(t *testing.T) {
	ts := []string{
		`echo "Hello, world!"`,
		`echo "This is test"`,
	}
	execute(ts)
}

func TestUserCommand(t *testing.T) {
	imp := "github.com/anonx/sunplate/skeleton"
	res := userCommand(imp)
	exp := "skeleton"
	if res != exp {
		t.Errorf("Incorrect user command. Expected `%s`, got `%s`.", exp, res)
	}
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
