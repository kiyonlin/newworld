package model

// SearchResult store the search result
type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextTo   int
	Items    []interface{}
}
