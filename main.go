package main

import (
	"log"
	"net/http"
	"os/exec"

	"GTapi/tracker"
	"GTapi/webserver"
)

func main() {
	port := ":8080"

	tracker.APiProcess("https://groupietrackers.herokuapp.com/api")

	http.HandleFunc("/", webserver.HomeHandle)
	http.HandleFunc("/getinfo", webserver.InfoHandle)

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")

	open("http://localhost" + port + "/")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func open(url string) {
	err := exec.Command("open", url).Start()
	if err != nil {
		println("Error:", err)
	}
}
