package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	if err := json.Unmarshal(res, &Data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

func Get_Artist_MoreData(URLS []string) {
	for i := 0; i < len(URLS); i++ {
		req, err := http.Get(URLS[i])
		fmt.Println(URLS[i])
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		if err := json.Unmarshal(res, &Datapls); err != nil {
			log.Fatalf("Error unmarshalling JSON: %v", err)
		}
	}
}
