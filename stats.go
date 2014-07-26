package phant

import (
	"strconv"
)

// StatsResponse holds information about a stream
type StatsResponse struct {
	PageCount      int
	RemainingBytes int
	UsedBytes      int
	CapacityBytes  int
}

type tempStatsResponse struct {
	PageCount      int    `json:"pageCount"`
	RemainingBytes int    `json:"remaining"`
	UsedBytes      int    `json:"used"`
	CapacityBytes  string `json:"cap"` // string for now until phant's next release
}

// Stats creates a stream
func Stats(publicKey string) (StatsResponse, error) {
	var statsRes = StatsResponse{}
	req, err := createHTTPRequest("GET", defaultEndpointPrefix+"output/"+publicKey+"/stats", nil)

	if err != nil {
		return statsRes, err
	}

	var tmp = tempStatsResponse{}
	err = decodeJsonAndClose(req, &tmp)

	if err != nil {
		return statsRes, err
	}

	// transfer over to statsRes
	statsRes.CapacityBytes, _ = strconv.Atoi(tmp.CapacityBytes)
	statsRes.PageCount = tmp.PageCount
	statsRes.RemainingBytes = tmp.RemainingBytes
	statsRes.UsedBytes = tmp.UsedBytes

	return statsRes, nil
}
