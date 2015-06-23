// Package log is a simple wrapper on golang standard log
// package and terminal colorizer (that colorizes result on win and *nix systems).
// There are four predefined loggers. They are Error, Warn, Trace, and Info.
package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/anonx/sunplate/internal/github.com/fatih/color"
)

// A list of loggers that are used by "sunplate" packages.
var (
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Trace *log.Logger
)

// context stores information about logger output and its color.
type context struct {
	c *color.Color
	w io.Writer
}

// Write is used for writing to predefined output using
// foreground color we want.
func (c *context) Write(p []byte) (n int, err error) {
	// Set the color that has been registered for this logger.
	c.c.Set()
	defer color.Unset() // Don't forget to stop using after we're done.

	// Write the result to writer.
	return c.w.Write(p)
}

func init() {
	// Initialize default loggers.
	Error = log.New(
		&context{
			c: color.New(color.FgRed),
			w: os.Stderr,
		}, "", 0,
	)
	Warn = log.New(
		&context{
			c: color.New(color.FgYellow),
			w: os.Stderr,
		}, "", 0,
	)
	Info = log.New(
		&context{
			c: color.New(color.FgGreen),
			w: os.Stdout,
		}, "", 0,
	)
	Trace = log.New(
		&context{
			c: color.New(color.FgCyan),
			w: ioutil.Discard,
		}, "", 0,
	)
}
