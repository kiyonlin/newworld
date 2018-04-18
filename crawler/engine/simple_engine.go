package engine

import (
	"log"

	"github.com/kiyonlin/newworld/crawler/fetcher"
)

// SimpleEngine is a simple engine
type SimpleEngine struct {
}

// Run all request seeds
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)

		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	log.Println("Fetching ", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
