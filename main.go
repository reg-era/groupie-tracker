package main

import (
	"log"
	"net/http"

	"groupie-tracker-filter/pkg/funcs"
)

func main() {
	funcs.ParseJson()

	port := ":8000"

	// Handlers
	http.HandleFunc("/", funcs.HomeHandler)
	http.HandleFunc("/group/", funcs.GroupHandler)
	http.HandleFunc("/search", funcs.FilterHandler)

	// Serving static files
	http.Handle("/website/", http.StripPrefix("/website", http.FileServer(http.Dir("./website"))))

	log.Println("Server is running on http://localhost" + port)

	// Start server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
