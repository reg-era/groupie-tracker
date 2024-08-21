package webserver

import (
	"html/template"
	"net/http"

	"GTapi/tracker"
)

func InfoHandle(w http.ResponseWriter, r *http.Request) {
	// check the reauest info error
	if r.URL.Path != "/getinfo" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	// // Get form TrackerServer Data to serve using the ID
	// Data, err := tracker.Get_Api_MoreData(r.PostFormValue("Art-ID"))
	// if err != nil {
	// 	http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
	// 	return
	// }

	// parse the web page template
	t, err := template.ParseFiles("./website/pages/moreinfo.html")
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
