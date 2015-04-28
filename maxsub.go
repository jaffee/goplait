package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		nums := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&nums[j])
		}
		printMaxes(nums)
	}
}

func printMaxes(nums []int) {
	maxcon := nums[0]
	curcon := nums[0]
	maxnoncon := nums[0]
	for _, n := range nums[1:] {
		if n >= 0 {
			if maxcon < 0 {
				maxcon = n
			}
			curcon += n
		} else { //n < 0
			if maxcon < n {
				maxcon = n
			} else {
				curcon += n
			}
		}
		if curcon > maxcon && curcon > 0 {
			maxcon = curcon
		}
		if curcon < 0 {
			curcon = 0
		}

		if maxnoncon < 0 {
			if n > maxnoncon {
				maxnoncon = n
			}
		} else if n > 0 {
			maxnoncon += n
		}
	}

	fmt.Println(maxcon, maxnoncon)
}
