package tracker

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
