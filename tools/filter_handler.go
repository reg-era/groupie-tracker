package tools

import (
	"net/http"
)

var Rel []string

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		ExecuteError(w, "Page not found", "404")
		return
	}
	if r.Method != "POST" {
		ExecuteError(w, "Method not allowed", "405")
		return
	}
	op1 := GetOptions(Data)
	op2 := GetLocations(Data.Cards)

	data := Filter(r.PostForm["member"], r.PostFormValue("firstAlbumMin"), r.PostFormValue("firstAlbumMax"), r.FormValue("creationDateMin"), r.FormValue("creationDateMax"), r.FormValue("location"))
	if len(data) == 0 {
		ExecuteTemplate(w, PageData{Cards: data, Option: op1, Oplocation: op2, Notfound: true})
	} else {
		ExecuteTemplate(w, PageData{Cards: data, Option: op1, Oplocation: op2, Notfound: false})
	}
}

func Filter(members []string, firstmin string, firstmax string, creatmin string, creatmax string, location string) []Card {
	var found []Card
	for _, v := range Data.Cards {
		if /*CreationDateVal(creatmax, creatmin, v.CreationDate) && FirstAlbumVal(firstmax, firstmin, v.FirstAlbum) &&*/ Checkloca(location, v.Locations) && MembersVal(members, len(v.Members)) {
			found = append(found, v)
		}
	}
	return found
}
