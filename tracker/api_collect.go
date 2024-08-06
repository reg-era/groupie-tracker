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
		return nil, fmt.Errorf("error")
	}
	// fetch for each Data the Api provide i is the Url and val is Data
	for i, val := range URLS {
		url := i + "/" + id
		Get_Api_Data(Info{url, val}) // fetching Data
	}
	// Get GEO localisation coordinates
	for i := 0; i < len(Locations.Locations); i++ {
		Get_Map_Coordinates(Locations.Locations[i])
		// loc := MapLoc{*MapLoc.Lat, lng, Locations.Locations[i]}
		fmt.Println(&MapLoc{})
		// Map = append(Map, loc)
	}

	mapCord, err := json.Marshal(Map)
	if err != nil {
		return nil, fmt.Errorf("err")
	}
	// return the morinfo data for a specific artist and nil err
	return &MoreInfo{
		Artists[inx-1],
		Dates,
		Relations,
		string(mapCord),
	}, nil
}

// https://maps.googleapis.com/maps/api/geocode/json?address=new_south_wales-australia&key=AIzaSyCCTAVP5kfJGMAH2KoX8qo-n7r90Iosbjg

func Get_Map_Coordinates(address string) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, API_KEY)
	Get_Api_Data(Info{url, &MapLoc{}})
}
