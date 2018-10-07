package main

import (
	"log"
	"net/http"
	"strings"
	"html/template"
)

type Visitor struct {
	IP string
}

func main() {

	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user_agent := request.UserAgent()
		log.Println("user_agent is " + user_agent)
		if strings.Contains(user_agent, "curl") {
			writer.Write([]byte("curl!"))
		} else {
			tmplt := template.New("index.go.html")       //create a new template with some name
			tmplt, _ = tmplt.ParseFiles("index.go.html")
			ip := "123.123.12.1"
			visitor := Visitor{IP: ip}
			tmplt.Execute(writer, visitor)
		}
	})

	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("API Go test!"))
	})

	log.Println("App is running on http://127.0.0.1:5000")
	http.ListenAndServe(":5000", nil)
}