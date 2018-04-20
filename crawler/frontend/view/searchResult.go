package view

import (
	"html/template"
	"io"

	"github.com/kiyonlin/newworld/crawler/frontend/model"
)

// SearchResultView is the view
type SearchResultView struct {
	template *template.Template
}

// CreateSearchResultView return the view
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}
}

// Render the html
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
