package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kiyonlin/newworld/pipeline"
)

func main() {
	// fmt.Println(runtime.NumCPU())
	p := createNetworkPipeline("../pipeline-demo/large.in", 800000000, 4)
	writeToFile(p, "large.out")
	printFile("large.out")
	// p := createPipeline("../pipeline-demo/large.in", 800000000, 4)
	// writeToFile(p, "large.out")
	// printFile("large.out")
}


func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()

	chunkSize := fileSize / chunkCount

	sortResult := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)
		sortResult = append(sortResult, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResult...)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()

	chunkSize := fileSize / chunkCount

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResult := []<-chan int{}
	for _, addr := range sortAddr {
		sortResult = append(sortResult, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResult...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}


func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 20 {
			break
		}
	}
}