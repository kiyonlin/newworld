package engine

// Request store url and parse func
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult store the parse result, include requests and items
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// Item store fetched infos
type Item struct {
	URL     string
	ID      string
	Type    string
	Payload interface{}
}

// NilParser returns a nil ParseResult
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
