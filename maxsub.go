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
		fmt.Scanf("%v", &nums)
		printMaxes(nums)
	}
}

func printMaxes(nums []int) {
	maxcon := nums[0]
	curcon := nums[0]
	maxnoncon := nums[0]
	fmt.Println(nums)
	for _, n := range nums[1:] {
		if n >= 0 {
			if maxcon < 0 {
				maxcon = n
				curcon += n
			} else {
				curcon += n
			}
		} else { //n < 0
			if maxcon < n {
				maxcon = n
				curcon = n
			} else {
				curcon = 0
			}
		}
		if curcon > maxcon {
			maxcon = curcon
		}

		if maxnoncon < 0 {
			if n < 0 && n > maxnoncon {
				maxnoncon = n
			} else {
				maxnoncon = n
			}
		} else {
			maxnoncon += n
		}
	}

	fmt.Println(maxcon, maxnoncon)
}
