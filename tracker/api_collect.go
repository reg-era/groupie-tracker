package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Get_Api_Data(info Info) {
	req, err := http.Get(info.Url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if err := json.Unmarshal(res, info.Data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func Get_Api_MoreData(id string) (MoreInfo, error) {
	inx, err := strconv.Atoi(id)
	if err != nil || inx > 52 || inx < 1 {
		return MoreInfos, fmt.Errorf("")
	}

	for i, val := range URLS {
		url := i + "/" + id
		Get_Api_Data(Info{url, val})
	}

	return MoreInfo{
		Artists[inx-1],
		Locations,
		Dates,
		Relations,
	}, nil
}
