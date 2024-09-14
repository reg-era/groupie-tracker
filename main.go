package main

import (
	"fmt"
	"net/http"

	help "tools/tools"
)

func main() {
	Port := "localhost:8080"

	http.HandleFunc("/static/", help.ServeHandle)
	http.HandleFunc("/", help.Index)
	http.HandleFunc("/bandsinfo", help.Bandinfo)

	fmt.Println("Server is running at http://" + Port)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
