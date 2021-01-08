package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//Measure program time
	start := time.Now()

	c := make(chan string)

	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, website := range websites {
		go getWebsite(website, c)
	}

	//Iterating over the range of channel. So keeps receiving messages until channel is closed
	for msg := range c {
		fmt.Println(msg)
		duration := time.Since(start)
		fmt.Println(duration)
	}

	close(c)
	//Alternate syntax
	/*
		for {
			msg, open := <-c
			if !open {
				break
			}
			fmt.Println(msg)
		}
	*/

	//Print time
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	duration := time.Since(start)
	fmt.Println(duration)

}
func getWebsite(website string, c chan string) {
	if _, err := http.Get(website); err != nil {
		c <- website + "is down"

	} else {
		c <- website + "is up"
	}

}
