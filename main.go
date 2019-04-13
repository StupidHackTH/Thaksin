package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var template = `
	<!doctype html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	 <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
	 <meta http-equiv="X-UA-Compatible" content="ie=edge">
	 <title>Secret Place</title>
	<style>
		body {
			display: flex;
			align-items: center;
			justify-content: center;
			min-height: 100vh;

			font-family: "Helvetica Neue", sans-serif;
		}

		h1 {
			font-size: 2.8em;
		}
	</style>
	</head>
	<body>
	  <h1>Hello, World! You've discovered my secret lair! ;)</h1>
	</body>
	</html>
`

func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, template)

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
		sendResponse(w, "You discovered the secret! 🐳")
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
