// Package run has two main functions:
// - work as a task runner, watching files and
// rebuilding them if necessary;
// - works as a proxy server, that runs a user application,
// proxies all requests to it, and shows detailed
// error messages if needed (TODO).
package run

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/log"
	p "github.com/anonx/sunplate/path"
	"github.com/anonx/sunplate/watcher"

	"github.com/tsuru/config"
)

// ConfigFile is a name of the file that is located at the
// root of user project and describes what the test runner should do.
var ConfigFile = "sunplate.yml"

// Handler is an instance of run subcommand.
var Handler = command.Handler{
	Name:  "run",
	Info:  "start a task runner",
	Usage: "run [path]",
	Desc: `Run is a watcher and task runner. It uses sunplate.yml
file at the root of your project to find out what it should watch
and how to build / start your application.

Tasks of "init" section are run first, but only once per starting "sunplate run"
command. "watch" section is to inform that when some files in the specified
directories are modified, some tasks are expected to be executed.
`,

	Main: main,
}

var (
	notify = make(chan os.Signal, 1)
)

// main is an entry point of the command.
var main = func(action string, params command.Data) {
	defer func() {
		if err := recover(); err != nil {
			for k := range stopExpected {
				stopExpected[k] <- true
			}
			for k := range startExpected {
				<-startExpected[k]
			}
			log.Warn.Panic("Application has been terminated.")
		}
	}()

	imp := p.AbsoluteImport(params.Default(action, "./"))
	dir := p.PackageDir(imp)
	cf := filepath.Join(dir, ConfigFile)

	// Execute all commands from the requested directory.
	os.Chdir(dir)

	// Trying to read a configuration file..
	err := config.ReadConfigFile(cf)
	log.AssertNil(err)

	// Parsing configuration file and extracting the values
	// we need.
	log.Trace.Printf(`Starting to parse "%s"...`, cf)
	c := parseConf(cf)

	// Start init tasks.
	c.init()

	// Start watching the requested directories.
	w := watcher.NewType()
	for pattern := range c.watch {
		w.Listen(pattern, c.watch[pattern])
	}

	// Cleaning up after we are done.
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)
	<-notify
	log.Trace.Println("Application has been stopped.")
}
