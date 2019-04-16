package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kataras/go-template/amber"

	tmpl "github.com/kataras/go-template"
)

type LandingPageData struct {
	Name string
}

var landingTemplate = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, req *http.Request) {
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

	err := landingTemplate.Execute(w, data)
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
	if secret == "thaksin" {
		sendResponse(w, "ไอสัส กูอีกแล้วเหรอ")
		return
	}

	if secret == QUERY_SECRET {
		sendResponse(w, "You discovered our secret! 🐳")
		return
	}

	sendResponse(w, "Everything is fine! 💖")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{"Name": "You!"}

	err := tmpl.ExecuteWriter(w, "hello.amber", data)
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
	}
}

func main() {
	tmpl.AddEngine(amber.New()).Directory(".", ".amber")
	err := tmpl.Load()
	if err != nil {
		panic("While parsing the template files: " + err.Error())
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	fmt.Println("Starting the server on Port", PORT)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/hello", helloHandler)

	err = http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Println("Unable to start the server:", err)
	}
}
