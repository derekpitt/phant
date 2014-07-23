package phant

func (c *Client) deleteUrl() string {
  return c.endpointPrefix + "streams/" + c.publicKey
}

// Delete will delete a stream. Cannot recover from this one!
func (c *Client) Delete(deleteKey string) error {
  request, err := c.createHttpRequestWithPrivateKey("DELETE", c.deleteUrl(), nil)

  if err != nil {
    return err
  }

  request.Header.Set("Phant-Delete-Key", deleteKey)

  _, err = doAndParseRequest(request)

  return err
}
