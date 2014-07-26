package phant

import (
	"errors"
	"net/http"
	"strconv"
)

func dataAtURL(url string, into interface{}) (statusCode int, err error) {
	req, err := createHTTPRequest("GET", url, nil)

	if err != nil {
		return
	}

	return decodeJsonAndClose(req, &into)
}

// AllData will return all of a streams data in `into`
// `into` should be passed in as a pointer
func AllData(publicKey string, into interface{}) (err error) {
	statusCode, err := dataAtURL(defaultEndpointPrefix+"output/"+publicKey+".json", into)
	if statusCode != http.StatusOK {
		err = errors.New("Status Not OK")
	}

	return
}

// DataOnPage will return a page of a streams data in `into`
// Use Stats("...") to get how many pages are on a stream
// `into` should be passed in as a pointer
func DataOnPage(publicKey string, page int, into interface{}) (err error) {
	statusCode, err := dataAtURL(defaultEndpointPrefix+"output/"+publicKey+".json?page="+strconv.Itoa(page), into)

	if statusCode != http.StatusOK {
		err = errors.New("Status Not OK")
	}

	return
}
