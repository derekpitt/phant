package phant

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
)

// Client is the main type
type Client struct {
  publicKey  string
  privateKey string

  endpointPrefix string
}

type postResponse struct {
  Message string `json:"message"`
  Success bool   `json:"success"`
}

type postError struct {
  message string
}

func (p postError) Error() string {
  return p.message
}

const (
  defaultEndpointPrefix = "http://data.sparkfun.com/input/"
  version               = "derekpitt/phant/0.0.1"
)

// Create creates a client for phant
func Create(publicKey, privateKey string) *Client {
  return &Client{
    privateKey: privateKey,
    publicKey:  publicKey,

    endpointPrefix: defaultEndpointPrefix,
  }
}

func (c *Client) postUrl() string {
  return c.endpointPrefix + c.publicKey + ".json"
}

func convertMapStringStringToUrlValues(f map[string]string) url.Values {
  v := url.Values{}

  for fK, fV := range f {
    v.Set(fK, fV)
  }

  return v
}

// Post will post a map of strings to phant
func (c *Client) Post(fields map[string]string) error {
  bodyReader := strings.NewReader(convertMapStringStringToUrlValues(fields).Encode())

  client := &http.Client{}
  request, err := http.NewRequest("POST", c.postUrl(), bodyReader)

  if err != nil {
    return err
  }

  request.Header.Set("User-Agent", version)
  request.Header.Set("Phant-Private-Key", c.privateKey)
  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(request)

  if err != nil {
    return err
  }

  defer res.Body.Close()
  bodyBytes, err := ioutil.ReadAll(res.Body)

  if err != nil {
    return err
  }

  postRes := postResponse{}

  err = json.Unmarshal(bodyBytes, &postRes)

  if err != nil {
    return err
  }

  // look at the res
  if postRes.Success == false {
    return postError{postRes.Message}
  }

  return nil
}
