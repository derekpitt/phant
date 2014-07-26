package phant

import (
	"testing"
)

func TestClear_Success(t *testing.T) {
	setup()
	defer teardown()

	handleSuccess()

	err := client.Clear()

	if err != nil {
		t.Error("expected no error")
	}
}

func TestClear_Fail(t *testing.T) {
	setup()
	defer teardown()

	handleError()

	err := client.Clear()

	if err == nil {
		t.Error("expected error")
	}
}
