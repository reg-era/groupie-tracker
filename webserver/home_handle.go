package webserver

import (
	"html/template"
	"net/http"

	"GTapi/tracker"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
	}

	t, err := template.ParseFiles("./website/pages/home.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, &tracker.Data)
	if err != nil {
		http.Error(w, "Method Not Allowed: error 500", http.StatusInternalServerError)
		return
	}
}
