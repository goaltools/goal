package command

import (
	"reflect"
	"strings"
	"testing"
)

func TestRun_Default(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Default: true,
			Run: func(hs []Handler, i int, args Data) {
				count++
			},
		},
		Handler{ // Second one must be ignored.
			Default: true,
			Run: func(h []Handler, i int, args Data) {
				count++
			},
		},
	)
	if err := c.Run([]string{}); err != nil || count != 1 {
		t.Errorf("No arguments received: single default handler expected to be started.")
	}
}

func TestRun_NoDefault(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Run: func(hs []Handler, i int, args Data) {
				count++
			},
		},
	)
	if err := c.Run([]string{}); err == nil || count != 0 {
		t.Errorf("No defaults defined and no arguments received: nothing was expected to be started.")
	}
}

func TestRun_NotFound(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Name: "run",
			Run: func(hs []Handler, i int, args Data) {
				count++
			},
		},
		Handler{
			Name: "go generate",
			Run: func(hs []Handler, i int, args Data) {
				count++
			},
		},
	)
	if err := c.Run([]string{"start --stuff xxx"}); count != 0 || err == nil {
		t.Errorf(`Non-existent command requested. Expected "nil", got "%s".`, err)
	}
}

func TestRun(t *testing.T) {
	count := 0
	c := NewContext(
		Handler{
			Name: "run",
			Run: func(hs []Handler, i int, args Data) {
			},
		},
		Handler{
			Name: "go generate",
			Run: func(hs []Handler, i int, args Data) {
				count++
			},
		},
		Handler{
			Name: "new",
			Run: func(hs []Handler, i int, args Data) {
			},
		},
	)
	exp := "z"
	res := c.list[1].Flags.String("x", "default", "comment")
	if err := c.Run([]string{"go", "generate", "-x", exp}); err != nil || count != 1 {
		t.Errorf(`Apparently, requested handler's entry function was not started. Got "%s".`, err)
	}
	if r := res; *r != exp {
		t.Errorf(`Incorrect value of flag. Expected "%s", got "%s".`, exp, *res)
	}
	if err := c.Run([]string{"go", "generate", "--incorrect", "flag"}); err == nil || count != 1 {
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
		if as, ok := res.h.Requested(args); !reflect.DeepEqual(as, res.args) || ok != res.ok {
			t.Errorf(
				`Incorrect result. Expected "%v", "%v" got "%v", "%v".`,
				res.args, res.ok,
				as, ok,
			)
		}
	}
}

func TestDataGetDefault(t *testing.T) {
	args := Data{}
	if v := args.GetDefault(1, "default"); v != "default" {
		t.Errorf(`Expected "default", got "%s".`, v)
	}

	args = Data{"0", "1", "2"}
	if v := args.GetDefault(10, "default"); v != "default" {
		t.Errorf(`Expected "default", got "%s".`, v)
	}
	if v := args.GetDefault(1, "1"); v != "1" {
		t.Errorf(`Expected "1", got "%s".`, v)
	}
}
