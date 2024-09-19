package main

import (
	"algorithms/alg"
	"fmt"
)

func main() {

	x := alg.RecursiveBinarySearch([]int{1, 2, 6, 7}, 5, 0, 4)

	fmt.Println(x)
}
