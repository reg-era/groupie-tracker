package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const artists = "https://groupietrackers.herokuapp.com/api/artists"

type conutry struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {
	body := body()

	var conut []conutry
	err := json.Unmarshal([]byte(body), &conut)
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
}

func body() string {
	url, err := http.Get(artists)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	body, err := io.ReadAll(url.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(body)
}
