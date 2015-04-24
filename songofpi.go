package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pi := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 3}
	var T int
	fmt.Scan(&T)
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < T; i++ {
		var b byte
		b, _ = r.ReadByte()
		count := 0
		index := 0
		pisong := true
		for b != 10 {
			for b != 32 && b != 10 {
				count = count + 1
				b, _ = r.ReadByte()
			}
			if count != pi[index] {
				pisong = false
			}
			count = 0
			index += 1
			if b == 10 {
				break
			}
			b, _ = r.ReadByte()
		}
		if pisong {
			fmt.Println("It's a pi song.")
		} else {
			fmt.Println("It's not a pi song.")
		}
	}

}
