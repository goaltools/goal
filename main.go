// Package main is used as an entry point of
// the framework. It validates user input parameters
// and runs subcommands (aka tools).
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/colegion/goal/internal/log"
	"github.com/colegion/goal/tools/create"
	"github.com/colegion/goal/tools/generate/handlers"
	"github.com/colegion/goal/tools/run"
	"github.com/colegion/goal/utils/tool"
)

var trace = flag.Bool("trace", true, "show stack trace in case of runtime errors")

// tools is a stores information about the registered subcommands (tools)
// the framework supports.
var tools = tool.NewContext(
	create.Handler,
	run.Handler,

	handlers.Handler,
)

func main() {
	// Do not show stacktrace if something goes wrong
	// but tracing is disabled.
	defer func() {
		if err := recover(); err != nil {
			if *trace {
				var buffer bytes.Buffer
				buffer.WriteString(fmt.Sprintf("[panic]%v", err))
				for i := 1; ; i += 1 {
					pc, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					funcinfo := runtime.FuncForPC(pc)
					if nil != funcinfo {
						buffer.WriteString(fmt.Sprintf("    %s:%d %s\r\n", file, line, funcinfo.Name()))
					} else {
						buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
					}
				}

				log.Warn.Fatalf("TRACE: %v.", buffer.String())
			}
			os.Exit(0)
		}
	}()

	// Try to run the command user requested.
	// Ignoring the first argument as it is name of the executable.
	flag.Parse()
	err := tools.Run(flag.Args())
	if err != nil {
		log.Warn.Printf(unknownCmd, err, os.Args[0])
	}
}

var unknownCmd = `Error: %v.
Run "%s help" for usage.`
