package engine

import (
	"log"

	"github.com/kiyonlin/newworld/crawler/fetcher"
)

// Run all request seeds
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Println("Fetching ", r.URL)
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v\n", r.URL, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
