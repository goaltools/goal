package run

import (
	"fmt"
	"strings"

	"github.com/colegion/goal/internal/log"
	"github.com/kylelemons/go-gypsy/yaml"
)

const (
	initSection  = "init"
	watchSection = "watch"
)

// conf is a parsed version of goal configuration file.
type conf struct {
	init  func()
	watch map[string]func() // Keys are patterns.

	m yaml.Map
}

// parseConf parses a requested file and returns
// it in a form of conf structure.
func parseConf(file string) *conf {
	var err error
	c := &conf{
		watch: map[string]func(){},
	}

	// Parse the configuration file..
	c.m, err = parseFile(file)
	if err != nil {
		log.Error.Panic(err)
	}

	// Extract init tasks.
	init, err := parseSlice(c.m, initSection)
	if err != nil {
		log.Error.Panic(err)
	}
	c.init, err = c.processTasksFn(init, initSection)
	if err != nil {
		log.Error.Panic(err)
	}

	// Extract patterns and tasks from watch section of config file.
	watch, err := parseMap(c.m, watchSection)
	if err != nil {
		log.Error.Panic(err)
	}
	for pattern, tasks := range watch {
		section := watchSection + ":" + pattern // It is used for debug messages.
		c.watch[pattern], err = c.processTasksFn(tasks, section)
		if err != nil {
			log.Error.Panic(err)
		}
	}

	log.Trace.Printf(`Config file "%s" has been parsed.`, file)
	return c
}

// processTasksFn gets a list of tasks, processing them
// and returns a function that can be used for start
// of their execution.
func (c *conf) processTasksFn(tasks []string, section string) (func(), error) {
	log.Trace.Printf(`Processing section "%s" of configuration file...`, section)
	fns := []func(){}
	for i := range tasks {
		// We are parsing everything first to show errors early,
		// not during runtime.
		fn, err := c.processTaskFn(tasks[i], section)
		if err != nil {
			return nil, err
		}
		fns = append(fns, fn)
	}
	return func() {
		for i := range fns {
			fns[i]()
		}
	}, nil
}

// processTaskFn gets a task as a string, a section name where
// the task was found and returns a function
// that can be used for starting the task.
func (c *conf) processTaskFn(task, section string) (func(), error) {
	log.Trace.Printf("\t`%s`", task)
	name, args := parseTask(task)

	switch name {
	case "/start":
		assertSingleNonLoopArg(name, section, args)
		lst, err := c.listSection(name, args[0])
		if err != nil {
			return nil, err
		}
		return func() {
			start(lst)
		}, nil
	case "/run":
		assertSingleNonLoopArg(name, section, args)
		lst, err := c.listSection(name, args[0])
		if err != nil {
			return nil, err
		}
		return func() {
			run(lst)
		}, nil
	case "/single":
		assertSingleNonLoopArg(name, section, args)
		lst, err := c.listSection(name, args[0])
		if err != nil {
			return nil, err
		}
		return func() {
			startSingleInstance(lst, args[0])
		}, nil
	case "/echo":
		return func() {
			log.Info.Println(strings.Join(args, " "))
		}, nil
	case "/pass":
		if len(args) > 0 {
			return nil, fmt.Errorf("%s: no arguments expected, got %v", name, args)
		}
		return func() {
			// Do nothing.
		}, nil
	}
	return func() {
		run([]string{task})
	}, nil
}

// parseTask gets a string representation and returns
// a name of the command and arguments.
func parseTask(s string) (string, []string) {
	s = replaceVars(s) // Replace vars in a task to our values.
	ps := strings.Split(s, " ")

	// We are not checking the length of ps as
	// a guaranteed minimum is 1.
	var as []string
	if len(ps) > 1 {
		as = ps[1:]
	}
	return ps[0], as
}

// assertSingleNonLoopArg gets a section and a list of arguments
// and makes sure the number of arguments is one and it is
// not equal to the current section.
func assertSingleNonLoopArg(name, section string, args []string) {
	if l := len(args); l != 1 {
		log.Error.Panicf(`%s: Incorrect number of arguments. Expected 1, got %d.`, name, l)
	}
	if args[0] == section {
		log.Error.Panicf(`%s: Use of "%s" as argument is not possible, loops are not allowed.`, name, section)
	}
}

// listSection gets a section, makes sure it is a list
// and returns its values if everything is OK.
func (c *conf) listSection(name, section string) ([]string, error) {
	tasks, err := parseSlice(c.m, section)
	if err != nil {
		return nil, fmt.Errorf(`Failed to parse command "%s" of "%s". Error: %v.`, name, section, err)
	}
	return tasks, nil
}
