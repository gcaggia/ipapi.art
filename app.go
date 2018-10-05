package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("App is running")
	log.Println("Curently listening on http://127.0.0.1:5000")
	http.ListenAndServe(":5000", nil)
}