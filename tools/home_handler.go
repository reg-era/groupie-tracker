package tools

import (
	"log"
	"net/http"
	"strings"
)

var data PageData

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ExecuteError(w, "Page not found", "404")
		return
	}

	cards, err := FetchArtistData("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		ExecuteError(w, "Status Internal Server Error", "500")
		log.Printf("Error fetching artist data: %v", err)
		return
	}

	data = PageData{Cards: cards}
	GetOptions(data)

	if r.Method == "POST" {
		value := r.PostFormValue("search")
		if value != "" {
			value = strings.ReplaceAll(value, " - first album date", "")
			value = strings.ReplaceAll(value, " - creation date", "")
			value = strings.ReplaceAll(value, " - members", "")
			value = strings.ReplaceAll(value, " - locations", "")

			var newcard []Card
			ids := SearchProcess(value)
			for _, i := range ids {
				newcard = append(newcard, data.Cards[i])
			}
			if len(newcard) != 0 {
				ExecuteTemplate(w, PageData{Cards: newcard, Option: Options, Notfound: false})
			} else {
				ExecuteTemplate(w, PageData{Cards: newcard, Option: Options, Notfound: true})
			}
		} else {
			ExecuteTemplate(w, PageData{Cards: cards, Option: Options, Notfound: false})
		}
	} else {
		ExecuteTemplate(w, PageData{Cards: cards, Option: Options, Notfound: false})
	}
}
