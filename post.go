package phant

import (
  "strings"
)

func (c *Client) postUrl() string {
  return c.endpointPrefix + "input/" + c.publicKey
}

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
