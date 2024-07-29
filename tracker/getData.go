package box

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ArtistsURL = "https://groupietrackers.herokuapp.com/api/artists"

var Data DataValu

func GData(i int) {
	var artApi []Api_artists
	var loca Api_attits_locatons
	var coner Api_artists_conertDates
	var relat Api_artists_relations

	var err error

	err = Decode(&artApi, ArtistsURL)

	err = Decode(&loca, artApi[i-1].Locations)
	err = Decode(&coner, artApi[i-1].ConcertDates)
	err = Decode(&relat, artApi[i-1].Relations)

	fmt.Println(
		artApi[i-1],
		"\n\n",
		loca,
		"\n\n",
		coner,
		"\n\n",
		relat,
	)

	

	Data.Image = artApi[i-1].Image
	Data.Name = artApi[i-1].Name

	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return 
	}
}

func Decode(mystruct interface{}, url string) error {
	resp, err := http.Get(url)
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
