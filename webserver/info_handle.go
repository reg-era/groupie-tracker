package webserver

import (
	"fmt"
	"html/template"
	"net/http"

	"GTapi/tracker"
)

func InfoHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getinfo" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
	}

	if r.Method != "post" {
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
	}

	t, err := template.ParseFiles("./website/pages/moreinfo.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, &tracker.Datapls)
	fmt.Println(tracker.Datapls[0])
	if err != nil {
		http.Error(w, "Method Not Allowed: error 500", http.StatusInternalServerError)
		return
	}
}
