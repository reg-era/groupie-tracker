package webserver

import (
	"net/http"

	"GTapi/tracker"
)

var Options = make(map[string]int)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	// check the reauest info error
	if r.URL.Path != "/" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		ExecuteTemplate(w, tracker.Artists)
	case "POST":
		GetOptions(tracker.Artists)
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
			ExecuteTemplate(w, data)
		}
	default:
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}
}
