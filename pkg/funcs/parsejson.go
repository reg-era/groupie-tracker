package funcs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// main struct
var All API

func ParseJson() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations := "https://groupietrackers.herokuapp.com/api/relation"
	ParseInfo(urlArtists, &All.Artists)
	ParseInfo(urlRelations, &All.Relation)
	for i, v := range All.Relation.Index {
		All.Artists[i].Rel = v.DatesLocations
	}
}

func ParseInfo(url string, temp interface{}) {
	res, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		return
	}
	text, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(text, &temp)
}
