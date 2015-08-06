package run

import (
	"syscall"
	"testing"

	"github.com/anonx/sunplate/internal/command"
)

func TestMain_TestData(t *testing.T) {
	defer expectPanic(`Application was terminated, panic expected.`)
	notify <- syscall.SIGTERM
	main("run", command.Data{
		"run": "./testdata/configs",
	})
}

func TestMain_IncorrectConfig(t *testing.T) {
	defer expectPanic(`A directory without configuration file. Panic expected.`)
	main("run", command.Data{
		"run": "./testdata", // Directory without config file.
	})
}

func TestMain(t *testing.T) {
	defer expectPanic(`Application was terminated, panic expected.`)
	notify <- syscall.SIGTERM
	main("run", command.Data{
		"run": "github.com/anonx/sunplate/internal/skeleton",
	})
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
