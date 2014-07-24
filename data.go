package phant

import (
  "encoding/json"
  "strconv"
)

func dataAtUrl(url string, into interface{}) (err error) {
  req, err := createHTTPRequest("GET", url, nil)

  if err != nil {
    return
  }

  bodyBytes, err := getBodyBytes(req)

  if err != nil {
    return
  }

  err = json.Unmarshal(bodyBytes, &into)
  return
}

func AllData(publicKey string, into interface{}) error {
  return dataAtUrl(defaultEndpointPrefix+"output/"+publicKey+".json", into)
}

func DataOnPage(publicKey string, page int, into interface{}) error {
  return dataAtUrl(defaultEndpointPrefix+"output/"+publicKey+".json?page="+strconv.Itoa(page), into)
}
