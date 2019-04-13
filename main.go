package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type LandingPageData struct {
	Name string
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	name := req.URL.Query().Get("name")
	nameInPath := req.URL.Path[1:]
	isNameUnspecified := name == ""

	if isNameUnspecified && nameInPath != "" {
		name = nameInPath
	} else if isNameUnspecified {
		name = "World"
	}

	name = strings.Title(name)

	data := LandingPageData{
		Name: name,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalln("Cannot send response..?")
	}
}

type ApiResponse struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}

func sendResponse(w http.ResponseWriter, data string) {
	apiRes := ApiResponse{
		Status: "success",
		Data:   data,
	}

	res, err := json.Marshal(apiRes)
	if err != nil {
		w.WriteHeader(500)
		_, err = w.Write([]byte("Unable to marshal JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
}

const QUERY_SECRET = "jesuschrist420"

func apiHandler(w http.ResponseWriter, req *http.Request) {
	secret := req.URL.Query().Get("secret")
	if secret == QUERY_SECRET {
		sendResponse(w, "You discovered our secret! 🐳")
		return
	}

	sendResponse(w, "Everything is fine! 💖")
}

func main() {
	PORT := os.Getenv("PORT")

	fmt.Println("Starting the server on Port", PORT)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api", apiHandler)

	_ = http.ListenAndServe(":"+PORT, nil)
}
