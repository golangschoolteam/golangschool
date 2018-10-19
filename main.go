package main

import (
	"log"
	"net/http"

	"blog/router"
)

func main() {
	router.Routes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("Server Start Failed")
	} else {
		log.Println("Listening on 0.0.0.0:8090")
	}
}