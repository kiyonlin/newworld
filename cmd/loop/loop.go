package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(15),
		convertToBin(2),
		convertToBin(3),
	)

	fmt.Println("abc.txt contents:")
	printFile("../data/abc.txt")

	s := `abc "ad"
	kddd
	123
	p
	22333'`
	printFileContents(strings.NewReader(s))
}
