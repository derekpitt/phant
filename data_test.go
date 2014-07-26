package phant

import (
	"fmt"
	"net/http"
	"testing"
)

type testType struct {
	Derek     string `json:"derek"`
	Test      string `json:"test"`
	Timestamp string `json:"timestamp"`
}

func TestAllData_ReturnACorrectResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/output/12345.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `[{"derek":"1","test":"2","timestamp":"2014-07-23T03:17:28.061Z"},{"derek":"3","test":"4","timestamp":"2014-07-23T03:17:28.061Z"}]`)
	})

	var results []testType
	err := AllData("12345", &results)

	if err != nil {
		t.Error("expected no error: " + err.Error())
	}

	if len(results) != 2 {
		t.Error("expected len(results) to be 2")
	}

	if results[0].Derek != "1" {
		t.Error("expected 1 in [0].Derek")
	}

	if results[0].Test != "2" {
		t.Error("expected 2 in [0].Test")
	}

	if results[1].Derek != "3" {
		t.Error("expected 1 in [0].Derek")
	}

	if results[1].Test != "4" {
		t.Error("expected 4 in [1].Test")
	}
}

func TestDataOnPage_ReturnACorrectResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/output/12345.json", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") != "2" {
			t.Error("expected page parameter to be 2")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `[{"derek":"1","test":"2","timestamp":"2014-07-23T03:17:28.061Z"},{"derek":"3","test":"4","timestamp":"2014-07-23T03:17:28.061Z"}]`)
	})

	var results []testType
	err := DataOnPage("12345", 2, &results)

	if err != nil {
		t.Error("expected no error: " + err.Error())
	}

	if len(results) != 2 {
		t.Error("expected len(results) to be 2")
	}

	if results[0].Derek != "1" {
		t.Error("expected 1 in [0].Derek")
	}

	if results[0].Test != "2" {
		t.Error("expected 2 in [0].Test")
	}

	if results[1].Derek != "3" {
		t.Error("expected 1 in [0].Derek")
	}

	if results[1].Test != "4" {
		t.Error("expected 4 in [1].Test")
	}
}

func TestAllData_HandlesA404(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/output/12345.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"success":false,"message":"notfound"}`)
	})

	var results []testType
	err := AllData("12345", &results)

	if err == nil {
		t.Error("expected error to not be nil")
	}

}

func TestDataOnPage_HandlesA404(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/output/12345.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"success":false,"message":"notfound"}`)
	})

	var results []testType
	err := DataOnPage("12345", 2, &results)

	if err == nil {
		t.Error("expected error to not be nil")
	}

}
