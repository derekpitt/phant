package phant

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "net/url"
  "testing"
)

var (
  mux    *http.ServeMux
  server *httptest.Server
  client *Client
)

func setup() {
  mux = http.NewServeMux()
  server = httptest.NewServer(mux)

  client = Create("public", "private")
  url, _ := url.Parse(server.URL)
  client.endpointPrefix = url.String() + "/"
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

func TestSuccessPost(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"success":true,"message":"ok"}`)
  })

  err := client.Post(map[string]string{
    "pow": "pow",
  })

  if err != nil {
    t.Error("expected no error")
  }
}

func TestFailPost(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, `{"success":false,"message":"not ok"}`)
  })

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
