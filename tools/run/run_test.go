package run

import (
	"io/ioutil"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/colegion/goal/internal/command"
)

var mu sync.Mutex

func TestMain_TestData(t *testing.T) {
	createConfig(t)
	createdFile := make(chan bool, 1)

	defer expectPanic(`Application was terminated, panic expected.`)

	// Paths relative to the root directory are used here.
	defer os.Remove("../tmp.test")
	go func() {
		ioutil.WriteFile("../tmp.test", []byte{}, 0755)
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
	os.Chdir("../../") // Previous call of main() changed current dir.
	bs := createConfig(t)

	defer expectPanic(`Application was terminated, panic expected.`)
	go func() {
		// Paths relative to the root directory are used here.
		err := ioutil.WriteFile("./goal.yml", bs, 0755)
		if err != nil {
			t.Error(err)
		}

		time.Sleep(time.Second * 4)
		notify <- syscall.SIGTERM
	}()
	main("run", command.Data{
		"run": "github.com/colegion/goal/internal/programs/run/testdata/configs",
	})
}

func TestMain_IncorrectConfig(t *testing.T) {
	os.Chdir("../../") // Previous call of main() changed current dir.

	defer expectPanic(`A directory without configuration file. Panic expected.`)
	main("run", command.Data{
		"run": "./testdata", // Directory without config file.
	})
}

func TestMain(t *testing.T) {
	defer expectPanic(`Application was terminated, panic expected.`)
	notify <- syscall.SIGTERM
	main("run", command.Data{
		"run": "github.com/colegion/goal/internal/skeleton",
	})
}

func createConfig(t *testing.T) []byte {
	bs, err := ioutil.ReadFile("./testdata/configs/goal.src.yml")
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("./testdata/configs/goal.yml", bs, 0755)
	if err != nil {
		t.Error(err)
	}
	return bs
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
