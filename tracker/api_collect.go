package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Get_Api_Data(URL string) {
	req, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if err := json.Unmarshal(res, &Api); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func Get_Artist_Data(URL string) {
	req, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if err := json.Unmarshal(res, &Artists); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func Get_Artist_MoreData(id string) (MoreInfo, error) {
	inx, err := strconv.Atoi(id)
	if err != nil || inx > 52 || inx < 1 {
		return MoreInfos, fmt.Errorf("500")
	}

	for i, val := range URLS {
		req, err := http.Get(i + "/" + id)
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}
		if err := json.Unmarshal(res, val); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	}

	return MoreInfo{
		Artists[inx-1],
		Locations,
		Dates,
		Relations,
	}, nil
}
