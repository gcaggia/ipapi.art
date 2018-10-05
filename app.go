package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user_agent := request.UserAgent()
		log.Println("user_agent is " + user_agent)
		if strings.Contains(user_agent, "curl") {
			writer.Write([]byte("curl!"))
		} else {

			http.ServeFile(writer, request, "public/" + request.URL.Path[1:])
		}
	})
	// fs := http.FileServer(http.Dir("public"))
	// http.Handle("/", fs)


	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("API Go test!"))
	})

	log.Println("App is running")
	log.Println("Curently listening on http://127.0.0.1:5000")
	http.ListenAndServe(":5000", nil)
}