package generate

import (
	"os"
	"os/exec"
	"testing"
)

func TestStart_IncorrectSubcommand(t *testing.T) {
	defer expectPanic("requested incorrect subcommand and thus expected panic")
	start("generate", map[string]string{
		"generate": "commandThatDoesNotExist",
	})

	// Remove the directory we have created.
	os.RemoveAll("./assets")
}

func TestStart_Handlers(t *testing.T) {
	start("generate", map[string]string{
		"generate":  "handlers",
		"--input":   "./handlers/testdata/controllers",
		"--output":  "./assets/handlers",
		"--package": "handlers",
	})

	// TODO: check correctness of generated package.

	// Remove the directory we have created.
	os.RemoveAll("./assets")
}

func TestStart_Views(t *testing.T) {
	start("generate", map[string]string{
		"generate":  "views",
		"--input":   "./testdata/views",
		"--output":  "./testdata/assets/views",
		"--package": "views",
	})

	cmd := exec.Command("go", "run", "testdata/app/views/main.go")
	cmd.Stderr = os.Stderr // Show the output of the program we run.
	if err := cmd.Run(); err != nil {
		t.Errorf("There are problems with generated views, error: '%s'.", err)
	}

	// Remove the directory we have created.
	os.RemoveAll("./testdata/assets")
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
