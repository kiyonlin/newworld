package main

import "fmt"

func main() {
	cache := make(map[string]string)
	cache["name"] = "jack"
	fmt.Println(cache)

	var arr1 [5]int
	arr2 := [3]int{2, 4, 6}
	arr3 := [...]int{2, 4, 5, 6, 78, 0}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][2][1]int
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(&arr1)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
