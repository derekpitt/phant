package phant

import (
  "encoding/json"
)

// StatsResponse holds information about a stream
type StatsResponse struct {
  PageCount      int `json:"pageCount"`
  RemainingBytes int `json:"remaining"`
  UsedBytes      int `json:"used"`
  CapacityBytes  int `json:"cap"`
}

// Stats creates a stream
func Stats(publicKey string) (StatsResponse, error) {
  var statsRes = StatsResponse{}
  req, err := createHTTPRequest("GET", defaultEndpointPrefix+"output/"+publicKey+"/stats", nil)

  if err != nil {
    return statsRes, err
  }

  bodyBytes, err := getBodyBytes(req)

  if err != nil {
    return statsRes, err
  }

  err = json.Unmarshal(bodyBytes, &statsRes)

  if err != nil {
    return statsRes, err
  }

  return statsRes, nil
}
