package main

import (
	"fmt"
	"net/http"
)

const url = "http://localhost"
const port = ":8090"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func startpage(w http.ResponseWriter, req *http.Request) {

	var template string = "<h1>Startpage</h1><p>Welcome to your Golang Website</p>"

	fmt.Fprintf(w, template)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", startpage)

	fmt.Println("Server is running on " + url + port)
	http.ListenAndServe(port, nil)
}
