package phant

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCreate_ReturnACorrectResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `{"success":true,"publicKey":"123","privateKey":"456","deleteKey":"789"}`)
	})

	res, err := CreateStream("title", "description", []string{}, []string{}, true)

	if err != nil {
		t.Error("expected no error")
	}

	if res.DeleteKey != "789" {
		t.Error("delete key not set")
	}

	if res.PrivateKey != "456" {
		t.Error("private key not set")
	}

	if res.PublicKey != "123" {
		t.Error("public key not set")
	}
}

func TestCreate_ReturnAnError(t *testing.T) {
	setup()
	defer teardown()

	handleError()

	_, err := CreateStream("title", "description", []string{}, []string{}, true)

	if err == nil {
		t.Error("expected error")
	}

	if err.Error() != "not ok" {
		t.Error("expected 'not ok' in error")
	}
}
