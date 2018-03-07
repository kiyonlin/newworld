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

// Poster can post to an url
type Poster interface {
	Post(url string, from map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "jack",
		"course": "golang",
	})
}

// RetrieverPoster combines Retriever and Poster
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "fake url"}
	r = &retriever
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

	fmt.Println("try a session with mockRetriever")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)

	fmt.Printf("%T %v\n", r, r)
	fmt.Println(r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
