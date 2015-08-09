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

	"github.com/anonx/sunplate/internal/command"
	p "github.com/anonx/sunplate/internal/path"
	"github.com/anonx/sunplate/internal/watcher"
	"github.com/anonx/sunplate/log"

	"github.com/tsuru/config"
	"gopkg.in/fsnotify.v1"
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
	notify  = make(chan os.Signal, 1)
	restart = make(chan bool, 1)
)

// main is an entry point of the command.
var main = func(action string, params command.Data) {
	imp := p.AbsoluteImport(params.Default(action, "./"))
	dir := p.PackageDir(imp)
	cf := filepath.Join(dir, ConfigFile)

	// Start a user tasks runner and instances controller.
	go instanceController()

	// Start a configuration.
	go configDaemon(imp, cf)

	// Show user friendly errors and terminate subprograms
	// in case of panics.
	defer func() {
		channel <- message{
			action: "exit",
		}
		<-stopped
		log.Trace.Panicln("Application has been terminated.")
	}()

	// Execute all commands from the requested directory.
	os.Chdir(dir)

	// Load the configuration.
	reloadConfig()

	// Cleaning up after we are done.
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)
	<-notify
}

func configDaemon(imp, file string) {
	var watchers []*fsnotify.Watcher

	// closeWatchers is iterating over available watchers
	// and closes them.
	closeWatchers := func() {
		for i := range watchers {
			watchers[i].Close()
		}
		watchers = []*fsnotify.Watcher{}
	}
	defer closeWatchers() // Close watchers when we are done.

	for {
		// Wait till we are asked to reload the config file.
		<-restart

		// Closing old watchers to create new ones.
		closeWatchers()

		// Trying to read a configuration file..
		err := config.ReadConfigFile(file)
		if err != nil {
			log.Error.Printf(
				`Are you sure "%s" is a path of sunplate project?
"%s" file is missing.`, imp, file,
			)
			notify <- syscall.SIGTERM
			return
		}

		// Parsing configuration file and extracting the values
		// we need.
		log.Trace.Printf(`Starting to parse "%s"...`, file)
		c := parseConf(file)

		// Start init tasks.
		c.init()

		// Start watching the requested directories.
		w := watcher.NewType()
		watchers = append(watchers, w.ListenFile("./"+ConfigFile, reloadConfig))
		for pattern := range c.watch {
			watchers = append(watchers, w.Listen(pattern, c.watch[pattern]))
		}
	}
}

func reloadConfig() {
	restart <- true
}
