package tracker

import (
	"encoding/json"
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

func Get_Artist_MoreData(id string) MoreInfo {
	for i, val := range URLS {
		// fmt.Println(i+"/"+id, "-----------")
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
		// fmt.Println(val)
	}
	inx, _ := strconv.Atoi(id)

	mr := MoreInfo{
		Artists[inx],
		Locations,
		Dates,
		Relations,
	}
	// fmt.Println()
	// fmt.Println()
	// fmt.Println("finish")
	// fmt.Println(MoreInfos.Artist.Name)
	return mr
}
