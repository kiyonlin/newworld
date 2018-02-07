package main

import (
	"fmt"

	"github.com/kiyonlin/newworld/pipeline"
)

func main() {
	p := pipeline.InMemSort(
		pipeline.ArraySource(2, 3, 4, 5, 2, 1, 33))

	for v := range p {
		fmt.Println(v)
	}
}
