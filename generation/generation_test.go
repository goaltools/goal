package generation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStart_IncorrectSubcommand(t *testing.T) {
	defer expectPanic("requested incorrect subcommand and thus expected panic")
	Start("generate", map[string]string{
		"generate": "commandThatDoesNotExist",
	})

	// Remove the directory we have created.
	os.RemoveAll(filepath.Join("./assets"))
}

func TestStart_Handlers(t *testing.T) {
	Start("generate", map[string]string{
		"generate": "handlers",
	})

	// Remove the directory we have created.
	os.RemoveAll(filepath.Join("./assets"))
}

func TestStart_Listing(t *testing.T) {
	Start("generate", map[string]string{
		"generate": "listing",
	})

	// Remove the directory we have created.
	os.RemoveAll(filepath.Join("./assets"))
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
