package tools

import (
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, data any) {
	// parse the web page template
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ExecuteError(w, "Status Internal Server Error", "500")
		return
	}
	// execute the templat on web page with the data serve
	err = t.Execute(w, data)
	if err != nil {
		ExecuteError(w, "Status Internal Server Error", "500")
		return
	}
}

func ExecuteError(w http.ResponseWriter, msg string, stat string) {
	data := struct {
		NumError     string
		MessageError string
	}{
		NumError:     stat,
		MessageError: msg,
	}
	t, err := template.ParseFiles("templates/PageError.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
