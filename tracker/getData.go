package box

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ArtistsURL = "https://groupietrackers.herokuapp.com/api/artists"

var (
	Data   []DataValu
	ArtApi []Api_artists
)

func GData(i int) {
	var loca Api_attits_locatons
	var coner Api_artists_conertDates
	var relat Api_artists_relations

	var err error

	err = Decode(&loca, ArtApi[i-1].Locations)
	err = Decode(&coner, ArtApi[i-1].ConcertDates)
	err = Decode(&relat, ArtApi[i-1].Relations)

	fmt.Println(
		ArtApi[i-1],
		"\n\n",
		loca,
		"\n\n",
		coner,
		"\n\n",
		relat,
	)

	Data[i-1].Image = ArtApi[i-1].Image
	Data[i-1].Name = ArtApi[i-1].Name

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
