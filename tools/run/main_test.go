package run

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/colegion/goal/utils/path"
	"github.com/colegion/goal/utils/tool"
)

var mu sync.Mutex

func TestMain_TestData(t *testing.T) {
	createConfig(t)
	createdFile := make(chan bool, 1)

	defer expectPanic(`Application was terminated, panic expected.`)

	// Paths relative to the root directory are used here.
	defer os.Remove("tmp.test")
	go func() {
		ioutil.WriteFile("tmp.test", []byte{}, 0666)
		createdFile <- true
	}()
	go func() {
		time.Sleep(time.Second)
		<-createdFile
		notify <- syscall.SIGTERM
	}()
	main(handlers, 0, tool.Data{"./testdata/configs"})
}

func TestMain_TestData2(t *testing.T) {
	createConfig(t)

	defer expectPanic(`Application was terminated, panic expected.`)
	go func() {
		createConfig(t)

		time.Sleep(time.Second * 4)
		notify <- syscall.SIGTERM
	}()
	main(handlers, 0, tool.Data{"github.com/colegion/goal/tools/run/testdata/configs"})
}

func TestMain_IncorrectConfig(t *testing.T) {
	defer expectPanic(`A directory without configuration file. Panic expected.`)
	notify <- syscall.SIGTERM

	// Directory without config file.
	main(handlers, 0, tool.Data{"./testdata"})
}

func TestMain(t *testing.T) {
	defer expectPanic(`Application was terminated, panic expected.`)
	go func() {
		time.Sleep(time.Second * 4)
		notify <- syscall.SIGTERM
	}()
	main(handlers, 0, tool.Data{"github.com/colegion/goal/internal/skeleton"})
}

func createConfig(t *testing.T) []byte {
	p, _ := path.ImportToAbsolute("github.com/colegion/goal/tools/run")

	bs, err := ioutil.ReadFile(
		filepath.Join(p, "./testdata/configs/goal.src.yml"),
	)
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile(
		filepath.Join(p, "./testdata/configs/goal.yml"), bs, 0666,
	)
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

var handlers = []tool.Handler{Handler}
