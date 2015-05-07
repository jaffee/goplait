package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &T)

	for i := 0; i < T; i++ {
		var px, py, qx, qy int
		fmt.Fscan(r, &px, &py, &qx, &qy)
		outx := qx + (qx - px)
		outy := qy + (qy - py)
		fmt.Printf("%v %v\n", outx, outy)
	}

}
