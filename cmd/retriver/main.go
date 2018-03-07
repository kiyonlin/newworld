package main

import (
	"fmt"
	"time"

	"github.com/kiyonlin/newworld/retriever/mock"
	"github.com/kiyonlin/newworld/retriever/real"
)

const url = "http://www.imooc.com"

// Retriever can get info from url
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "fake url"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozile/5.0", TimeOut: time.Minute}
	inspect(r)
	// fmt.Println(download(r))

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not mock retriever")
	}

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
