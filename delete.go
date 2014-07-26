package phant

func (c *Client) deleteURL() string {
	return c.endpointPrefix + "streams/" + c.publicKey
}

// Delete will delete a stream. Cannot recover from this one!
func (c *Client) Delete(deleteKey string) error {
	request, err := c.createHTTPRequestWithPrivateKey("DELETE", c.deleteURL(), nil)

	if err != nil {
		return err
	}

	request.Header.Set("Phant-Delete-Key", deleteKey)

	_, err = doAndParseRequest(request)

	return err
}
