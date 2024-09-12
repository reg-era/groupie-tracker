package tools

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var data PageData

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	apiURL := "https://groupietrackers.herokuapp.com/api"
	cards, err := FetchArtistData(apiURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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

func NotFound(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("404.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "404.html", nil)
}

func Bandinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
	}
	id := strings.TrimPrefix(r.URL.RawQuery, "=id")
	if id == "" {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil || (i <= 0 || i > 52) {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	card_with_geo := addLocations(data.Cards[i-1].Locations, i-1)

	t, err := template.ParseFiles("templates/bandsinfo.html")
	if err != nil {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	t.ExecuteTemplate(w, "bandsinfo.html", card_with_geo)
}

func ExecuteTemplate(w http.ResponseWriter, data any) {
	// parse the web page template
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	// execute the templat on web page with the data serve
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
