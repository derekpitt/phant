package phant

import (
  "strings"
)

// Post will post a map of strings to phant
func (c *Client) Post(fields map[string]string) error {
  bodyReader := strings.NewReader(convertMapStringStringToUrlValues(fields).Encode())

  request, err := c.createHttpRequestWithPrivateKey("POST", c.postUrl(), bodyReader)

  if err != nil {
    return err
  }

  _, err = doAndParseRequest(request)

  return err
}
