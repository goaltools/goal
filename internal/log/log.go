// Package log is a simple wrapper around Go's standard log
// package and terminal colorizer that works both on win and *nix.
// It includes a few standard loggers. They are:
// Error, Warn, Info, and Trace.
package log

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

// Default loggers that are used by packages of Goal project.
var (
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Trace *log.Logger
)

// context is a type that implements io.Writer interface.
type context struct {
	c *color.Color
	w io.Writer
}

// Write is a method that's required for context type
// to be an implementation of io.Writer interface.
// It writes data to the predefined writer using
// the previously defined color.
func (c *context) Write(d []byte) (n int, err error) {
	// Set the color that has been registered for this logger.
	c.c.Set()
	defer color.Unset() // Stop using the color after we're done.

	// Write the result to writer.
	return c.w.Write(d)
}

// newContext allocates and returns a new context.
func newContext(w io.Writer, cs ...color.Attribute) *context {
	return &context{
		c: color.New(cs...),
		w: w,
	}
}

func init() {
	// Initialize default loggers.
	Error = log.New(newContext(os.Stderr, color.FgRed, color.Bold), "ERROR: ", log.Ltime)
	Warn = log.New(newContext(os.Stderr, color.FgYellow), "WARN: ", log.Ltime)
	Info = log.New(newContext(os.Stdout, color.FgGreen), "INFO: ", log.Ltime)
	Trace = log.New(newContext(os.Stdout, color.FgCyan), "TRACE: ", log.Ltime)
}
