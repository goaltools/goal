package run

import (
	"os"
	"os/exec"
	"sync"

	"github.com/anonx/sunplate/log"
)

var (
	// started indicates whether a previous instance of user
	// app with a specific section name already exists.
	started = map[string]bool{}

	// stopExpected, instanceStopped, and stopped are channels that
	// are used to let us know the user app has to be stopped or has
	// already been stopped, so we can safely start it again.
	stopExpected    = map[string]chan bool{}
	instanceStopped = map[string]chan bool{}
	stopped         = make(chan bool, 1)

	mu sync.Mutex
)

// start runs commands but does not wait for them to complete.
func start(tasks []string) {
	// Iterate over all available tasks.
	for i := range tasks {
		n, as := parseTask(tasks[i])
		cmd := exec.Command(n, as...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		log.Trace.Printf("Starting `%s`...", replaceVars(tasks[i]))
		err := cmd.Start()
		if err != nil {
			log.Error.Panicf(`Failed to start a command "%s", error: %v.`, tasks[i], err)
		}
	}
}

// run is used for starting commands and waiting
// for them to complete.
func run(tasks []string) {
	// Iterate over all available tasks.
	for i := range tasks {
		// Run and wait till we get an output.
		n, as := parseTask(tasks[i])
		cmd := exec.Command(n, as...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		log.Trace.Printf("Running `%s`...", replaceVars(tasks[i]))
		err := cmd.Run()
		if err != nil {
			log.Error.Panicf(`Failed to run a command "%s", error: %v.`, tasks[i], err)
		}
	}
}

// startSingleInstance starts an instance asynchronyously just
// as start. However, if there is already an instance with
// the same name, it will be stopped first
// before running a new one.
func startSingleInstance(name, task string) *exec.Cmd {
	// Initialize channels if we haven't done it yet.
	active := started[name]
	if !active {
		stopExpected[name] = make(chan bool, 1)
		instanceStopped[name] = make(chan bool, 1)
	}

	// Stopping the previous instance
	// if it already exists.
	if active {
		log.Trace.Printf(`Terminating the old instance of "%s"...`, name)
		stopExpected[name] <- true
		<-instanceStopped[name]
	}
	log.Trace.Printf(`Starting a new instance of "%s"...`, name)

	// Parse the input task, prepare a command.
	n, as := parseTask(task)
	cmd := exec.Command(n, as...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	log.Trace.Printf("\t`%s`", replaceVars(task))

	// Starting the user task.
	err := cmd.Start()
	if err != nil {
		log.Error.Panicf("Failed to start a command `%s`, error: %v.", task, err)
	}
	started[name] = true

	// Make sure we'll be able to stop the app.
	go func() {
		<-stopExpected[name]
		mu.Lock()
		defer mu.Unlock()

		pid := cmd.Process.Pid
		cmd.Process.Kill()
		cmd.Process.Wait()
		log.Trace.Printf("\tProcess with PID %d has been killed.", pid)

		delete(started, name)

		instanceStopped[name] <- true
		if len(started) == 0 {
			stopped <- true
		}
	}()

	return cmd
}
