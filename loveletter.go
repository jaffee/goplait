package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var s string
		fmt.Scan(&s)
		fmt.Println(numToPalindrome(s))
	}
}

func numToPalindrome(s string) int {
	l := len(s)
	num := 0
	for i := 0; i < l/2; i++ {
		c1 := int(s[i])
		c2 := int(s[(l-i)-1])
		num += diff(c1, c2)
	}
	return num
}

func diff(n1, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	} else {
		return n2 - n1
	}
}
