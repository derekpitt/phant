package phant

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"
)

var (
  mux       *http.ServeMux
  server    *httptest.Server
  client    *Client
  serverURL string
)

func setup() {
  mux = http.NewServeMux()
  server = httptest.NewServer(mux)
  serverURLParsed, _ := url.Parse(server.URL)
  serverURL = serverURLParsed.String() + "/"
  defaultEndpointPrefix = serverURL

  client = Create("public", "private")
}

func teardown() {
  server.Close()
}

func TestCreate(t *testing.T) {
  c := Create("public", "private")

  if c.publicKey != "public" {
    t.Error("expected publicKey to be public")
  }

  if c.privateKey != "private" {
    t.Error("expected privateKey to be private")
  }
}

// some helpers
func handleSuccess() {
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"success":true,"message":"ok"}`)
  })
}

func handleError() {
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, `{"success":false,"message":"not ok"}`)
  })
}

func TestParseErrorResponse(t *testing.T) {
  setup()
  defer teardown()

  handleError()

  req, err := createHTTPRequest("POST", serverURL, nil)

  if err != nil {
    t.Error("expected no error when creating request")
  }

  _, err = doAndParseRequest(req)

  if err == nil {
    t.Error("expected error in doAndParseRequest")
  }

  if err.Error() != "not ok" {
    t.Error("expected 'not ok' in .Error()")
  }

}
