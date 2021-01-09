package main

import (
	"fmt"
	"net/http"
	"time"
)

// Step 1: Run program without Groroutine
// Step 2: Add go to function to enable Goroutine
// Step 3: Add Wait Group to synchronize

//var wg sync.WaitGroup

func main() {

	//Measure program time
	start := time.Now()

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
		//go
		getWebsite(website)
		//wg.Add(1)
	}

	//wg.Wait()

	//Print time
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	duration := time.Since(start)
	fmt.Println(duration)
}

func getWebsite(website string) {

	//defer wg.Done()

	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")

	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}

}
