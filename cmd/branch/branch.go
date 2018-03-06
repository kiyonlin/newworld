package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "../data/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(grade(0),
		grade(53),
		grade(100),
		grade(99))
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score < 80:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
