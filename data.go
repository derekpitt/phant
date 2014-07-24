package phant

import (
  "encoding/json"
  "strconv"
)

func dataAtURL(url string, into interface{}) (err error) {
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

// AllData will return all of a streams data in `into`
// `into` should be passed in as a pointer
func AllData(publicKey string, into interface{}) error {
  return dataAtURL(defaultEndpointPrefix+"output/"+publicKey+".json", into)
}

// DataOnPage will return a page of a streams data in `into`
// Use Stats("...") to get how many pages are on a stream
// `into` should be passed in as a pointer
func DataOnPage(publicKey string, page int, into interface{}) error {
  return dataAtURL(defaultEndpointPrefix+"output/"+publicKey+".json?page="+strconv.Itoa(page), into)
}
