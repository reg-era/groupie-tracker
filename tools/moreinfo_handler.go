package tools

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Bandinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ExecuteError(w, "Method Not allowed", "405")
		return
	}

	id := strings.TrimPrefix(r.URL.RawQuery, "=id")
	i, err := strconv.Atoi(id)
	if err != nil || id == "" || (i <= 0 || i > 52) {
		ExecuteError(w, "Bad Request", "400")
		return
	}

	card_with_geo := addLocations(data.Cards[i-1].Locations, i-1)
	t, err := template.ParseFiles("templates/bandsinfo.html")
	if err != nil {
		ExecuteError(w, "Status Internal Server Error", "500")
		return
	}

	t.ExecuteTemplate(w, "bandsinfo.html", card_with_geo)
}
