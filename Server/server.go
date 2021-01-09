package main

import (
	"fmt"
	"net/http"
)

const url = "http://localhost/"
const port = ":8090"

/*
type PageVariables struct {
	Date string
	Time string
}
*/

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func startpage(w http.ResponseWriter, req *http.Request) {

	var template string = "<h1>Startpage</h1><p>Welcome to your Golang Website</p>"
	fmt.Fprintf(w, template)

	/* More Adcanved with html/template package

	now := time.Now()              // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("startpage.html") //parse the html file homepage.html
	if err != nil {                                 // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	*/

}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", startpage)

	fmt.Println("Server is running on " + url + port)
	http.ListenAndServe(port, nil)
}
