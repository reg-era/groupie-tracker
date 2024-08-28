package webserver

import (
	"net/http"
	"strconv"

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
			}
			ExecuteTemplate(w, data)
		}
	default:
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}
}

func GetOptions(data []tracker.Artist) {
	for i, c := range data {
		Options[c.Name+" - artist/band"] = i
		Options[c.FirstAlbum+" - first album date"] = i
		Options[strconv.Itoa(c.CreationDate)+" - creation date"] = i
		for _, j := range c.Members {
			Options[j+" - members"] = i
		}
		for _, j := range c.LocationST.Locations {
			Options[j+" - locations"] = i
		}
	}
}

// add func FilterHandele(&tracker.Artists) Artists
