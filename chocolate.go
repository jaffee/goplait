package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N, C, M int
		fmt.Scan(&N, &C, &M)
		nc := N / C
		nw := nc
		for nw >= M {
			new_chocs := nw / M
			left_wrappers := nw % M
			nw = left_wrappers + new_chocs
			nc += new_chocs
		}
		fmt.Println(nc)
	}
}
