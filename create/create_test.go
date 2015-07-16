package create

import (
	"errors"
	"testing"
)

func TestWalkFunc_Error(t *testing.T) {
	_, fn := walkFunc("")
	TestError := errors.New("this is a test error")
	if err := fn("", nil, TestError); err != TestError {
		t.Errorf(`walkFunc expected to return TestError, returned "%s".`, err)
	}
}
