package tools

import (
	"net/http"
	"strings"
)

var Data PageData

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ExecuteError(w, "Page not found", "404")
		return
	}

	op1 := GetOptions(Data)
	op2 := GetLocations(Data.Cards)
	Data.Notfound = false

	if r.Method == "POST" {
		value := r.PostFormValue("search")
		if value != "" {
			value = strings.ReplaceAll(value, " - first album date", "")
			value = strings.ReplaceAll(value, " - creation date", "")
			value = strings.ReplaceAll(value, " - members", "")
			value = strings.ReplaceAll(value, " - locations", "")

			var newcard []Card
			ids := SearchProcess(value)
			for _, i := range ids {
				newcard = append(newcard, Data.Cards[i])
			}
			if len(newcard) != 0 {
				ExecuteTemplate(w, PageData{Cards: newcard, Option: op1, Oplocation: op2, Notfound: false})
			} else {
				ExecuteTemplate(w, PageData{Cards: newcard, Option: op1, Oplocation: op2, Notfound: true})
			}
		} else {
			ExecuteTemplate(w, PageData{Cards: Data.Cards, Option: op1, Oplocation: op2, Notfound: false})
		}
	} else if r.Method == "GET" {
		ExecuteTemplate(w, PageData{Cards: Data.Cards, Option: op1, Oplocation: op2, Notfound: false})
	} else {
		ExecuteError(w, "Method not allowed", "405")
		return
	}
}
