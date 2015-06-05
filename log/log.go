// Package log is a simple wrapper on golang standard log
// package and agtorre/gocolorize (that colorizes the result on *nix systems).
// There are four predefined loggers.
// They are Error, Warn, Trace, and Info.
package log

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/anonx/sunplate/internal/github.com/agtorre/gocolorize"
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
	c gocolorize.Colorize
	w io.Writer
}

// Write is used for writing to predefined output using
// foreground color we want.
func (c *context) Write(p []byte) (n int, err error) {
	return c.w.Write([]byte(c.c.Paint(string(p))))
}

func init() {
	// Initialize default loggers.
	Error = log.New(
		&context{
			c: gocolorize.NewColor("red"),
			w: os.Stderr,
		}, "", 0,
	)
	Warn = log.New(
		&context{
			c: gocolorize.NewColor("yellow"),
			w: os.Stderr,
		}, "", 0,
	)
	Info = log.New(
		&context{
			c: gocolorize.NewColor("green"),
			w: os.Stdout,
		}, "", 0,
	)
	Trace = log.New(
		&context{
			c: gocolorize.NewColor("blue"),
			w: os.Stdout,
		}, "", 0,
	)

	// Do not use colorize when on windows.
	if runtime.GOOS == "windows" {
		gocolorize.SetPlain(true)
	}
}
