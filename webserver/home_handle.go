package webserver

import (
	"fmt"
	"net/http"
	"strconv"

	"GTapi/tracker"
)

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
		value := r.PostFormValue("search")
		if value != "" {
			data, err := SearchHandele(value, tracker.Artists)
			if err != nil {
				http.Error(w, "Status Status Bad Request 400", http.StatusBadRequest)
				return
			} else {
				ExecuteTemplate(w, data)
			}
		}
	default:
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}
}

func SearchHandele(value string, data []tracker.Artist) ([]tracker.Artist, error) {
	id, err := strconv.Atoi(value)
	if err != nil || id <= 0 || id > len(data) {
		return nil, fmt.Errorf("Error")
	} else {
		return []tracker.Artist{data[id-1]}, nil
	}
}

// add func FilterHandele(&tracker.Artists) Artists
