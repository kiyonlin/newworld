package main

import "fmt"

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	var ptr *int
	fmt.Printf("ptr 的值为：%x\n", ptr)
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = swap1(a, b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}
