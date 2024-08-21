package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"GTapi/tracker"
	"GTapi/webserver"
)

var API = "https://groupietrackers.herokuapp.com/api"

func main() {
	start := time.Now()
	port := ":8080"

	// fetch the Api content in another routine
	go tracker.APiProcess(API)

	// serving style
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./website/style/"))))

	// handle web functions
	http.HandleFunc("/", webserver.HomeHandle)
	http.HandleFunc("/getinfo", webserver.InfoHandle)

	fmt.Println("time :", time.Since(start))
	return

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")
	// lanche the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
