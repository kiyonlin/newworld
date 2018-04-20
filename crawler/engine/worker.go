package engine

import (
	"log"

	"github.com/kiyonlin/newworld/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	// log.Println("Fetching ", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.URL), nil
}
