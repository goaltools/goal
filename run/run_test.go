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
		passCommand,
		`sunplate echo "Hello, world!"`,
		`sunplate echo "This is test. Ext is '%ext'"`,
	}
	execute(ts)
}

func TestRun_IncorrectTask(t *testing.T) {
	c := "gxgbhsjdjdduuhdhsh"
	defer expectPanic(fmt.Sprintf(`Command "%s" does not exist. Panic expected.`, c))
	run(c)
}

func TestRun(t *testing.T) {
	run("sunplate run github.com/anonx/sunplate/skeleton") // A program that runs forever.
	cmd := run(`sunplate echo "Starting a new instance"`)  // This should kill the previous instance.
	err := cmd.Wait()
	if err != nil {
		t.Error("Command is not started at the time of check.")
	}
}

func TestStart_EmptyFile(t *testing.T) {
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
	ts := []string{`sunplate echo "task 1"`, `sunplate echo "task 2"`}
	as := []string{`sunplate echo "after 1"`, `sunplate echo "after 2"`}
	c := `sunplate echo "command"`
	fn := rebuildFunc(ts, as, c)
	fn()
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
