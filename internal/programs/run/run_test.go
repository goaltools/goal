package run

import (
	"io/ioutil"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/anonx/sunplate/internal/command"
)

var mu sync.Mutex

func TestMain_TestData(t *testing.T) {
	createdFile := make(chan bool, 1)

	defer expectPanic(`Application was terminated, panic expected.`)
	defer os.Remove("./testdata/tmp.test")
	go func() {
		ioutil.WriteFile("./testdata/tmp.test", []byte{}, 0644)
		createdFile <- true
	}()
	go func() {
		time.Sleep(time.Second)
		<-createdFile
		notify <- syscall.SIGTERM
	}()
	main("run", command.Data{
		"run": "./testdata/configs",
	})
}

func TestMain_TestData2(t *testing.T) {
	defer expectPanic(`Application was terminated, panic expected.`)
	go func() {
		bs, err := ioutil.ReadFile("./sunplate.yml")
		if err != nil {
			t.Error(err)
		}
		err = ioutil.WriteFile("./sunplate.yml", bs, 0644)
		if err != nil {
			t.Error(err)
		}

		time.Sleep(time.Second * 4)
		notify <- syscall.SIGTERM
	}()
	main("run", command.Data{
		"run": "github.com/anonx/sunplate/internal/programs/run/testdata/configs",
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
