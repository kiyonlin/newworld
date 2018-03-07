package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[rune]rune)

	for _, ch := range []rune(s) {
		lastOccured[ch] = ch
	}
	return len(lastOccured)
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(
			"黑化肥挥发灰会花飞黑化肥挥发灰会花飞"))
}
