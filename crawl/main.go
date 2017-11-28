package main

import (
	"fmt"
	"log"
	"os"
)

const usage = `
usage:
	crawl <starting-url>
`

// Channel is like a message box in a buffered case
// Unbuffered is you wait until the other GoRoutine takes the message out of your hand

type JobResult struct {
	URL   string
	PL    *PageLinks
	Error error
}

func reportResults(result *JobResult, results chan *JobResult) {
	log.Printf("reporting results for %s", result.URL)
	results <- result // Write the first result into the channel results
}

func startWorking(toFetch chan string, results chan *JobResult) {
	for URL := range toFetch {
		log.Printf("crawling %s", URL)
		links, err := GetPageLinks(URL)
		result := &JobResult{URL, links, err}
		go reportResults(result, results) // Use go because we are going to reportResults on a brand new goroutine that only lasts as long as it take to put the resultsi nto the go routine
	}
}

//numWorkers is the number of worker goroutines
//we will start: begin with just 1 and increase
//to see the benefits of concurrent execution,
//but don't increase beyond the number of concurrent
//socket connections allowed by your OS
const numWorkers = 20

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	//use the first argument as our starting URL
	startingURL := os.Args[1]

	//TODO: build a concurrent web crawler
	//with `numWorkers` worker goroutines,
	//using a channel to pass URLs to fetch
	//form the main goroutine to the workers,
	//and a channel to pass *PageLinks structs
	//from the workers back to the main goroutine.
	//Use the `GetPageLinks()` function in `links.go`
	//from your worker goroutines to fetch links
	//for a given URL.
	toFetch := make(chan string)
	results := make(chan *JobResult)
	seen := map[string]bool{} // Have we seen the result?

	for i := 0; i < numWorkers; i++ {
		go startWorking(toFetch, results)
	}

	seen[startingURL] = true
	toFetch <- startingURL
	outstandingJobs := 1

	for result := range results {
		outstandingJobs--
		if result.Error != nil {
			log.Printf("error crawling %s: %v", result.URL, result.Error)
			continue
		}
		log.Printf("processing %d links found in %s", len(result.PL.Links), result.URL)

		for _, URL := range result.PL.Links {
			if !seen[URL] {
				seen[URL] = true
				log.Printf("adding %s to the queue", URL)
				toFetch <- URL // Add it to the URls channel
				outstandingJobs++
			}
		}
		if outstandingJobs == 0 {
			log.Println("All done")
			return
		}
	}
}
