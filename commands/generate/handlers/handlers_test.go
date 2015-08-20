package handlers

import (
	"os"
	"os/exec"
	"testing"

	"github.com/anonx/sunplate/internal/command"
)

func TestStart(t *testing.T) {
	Start(command.Data{
		"--input":   "./testdata/controllers",
		"--output":  "./testdata/assets/handlers",
		"--package": "handlers",
	})

	cmd := exec.Command("go", "get", "-t", "./testdata/assets/handlers/...")
	cmd.Stderr = os.Stderr // Show the output of the program we run.
	if err := cmd.Run(); err != nil {
		t.Errorf(`There are problems with generated handlers, error: "%s".`, err)
	}

	// Remove the directory we have created.
	os.RemoveAll("./testdata/assets")
}
