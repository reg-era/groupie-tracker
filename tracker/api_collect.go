package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Get_Api_Data(mystruct interface{}, url string) error { //my struct holds any data that comes fro mthe webpage
	resp, err := http.Get(url) //sends a get request to the url
	if err != nil {
		return fmt.Errorf("failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(mystruct); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	return nil
}

func Get_Artist_MoreData(id string) (interface{}, error) {
	inx, err := strconv.Atoi(id)
	if err != nil || inx > 52 || inx < 1 {
		return nil, fmt.Errorf("500")
	}

	for i, val := range URLS {
		req, err := http.Get(i + "/" + id)
		if err != nil {
			return nil, fmt.Errorf("Error fetching data: %v", err)
		}
		defer req.Body.Close()

		if err := json.NewDecoder(req.Body).Decode(val); err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %v", err)
		}
	}

	return MoreInfo{
		Artists[inx-1],
		Locations,
		Dates,
		Relations,
	}, nil
}
