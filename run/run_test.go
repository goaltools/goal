package run

import (
	"fmt"
	"syscall"
	"testing"

	"github.com/anonx/sunplate/command"
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
	res := userCommand("%application", imp)
	exp := "skeleton"
	if res != exp {
		t.Errorf("Incorrect user command. Expected `%s`, got `%s`.", exp, res)
	}
}

func TestRun_IncorrectTask(t *testing.T) {
	c := "gxgbhsjdjdduuhdhsh"
	defer expectPanic(fmt.Sprintf(`Command "%s" does not exist. Panic expected.`, c))
	run(c)
}

func TestRun(t *testing.T) {
	run("sleep 10")
	cmd := run("sleep 1") // Starting a new instance.
	err := cmd.Wait()
	if err != nil {
		t.Error("Command is not started at the time of check.")
	}
}

func TestStart_EmptyFile(t *testing.T) {
	defer expectPanic("Application has been killed, panic expected.")
	notify <- syscall.SIGTERM
	start("run", command.Data{
		"run": "github.com/anonx/sunplate/run/testdata/empty",
	})
}

func TestStart(t *testing.T) {
	defer expectPanic("Application has been killed, panic expected.")
	notify <- syscall.SIGTERM
	start("run", command.Data{
		"run": "github.com/anonx/sunplate/run/testdata/nonempty",
	})
}

func TestRebuildFunc(t *testing.T) {
	ts := []string{`echo "task 1"`, `echo "task 2"`}
	as := []string{`echo "after 1"`, `echo "after 2"`}
	c := `echo "command"`
	fn := rebuildFunc(ts, as, c)
	fn()
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
