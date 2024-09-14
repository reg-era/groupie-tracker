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

	GetOptions(Data)
	Data.Option = Options
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
				Data.Notfound = false
				ExecuteTemplate(w, PageData{Cards: newcard, Option: Options, Notfound: false})
			} else {
				Data.Notfound = true
				ExecuteTemplate(w, Data)
			}
		} else {
			ExecuteTemplate(w, Data)
		}
	} else {
		ExecuteTemplate(w, Data)
	}
}
