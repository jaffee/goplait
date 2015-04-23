package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	//	fmt.Printf("%v, %v", num, err)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
		//		fmt.Printf("%v, %v", num, err)
	}
	for n > 0 {
		min := get_min(n, arr)
		var newarr []int
		for i, _ := range arr {
			newnum := arr[i] - min
			if newnum > 0 {
				newarr = append(newarr, newnum)
			}
		}
		n = len(newarr)
		arr = newarr
		fmt.Println(n)
	}

	//fmt.Printf("%v %v", n, arr)

	//	n! / (r! * (n-r)!)
}

func get_min(size int, arr []int) int {
	min := 999999999
	for i := 0; i < size; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
