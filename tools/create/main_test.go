package create

import (
	"testing"

	"github.com/colegion/goal/internal/command"
)

func TestMain_NoArgs(t *testing.T) {
	for _, v := range [][]string{[]string{}, nil} {
		err := main([]command.Handler{Handler}, 0, v)
		if err == nil {
			t.Errorf("No arguments passed. Error expected, got nil.")
		}
	}
}

func TestMain_RelPath(t *testing.T) {
	for _, v := range [][]string{[]string{"./something"}, []string{"something\\cool"}} {
		err := main([]command.Handler{Handler}, 0, v)
		if err == nil {
			t.Errorf("Relative path requested. Error expected, got nil.")
		}
	}
}
