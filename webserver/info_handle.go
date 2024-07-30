package webserver

import (
	"html/template"
	"net/http"

	"GTapi/tracker"
)

func InfoHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getinfo" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}
	Data := tracker.Get_Artist_MoreData(r.PostFormValue("Art-ID"))
	t, err := template.ParseFiles("./website/pages/moreinfo.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Data)
	// fmt.Println(tracker.MoreInfos.Artist.Name)
	if err != nil {
		http.Error(w, "Method Not Allowed: error 500", http.StatusInternalServerError)
		return
	}
}
