package main

import (
	"net/http"

	"github.com/kiyonlin/newworld/crawler/frontend/controller"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
