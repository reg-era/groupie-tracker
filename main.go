package main

import (
	"fmt"
	"net/http"

	help "tools/tools"
)

func ServeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}
	file := http.Dir("static")
	fs := http.StripPrefix("/static/", http.FileServer(file))
	fs.ServeHTTP(w, r)
}

func setupRoutes() {
	http.HandleFunc("/static/", ServeHandle)
	http.HandleFunc("/", help.Index)
	http.HandleFunc("/404", help.NotFound)
	http.HandleFunc("/bandsinfo", help.Bandinfo)
}

func main() {
	setupRoutes()
	Port := "localhost:8080"
	fmt.Println("Server is running at http://" + Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
