package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	var pm, pn, bm, bn int

	for m := 0; m < T; m++ {
		var row string
		fmt.Scan(&row)
		for n := 0; n < len(row); n++ {
			if row[n] == 'm' {
				bm = m
				bn = n
			} else if row[n] == 'p' {
				pm = m
				pn = n
			}
		}
	}

	for ; pm > bm; bm++ {
		fmt.Println("DOWN")
	}
	for ; bm > pm; bm-- {
		fmt.Println("UP")
	}
	for ; pn > bn; bn++ {
		fmt.Println("RIGHT")
	}
	for ; pn < bn; bn-- {
		fmt.Println("LEFT")
	}
}
