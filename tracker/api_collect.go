package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// fetch data from Url and give the struct to reload
func Get_Api_Data(info Info) {
	// get respons from api
	req, err := http.Get(info.Url)
	if err != nil {
		log.Fatalf("failed to get URL: %v", err)
	}
	defer req.Body.Close()
	// read the respons body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	// parse the body and store it in the structure provided
	if err := json.Unmarshal(res, info.Data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func Get_Api_MoreData(id string) (*MoreInfo, error) {
	// convert the Id to number for seraching in data
	inx, err := strconv.Atoi(id)
	if err != nil || inx > 52 || inx < 1 {
		return nil, fmt.Errorf("")
	}
	// fetch for each Data the Api provide i is the Url and val is Data
	for i, val := range URLS {
		url := i + "/" + id
		Get_Api_Data(Info{url, val}) // fetching Data
	}
	// return the morinfo data for a specific artist and nil err
	return &MoreInfo{
		Artists[inx-1],
		Locations,
		Dates,
		Relations,
	}, nil
}
