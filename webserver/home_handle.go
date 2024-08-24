package webserver

import (
	"html/template"
	"net/http"

	"GTapi/tracker"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	// check the reauest info error
	if r.URL.Path != "/" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	// parse the web page template
	t, err := template.ParseFiles("./website/pages/home.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	// execute the templat on web page with the data serve
	err = t.Execute(w, &tracker.Artists)
	if err != nil {
		http.Error(w, "Method Not Allowed: error 500", http.StatusInternalServerError)
		return
	}
}

// add func SearchHandele(&tracker.Artists) Artists

// add func FilterHandele(&tracker.Artists) Artists
