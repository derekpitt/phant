package phant

import (
  "encoding/json"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
)

// Client is the main type
type Client struct {
  publicKey  string
  privateKey string

  endpointPrefix string
}

type standardResponse struct {
  Message string `json:"message"`
  Success bool   `json:"success"`
}

type standardError struct {
  message string
}

func (e standardError) Error() string {
  return e.message
}

var defaultEndpointPrefix = "https://data.sparkfun.com/"

const (
  version = "derekpitt/phant/0.0.1"
)

// Create creates a client for phant
func Create(publicKey, privateKey string) *Client {
  return &Client{
    privateKey: privateKey,
    publicKey:  publicKey,

    endpointPrefix: defaultEndpointPrefix,
  }
}

func convertMapStringStringToUrlValues(f map[string]string) url.Values {
  v := url.Values{}

  for fK, fV := range f {
    v.Set(fK, fV)
  }

  return v
}

func createHttpRequest(reqType, url string, reader io.Reader) (*http.Request, error) {
  request, err := http.NewRequest(reqType, url, reader)

  if err == nil {
    request.Header.Set("User-Agent", version)
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Set("Accept", "application/json")
  }

  return request, err
}

func (c *Client) createHttpRequestWithPrivateKey(reqType, url string, reader io.Reader) (*http.Request, error) {
  request, err := createHttpRequest(reqType, url, reader)

  if err == nil {
    request.Header.Set("Phant-Private-Key", c.privateKey)
  }

  return request, err
}

func getBodyBytes(request *http.Request) ([]byte, error) {
  client := &http.Client{}
  res, err := client.Do(request)

  if err != nil {
    return nil, err
  }

  defer res.Body.Close()
  return ioutil.ReadAll(res.Body)
}

func doAndParseRequest(request *http.Request) (standardResponse, error) {
  postRes := standardResponse{}
  bodyBytes, err := getBodyBytes(request)

  if err != nil {
    return postRes, err
  }

  err = json.Unmarshal(bodyBytes, &postRes)

  if err != nil {
    return postRes, err
  }

  // look at the res
  if postRes.Success == false {
    return postRes, standardError{postRes.Message}
  }

  return postRes, nil
}
