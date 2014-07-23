package phant

import (
  "strings"
)

func (c *Client) postURL() string {
  return c.endpointPrefix + "input/" + c.publicKey
}

// Post will post a map of strings to phant
func (c *Client) Post(fields map[string]string) error {
  bodyReader := strings.NewReader(convertMapStringStringToURLValues(fields).Encode())

  request, err := c.createHTTPRequestWithPrivateKey("POST", c.postURL(), bodyReader)

  if err != nil {
    return err
  }

  _, err = doAndParseRequest(request)

  return err
}
