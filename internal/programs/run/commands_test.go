package run

import (
	"testing"
)

func TestStart_IncorrectTask(t *testing.T) {
	ts := []string{
		"gxgbhsjdjdduuhdhsh",
	}
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
		"gxgbhsjdjdduuhdhsh",
	}
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
	go instanceController()

	startSingleInstance("smth", c)

	channel <- message{action: "exit"}
	<-stopped
}

func TestStartSingleInstance(t *testing.T) {
	go instanceController()

	startSingleInstance("app", "sunplate run github.com/anonx/sunplate/internal/skeleton")
	startSingleInstance("smth", "sunplate help")
	startSingleInstance("app", "sunplate run github.com/anonx/sunplate/internal/skeleton")

	channel <- message{action: "exit"}
	<-stopped
}
