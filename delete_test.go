package phant

import (
	"testing"
)

func TestDelete_Success(t *testing.T) {
	setup()
	defer teardown()

	handleSuccess()

	err := client.Delete("123")

	if err != nil {
		t.Error("expected no error")
	}
}

func TestDelete_Fail(t *testing.T) {
	setup()
	defer teardown()

	handleError()

	err := client.Delete("123")

	if err == nil {
		t.Error("expected error")
	}
}
