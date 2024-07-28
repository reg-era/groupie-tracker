package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const artists = "https://groupietrackers.herokuapp.com/api/artists"

type locatry struct{
	Id int
	Locations []string
	Dates  string
}

type conutry struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    locatry   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {
	body, errr := body()
	if errr != nil {
		fmt.Println(errr)
		return
	}
	var conut []conutry

	dec := json.NewDecoder(body)
	err := dec.Decode(&conut)
	// err := json.Unmarshal([]byte(body), &conut)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*for i := range conut{
		fmt.Printf(
			"Id: %d\nImage: %s\nName: %s\nMembers: %s\nCreationDate: %d\nFirstAlbum: %s\nLocations: %s\nConcertDates: %s\nRelations: %s",
			conut[i].Id, conut[i].Image, conut[i].Name, strings.Join(conut[i].Members, ", "), conut[i].CreationDate, conut[i].FirstAlbum, conut[i].Locations, conut[i].ConcertDates, conut[i].Relations,
		)
	}*/
	fmt.Println(conut)
}

func body() (io.Reader, error) {
	url, err := http.Get(artists)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}
	
	return url.Body, nil
}
