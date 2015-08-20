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
		t := replaceVars(tasks[i])
		log.Trace.Printf("Starting `%s`...", t)
		err := cmd.Start()
		if err != nil {
			log.Error.Printf(`Failed to start a command "%s", error: %v.`, t, err)
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
		t := replaceVars(tasks[i])
		log.Trace.Printf("Running `%s`...", t)
		err := cmd.Run()
		if err != nil {
			log.Error.Printf(`Failed to run a command "%s", error: %v.`, t, err)
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
	// terminate is used for killing an instance of a task.
	var terminate = func(name string, cmd *exec.Cmd) {
		if cmd.Process == nil {
			return
		}

		pid := cmd.Process.Pid
		err := cmd.Process.Kill()
		if err == nil {
			cmd.Process.Wait()
		}

		cmd.Process = nil // Manually set it to nil, so we can reuse the command.
		log.Trace.Printf(`Active instance of "%s" (PID %d) has been terminated.`, name, pid)
	}
	commands := map[string]*exec.Cmd{}

	// Clean up on termination.
	defer func() {
		for name, cmd := range commands {
			terminate(name, cmd)
		}
		stopped <- true
	}()

	// Waiting till we are asked to run/restart some tasks or exit
	// and following the orders.
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
				cmd = exec.Command(n, as...)
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout
			}

			// Starting the task.
			t := replaceVars(m.task)
			log.Trace.Printf("Starting a new instance of `%s`...", t)
			err := cmd.Start()
			if err != nil {
				log.Error.Printf("Failed to start a command `%s`, error: %v.", t, err)
			}

			// If this is the first time this command is requested
			// and the program has been started successfully, register it
			// so we'll be able to terminate it.
			if !ok {
				commands[m.name] = cmd
			}
		case "exit":
			return
		}
	}
}
