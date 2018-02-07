package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kiyonlin/newworld/pipeline"
)

func main() {
	const filename = "large.in"
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	defer writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)

	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 20 {
			break
		}
	}
}

func mergeDemo() {
	p :=
		pipeline.Merge(
			pipeline.InMemSort(
				pipeline.ArraySource(2, 13, 41, 5, 2, 1, 33)),
			pipeline.InMemSort(
				pipeline.ArraySource(12, 0, 42, 5, 21, 7, 3)))

	for v := range p {
		fmt.Println(v)
	}
}
