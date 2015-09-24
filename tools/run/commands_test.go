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
		`goal start_smth`,
		`goal start_smth_else`,
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
		`goal run_smth`,
		`goal run_smth_else`,
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

	startSingleInstance("app", "goal run github.com/colegion/goal/internal/skeleton")
	startSingleInstance("smth", "goal help")
	startSingleInstance("app", "goal run github.com/colegion/goal/internal/skeleton")

	channel <- message{action: "exit"}
	<-stopped
}
