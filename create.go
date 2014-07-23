package phant

import (
  "encoding/json"
  "strings"
)

type CreateResponse struct {
  Success bool   `json:"success"`
  Message string `json:"message"`

  PublicKey  string `json:"publickey"`
  PrivateKey string `json:"privateKey"`
  DeleteKey  string `json:"deleteKey"`
}

func CreateStream(title, description string, tags []string, fields []string, hidden bool) (CreateResponse, error) {
  var hiddenFormFalue = "0"
  if hidden == true {
    hiddenFormFalue = "1"
  }

  bodyReader := strings.NewReader(convertMapStringStringToUrlValues(map[string]string{
    "title":       title,
    "description": description,
    "hidden":      hiddenFormFalue,
    "tags":        strings.Join(tags, ","),
    "fields":      strings.Join(fields, ","),
    "check":       "",
  }).Encode())

  var createRes = CreateResponse{}
  req, err := createHttpRequest("POST", defaultEndpointPrefix+"streams", bodyReader)

  if err != nil {
    return createRes, err
  }

  bodyBytes, err := getBodyBytes(req)

  if err != nil {
    return createRes, err
  }

  err = json.Unmarshal(bodyBytes, &createRes)

  if err != nil {
    return createRes, err
  }

  if createRes.Success == false {
    return createRes, standardError{createRes.Message}
  }

  return createRes, nil
}
