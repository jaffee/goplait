package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pi = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 3}

func main() {
	var T int
	fmt.Scan(&T)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < T; i++ {
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
		strlist := strings.Split(line, " ")
		var j int
		for j = 0; j < len(strlist); j++ {
			if len(strlist[j]) != pi[j] {
				break
			}
		}
		if j < len(strlist) {
			fmt.Println("It's not a pi song.")
		} else {
			fmt.Println("It's a pi song.")
		}
	}
}
