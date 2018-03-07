package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// Retriever store the info
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// Get info from url
func (r *Retriever) Get(url string) string {
	rep, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(rep, true)
	defer rep.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
