package webserver

import (
	"net/http"

	"GTapi/tracker"
)

var Options = make(map[string]int)

type Asemble struct {
	Data     []tracker.Artist
	Option   map[string]int
	Notfound bool
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	GetOptions(tracker.Artists)
	// check the reauest info error
	if r.URL.Path != "/" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		ExecuteTemplate(w, Asemble{tracker.Artists, Options, false})
	case "POST":
		value := r.PostFormValue("search")
		if value != "" {
			var data []tracker.Artist
			if v, ok := Options[value]; ok {
				data = append(data, tracker.Artists[v])
			} else {
				ids := SearchPro(value)
				for _, i := range ids {
					data = append(data, tracker.Artists[i])
				}
			}
			if len(data) != 0 {
				ExecuteTemplate(w, Asemble{data, Options, false})
			} else {
				ExecuteTemplate(w, Asemble{data, Options, true})
			}
		}
	default:
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}
}
