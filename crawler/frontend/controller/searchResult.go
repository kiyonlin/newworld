package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kiyonlin/newworld/crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
)

// SearchResultHandler is a handler
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// ServeHTTP start a server
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	fmt.Fprintf(w, "q=%s,from=%d", q, from)
}

// CreateSearchResultHandler get a SearchResultHandler
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}
