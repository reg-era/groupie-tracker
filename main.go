package main

import (
	"log"
	"net/http"

	"GTapi/tracker"
	"GTapi/webserver"
)

var API = "https://groupietrackers.herokuapp.com/api"

// ********* there is probably 4s delai in the data fetching !!!!!!
func main() {
	port := ":8080"

	// fetch the Api content in another routine
	tracker.APiProcess(API)

	// serving style
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./website/style/"))))

	// handle web functions
	http.HandleFunc("/", webserver.HomeHandle)

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")
	// lanche the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
