package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"

	box "box/tracker"
)

func main() {
	open("http://localhost:8080/")
	http.HandleFunc("/", root)

	fmt.Println("Server started on port 8080...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func open(url string) {
	err := exec.Command("open", url).Start()
	if err != nil {
		println("Error:", err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./website/pages/artists.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	box.GData(5)
	tmp.Execute(w, box.Data)
}
