package main

import (
	"errors"
	"fmt"
	"log"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			log.Println("Error occured:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do:%v", r))
		}
	}()

	panic(errors.New("this is an error"))
	b := 0
	a := 5 / b
	fmt.Println(a)
	panic(922)
}

func main() {
	tryRecover()
}
