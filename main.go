package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"GTapi/tracker"
	"GTapi/webserver"
)

var API = "https://groupietrackers.herokuapp.com/api"

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	port := ":8080"

	// fetch the Api content in another routine
	wg.Add(1)
	go tracker.APiProcess(&wg, API)

	// serving style
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./website/style/"))))

	// handle web functions
	http.HandleFunc("/", webserver.HomeHandle)

	wg.Wait()
	fmt.Println("time :", time.Since(start))

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")
	// lanche the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
