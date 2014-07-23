package phant

func (c *Client) clearUrl() string {
  return c.endpointPrefix + "input/" + c.publicKey + "/clear"
}

// Clear will clear the stream
func (c *Client) Clear() error {
  request, err := c.createHttpRequestWithPrivateKey("POST", c.clearUrl(), nil)

  if err != nil {
    return err
  }
  _, err = doAndParseRequest(request)

  return err
}
