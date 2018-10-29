package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"html/template"
)

type Visitor struct {
	IP string
}

func GetIP(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	// user_agent = r.UserAgent()
	return ip
}

func main() {

	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user_agent := request.UserAgent()
		ip := GetIP(request)
		log.Println("user_agent is " + user_agent)
		if strings.Contains(user_agent, "curl") {
			writer.Write([]byte("Your public ip is " + ip + " "))
		} else {
			absPath, _ := filepath.Abs("")
			tmplt, err := template.ParseFiles(absPath + "/templates/index.go.html")

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			visitor := Visitor{IP: ip}
			tmplt.Execute(writer, visitor)
		}
	})

	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		visitor := Visitor{IP: GetIP(request)}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(visitor)
	})

	log.Println("App is running on http://127.0.0.1:18000")
	http.ListenAndServe(":18000", nil)
}