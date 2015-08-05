package run

import (
	"fmt"
	"testing"
)

func TestStart_IncorrectTask(t *testing.T) {
	ts := []string{
		"gxgbhsjdjdduuhdhsh", // This is expected to cause a panic.
	}
	defer expectPanic(fmt.Sprintf(`Command "%s" does not exist. Panic expected.`, ts[0]))
	start(ts)
}

func TestStart(t *testing.T) {
	ts := []string{
		`sunplate start_smth`,
		`sunplate start_smth_else`,
	}
	start(ts)
}

func TestRun_IncorrectTask(t *testing.T) {
	ts := []string{
		"gxgbhsjdjdduuhdhsh", // This is expected to cause a panic.
	}
	defer expectPanic(fmt.Sprintf(`Command "%s" does not exist. Panic expected.`, ts[0]))
	run(ts)
}

func TestRun(t *testing.T) {
	ts := []string{
		`sunplate run_smth`,
		`sunplate run_smth_else`,
	}
	run(ts)
}

func TestStartSingleInstance_IncorrectCommand(t *testing.T) {
	c := "gxgbhsjdjdduuhdhsh"
	go func() {
		//defer expectPanic(fmt.Sprintf(`Command "%s" does not exist. Panic expected.`, c))
		instanceController()
	}()

	startSingleInstance("smth", c)
	<-stopped
}

func TestStartSingleInstance(t *testing.T) {
	go instanceController()

	startSingleInstance("app", "sunplate run github.com/anonx/sunplate/skeleton")
	startSingleInstance("smth", "sunplate help")
	startSingleInstance("app", "sunplate run github.com/anonx/sunplate/skeleton")

	channel <- message{action: "exit"}
	<-stopped
}
