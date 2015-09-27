package command

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestRun_Default(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Default: true,
			Run: func(hs []Handler, i int, args []string) error {
				count++
				return nil
			},
		},
		Handler{ // Second one must be ignored.
			Default: true,
			Run: func(h []Handler, i int, args []string) error {
				count++
				return nil
			},
		},
	)
	if err := c.Run([]string{}); err != nil || count != 1 {
		t.Errorf("No arguments received: default handlers expected to be started.")
	}
}

func TestRun_NoDefault(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Run: func(hs []Handler, i int, args []string) error {
				count++
				return nil
			},
		},
	)
	if err := c.Run([]string{}); err != ErrIncorrectArgs || count != 0 {
		t.Errorf("No defaults defined and no arguments received: nothing was expected to be started.")
	}
}

func TestRun_NotFound(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Name: "run",
			Run: func(hs []Handler, i int, args []string) error {
				count++
				return errors.New("test")
			},
		},
		Handler{
			Name: "go generate",
			Run: func(hs []Handler, i int, args []string) error {
				count++
				return errors.New("test")
			},
		},
	)
	if err := c.Run([]string{"start --stuff xxx"}); count != 0 || err != ErrIncorrectArgs {
		t.Errorf(`Non-existent command requested. Expected "%s", got "%s".`, ErrIncorrectArgs, err)
	}
}

func TestRun(t *testing.T) {
	testErr := errors.New("Test error")
	c := NewContext(
		Handler{
			Name: "run",
			Run: func(hs []Handler, i int, args []string) error {
				return nil
			},
		},
		Handler{
			Name: "go generate",
			Run: func(hs []Handler, i int, args []string) error {
				return testErr
			},
		},
		Handler{
			Name: "new",
			Run: func(hs []Handler, i int, args []string) error {
				return nil
			},
		},
	)
	exp := "z"
	res := c.list[1].Flags.String("x", "default", "comment")
	if err := c.Run([]string{"go", "generate", "-x", exp}); err != testErr {
		t.Errorf(`Apparently, requested handler's entry function was not started. Got "%s".`, err)
	}
	if r := res; *r != exp {
		t.Errorf(`Incorrect value of flag. Expected "%s", got "%s".`, exp, *res)
	}
	if err := c.Run([]string{"go", "generate", "--incorrect", "flag"}); err == nil || err == testErr {
		t.Errorf(`Incorrect flag is used. Error expected, got "%s".`, err)
	}
}

func TestHandlerRequested(t *testing.T) {
	ts := map[string]struct {
		h    Handler
		args []string
		ok   bool
	}{
		"run": {
			h: Handler{
				Name: "run",
			},
			args: []string{},
			ok:   true,
		},
		"new": {
			h: Handler{
				Name: "run",
			},
			args: nil,
			ok:   false,
		},
		"generate stuff --something x --cool z": {
			h: Handler{
				Name: "generate stuff",
			},
			args: []string{"--something", "x", "--cool", "z"},
			ok:   true,
		},
		"generate": {
			h: Handler{
				Name: "generate stuff",
			},
			args: nil,
			ok:   false,
		},
	}
	for cmd, res := range ts {
		args := strings.Split(cmd, commandWordSep)
		if as, ok := res.h.requested(args); !reflect.DeepEqual(as, res.args) || ok != res.ok {
			t.Errorf(
				`Incorrect result. Expected "%v", "%v" got "%v", "%v".`,
				res.args, res.ok,
				as, ok,
			)
		}
	}
}
