package main

import (
	"fmt"
)

func main() {
	var L int
	var R int
	fmt.Scan(&L)
	fmt.Scan(&R)
	max := 0
	for a := L; a <= R; a++ {
		for b := a; b <= R; b++ {
			val := a ^ b
			if val > max {
				max = val
			}
		}
	}
	fmt.Println(max)
}
