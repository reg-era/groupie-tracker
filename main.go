package main

import (
	"log"
	"net/http"

	"GTapi/tracker"
	"GTapi/webserver"
)

var (
	API_KEY = "AIzaSyCCTAVP5kfJGMAH2KoX8qo-n7r90Iosbjg"
	API     = "https://groupietrackers.herokuapp.com/api"
)

func main() {
	port := ":8008"

	// fetch the Api content in another routine
	go tracker.API_Process(API, API_KEY)

	// handle web functions
	http.HandleFunc("/", webserver.HomeHandle)
	http.HandleFunc("/getinfo", webserver.InfoHandle)

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./website/style/"))))

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")

	// lanche the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
