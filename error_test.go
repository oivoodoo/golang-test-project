package music

import (
	"testing"
)

func TestAddNewError(t *testing.T) {
	error := new(Error)
	error.Add("error1")
	error.Add("error2")

	if len(error.Messages) != 2 {
		t.Errorf("Could not add new error message to the container")
	}
}

