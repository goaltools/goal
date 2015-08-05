package run

import (
	"os"
	"os/exec"

	"github.com/anonx/sunplate/log"
)

var (
	// stopped is a channel that is used for notifying the main program
	// that all subprograms have been terminated.
	stopped = make(chan bool, 1)

	// channel is used for communication with a user tasks starter
	// and their instances controller.
	channel = make(chan message, 1)
)

type message struct {
	action string
	name   string
	task   string
}

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
func startSingleInstance(name, task string) {
	channel <- message{
		action: "start",
		name:   name,
		task:   task,
	}
}

// instanceController is a function that is expected to be run
// as a separate goroutine. It starts and stops instances
// of user apps.
func instanceController() {
	// Clean up on termination.
	defer func() {
		if err := recover(); err != nil {
			stopped <- true
		}
	}()

	// terminate is used for killing an instance of a task.
	var terminate = func(name string, cmd *exec.Cmd) {
		pid := cmd.Process.Pid
		err := cmd.Process.Kill()
		log.AssertNil(err)

		cmd.Process.Wait()
		log.Trace.Printf(`Active instance of "%s" (PID %d) has been terminated.`, name, pid)
		cmd.Process = nil
	}

	// Waiting till we are asked to run/restart some tasks or exit
	// and following the orders.
	commands := map[string]*exec.Cmd{}
	for {
		switch m := <-channel; m.action {
		case "start":
			// Check whether we have already had an instance of the
			// requested task.
			cmd, ok := commands[m.name]
			if ok {
				// If so, terminate it first.
				terminate(m.name, cmd)
			}

			// If this is the first time this command is requested
			// to be run, initialize things.
			if !ok {
				n, as := parseTask(m.task)
				log.Trace.Printf(`Preparing "%s"...`, n)
				commands[m.name] = exec.Command(n, as...)
				commands[m.name].Stderr = os.Stderr
				commands[m.name].Stdout = os.Stdout
			}

			// Starting the task.
			t := replaceVars(m.task)
			log.Trace.Printf("Starting a new instance of `%s`...", t)
			err := commands[m.name].Start()
			if err != nil {
				log.Error.Panicf("Failed to start a command `%s`, error: %v.", t, err)
			}
		case "exit":
			for name, cmd := range commands {
				terminate(name, cmd)
			}
			stopped <- true
		}
	}
}
