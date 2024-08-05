package main

import (
	"log"
	"net/http"

	"GTapi/tracker"
	"GTapi/webserver"
)

//  Google map API KEY = AIzaSyCCTAVP5kfJGMAH2KoX8qo-n7r90Iosbjg

var API = "https://groupietrackers.herokuapp.com/api"

func main() {
	port := ":8080"

	// fetch the Api content in another routine
	go tracker.APiProcess(API)

	// handle web functions
	http.HandleFunc("/", webserver.HomeHandle)
	http.HandleFunc("/getinfo", webserver.InfoHandle)

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")

	// lanche the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
