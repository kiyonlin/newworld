package main

import (
	"fmt"
	"io/ioutil"
	"time"
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

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go f1(ch1)
	go f2(ch2)
	select {
	case <-ch1:
		fmt.Println("The first case is selected.")
	case <-ch2:
		fmt.Println("The second case is selected.")
	}
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

func f1(ch chan int) {
	time.Sleep(time.Second * 1)
	ch <- 1
}

func f2(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 1
}
