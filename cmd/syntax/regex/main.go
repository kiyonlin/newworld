package main

import (
	"fmt"
	"regexp"
)

const text = "My email is jack@gmail.com " +
	"My email is aa@gmail.com " +
	"My email js jack@gmail.com.com"

func main() {
	re := regexp.MustCompile("([a-zA-Z0-9]+)@([a-zA-z0-9]+)\\.([a-zA-Z0-9.]+)")
	amtch := re.FindAllStringSubmatch(text, -1)
	fmt.Println(amtch)
}
