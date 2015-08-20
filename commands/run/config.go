package run

import (
	"strings"

	"github.com/anonx/sunplate/log"

	"github.com/tsuru/config"
)

const (
	initSection  = "init"
	watchSection = "watch"
)

// conf is a parsed version of sunplate configuration file.
type conf struct {
	init  func()
	watch map[string]func() // Keys are patterns.
}

// parseConf parses a requested file and returns
// it in a form of conf structure.
func parseConf(file string) *conf {
	c := &conf{
		watch: map[string]func(){},
	}

	// Trying to read a configuration file..
	err := config.ReadConfigFile(file)
	log.AssertNil(err)

	// Extract init tasks, if presented.
	init, err := config.GetList(initSection)
	log.AssertNil(err)
	c.init = processTasksFn(init, initSection)

	// Extract patterns and tasks from watch section of config file.
	watch, err := config.Get(watchSection)
	log.AssertNil(err)
	switch typ := watch.(type) {
	case map[interface{}]interface{}:
		for key := range typ {
			pattern := key.(string)
			section := watchSection + ":" + pattern

			tasks, err := config.GetList(section)
			log.AssertNil(err)
			c.watch[pattern] = processTasksFn(tasks, section)
		}
	default:
		log.Warn.Printf(`No watch rules found in "%s".`, ConfigFile)
	}

	log.Trace.Printf(`Config file "%s" has been parsed.`, file)
	return c
}

// processTasksFn gets a list of tasks, processing them
// and returns a function that can be used for start
// of their execution.
func processTasksFn(tasks []string, section string) func() {
	log.Trace.Printf(`Processing section "%s" of configuration file...`, section)
	fns := []func(){}
	for i := range tasks {
		// We are parsing everthing first to show errors early,
		// not during runtime.
		fn := processTaskFn(tasks[i], section)
		fns = append(fns, fn)
	}
	return func() {
		for i := range fns {
			fns[i]()
		}
	}
}

// processTaskFn gets a task as a string, a section name where
// the task was found and returns a function
// that can be used for starting the task.
func processTaskFn(task, section string) func() {
	log.Trace.Printf("\t`%s`", task)
	name, args := parseTask(task)

	switch name {
	case "/start":
		assertSingleNonLoopArg(name, section, args)
		lst := listSection(name, args[0])
		return func() {
			start(lst)
		}
	case "/run":
		assertSingleNonLoopArg(name, section, args)
		lst := listSection(name, args[0])
		return func() {
			run(lst)
		}
	case "/single":
		assertSingleNonLoopArg(name, section, args)
		txt := textSection(name, args[0])
		return func() {
			startSingleInstance(args[0], txt)
		}
	case "/echo":
		return func() {
			log.Info.Println(strings.Join(args, " "))
		}
	case "/pass":
		if len(args) > 0 {
			log.Error.Panicf("%s: no arguments expected, got %v.", name, args)
		}
		return func() {
			// Do nothing.
		}
	}
	return func() {
		run([]string{task})
	}
}

// parseTask gets a string representation and returns
// a name of the command and arguments.
func parseTask(s string) (string, []string) {
	s = replaceVars(s) // Replace vars in a task to out values.
	ps := strings.Split(s, " ")

	// We are not checking the length of ps as
	// a garanteed minimum is 1.
	// tsuru/config returns <nil> instead of empty values.
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
// It panics otherwise.
func listSection(name, section string) []string {
	tasks, err := config.GetList(section)
	if err != nil {
		log.Error.Panicf(`Command "%s" expects section "%s" to be a list. Error: %v.`, name, section, err)
	}
	return tasks
}

// listSection gets a section, makes sure it is a single value
// section and returns it if everything is OK. It panics otherwise.
func textSection(name, section string) string {
	task, err := config.GetString(section)
	if err != nil {
		log.Error.Panicf(`Command "%s" expects section "%s" to be a single command. Error: %v.`, name, section, err)
	}
	return task
}
