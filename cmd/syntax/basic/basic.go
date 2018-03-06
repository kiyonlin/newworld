package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ss = "string"
	aa = 1
	bb = 2
)

//默认 没有初值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d,%q\n", a, s)
}

//赋予初值和类型
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//赋予初值
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 3
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3+4i
	//fmt.Println(cmplx.Abs(c))
	//cmplx.Pow(math.E,1i*math.Pi)+1

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	//a,b可以不指定类型
	c := math.Sqrt(a*a + b*b)
	fmt.Println(filename, c)

}

const (
	cpp = iota
	_
	python
	golang
	javascript
)

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	fmt.Println(b, kb, mb, gb, tb, pb)
}
