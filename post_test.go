package phant

import (
  "testing"
)

func TestPost_Success(t *testing.T) {
  setup()
  defer teardown()

  handleSuccess()

  err := client.Post(map[string]string{
    "pow": "pow",
  })

  if err != nil {
    t.Error("expected no error")
  }
}

func TestPost_Fail(t *testing.T) {
  setup()
  defer teardown()

  handleError()

  err := client.Post(map[string]string{
    "pow": "pow",
  })

  if err == nil {
    t.Error("expected error")
  }

  if err.Error() != "not ok" {
    t.Error("expected 'not ok' in .Error()")
  }
}
