package generation

import (
	"testing"
)

func TestStart_IncorrectSubcommand(t *testing.T) {
	defer expectPanic("requested incorrect subcommand and thus expected panic")
	Start("generate", map[string]string{
		"generate": "commandThatDoesNotExist",
	})
}

func TestStart_Handlers(t *testing.T) {
	Start("generate", map[string]string{
		"generate": "handlers",
	})
}

func TestStart_Listing(t *testing.T) {
	Start("generate", map[string]string{
		"generate": "listing",
	})
}

func expectPanic(msg string) {
	if err := recover(); err == nil {
		panic(msg)
	}
}
