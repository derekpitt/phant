package phant

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStats_ReturnACorrectResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"pageCount":1,"remaining":2,"used":3,"cap":"4"}`)
	})

	res, err := Stats("123")

	if err != nil {
		t.Error("expected no error: " + err.Error())
	}

	if res.PageCount != 1 {
		t.Error("PageCount not set correctly")
	}

	if res.RemainingBytes != 2 {
		t.Error("RemainingBytes not set correctly")
	}

	if res.UsedBytes != 3 {
		t.Error("UsedBytes not set correctly")
	}

	if res.CapacityBytes != 4 {
		t.Error("CapacityBytes not set correctly")
	}
}
