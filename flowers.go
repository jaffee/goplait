package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var N, K int
	fmt.Fscan(r, &N)
	fmt.Fscan(r, &K)
	cis := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &cis[i])
	}

	sort.Ints(cis)
	multiplier := 1
	cost := 0
	k := 1
	for i := len(cis) - 1; i >= 0; i-- {
		cost += multiplier * cis[i]

		if k == K {
			k = 1
			multiplier++
		} else {
			k += 1
		}
	}

	fmt.Println(cost)
}
