package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func getCoordinates(location string) (float64, float64, bool) {
	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + location + "&key=AIzaSyCCTAVP5kfJGMAH2KoX8qo-n7r90Iosbjg"

	req, err := http.Get(url)
	if err != nil {
		return 0, 0, true
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return 0, 0, true
	}
	var geocode GeocodeResponse
	err = json.Unmarshal(res, &geocode)
	if err != nil {
		return 0, 0, true
	}
	locationData := geocode.Results[0].Geometry.Location
	return locationData.Lat, locationData.Lng, false
}

func addLocations(cities []string, id int) Card {
	for _, city := range cities {
		Lat, Lng, err := getCoordinates(city)
		if err {
			os.Exit(1)
		}
		geo := strconv.FormatFloat(Lat, 'f', -1, 64) + "," + strconv.FormatFloat(Lng, 'f', -1, 64)
		Data.Cards[id].Coordinates = append(Data.Cards[id].Coordinates, geo)
	}

	return Data.Cards[id]
}

func FetchArtistData(baseURL string) ([]Card, error) {
	var api APIindex
	if err := FetchData(baseURL, &api); err != nil {
		return nil, err
	}
	var artists []Artist
	if err := FetchData(api.Artists, &artists); err != nil {
		return nil, err
	}
	var location APIlocations
	if err := FetchData(api.Locations, &location); err != nil {
		return nil, err
	}
	var dates APIdates
	if err := FetchData(api.Dates, &dates); err != nil {
		return nil, err
	}
	var relation APIrelations
	if err := FetchData(api.Relations, &relation); err != nil {
		return nil, err
	}

	var cards []Card
	for i, artist := range artists {
		cards = append(cards, Card{
			Id:           artist.Id,
			Image:        artist.Image,
			Name:         artist.Name,
			Members:      artist.Members,
			CreationDate: artist.CreationDate,
			FirstAlbum:   artist.FirstAlbum,
			Locations:    location.Index[i].Location,
			ConcertDates: dates.Index[i].Dates,
			Relation:     relation.Index[i].DatesLocations,
		})
	}
	return cards, nil
}
