package phant

func (c *Client) clearURL() string {
  return c.endpointPrefix + "input/" + c.publicKey + "/clear"
}

// Clear will clear the stream
func (c *Client) Clear() error {
  request, err := c.createHTTPRequestWithPrivateKey("POST", c.clearURL(), nil)

  if err != nil {
    return err
  }
  _, err = doAndParseRequest(request)

  return err
}
