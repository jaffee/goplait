package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &N)
	la := 10000000
	lb := 10000000
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		if a < la {
			la = a
		}
		if b < lb {
			lb = b
		}
	}
	fmt.Println(la * lb)
}
