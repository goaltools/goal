package handlers

import (
	"os"
	"os/exec"
	"testing"

	"github.com/colegion/goal/internal/command"
)

func TestStart(t *testing.T) {
	main(handlers, 0, command.Data{})

	cmd := exec.Command("go", "install", "github.com/colegion/goal/tools/generate/handlers/testdata/assets/handlers")
	cmd.Stderr = os.Stderr // Show the output of the program we run.
	if err := cmd.Run(); err != nil {
		t.Errorf(`There are problems with generated handlers, error: "%s".`, err)
	}

	// Remove the directory we have created.
	os.RemoveAll(*output)
}

var handlers []command.Handler

func init() {
	Handler.Flags.Set("input", "./testdata/controllers")
	Handler.Flags.Set("output", "./testdata/assets/handlers")

	handlers = []command.Handler{Handler}
}
