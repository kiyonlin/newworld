package mock

import "fmt"

// Retriever get contents
type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{contents=%s}", r.Contents)
}

// Post url from map
func (r *Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]
	return "ok"
}

// Get info from url
func (r *Retriever) Get(url string) string {
	return r.Contents
}
