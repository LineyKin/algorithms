package main

import (
	"algorithms/alg"
	"fmt"
)

func main() {
	ar1 := [3]int{1, 2, 3}
	ar2 := [3]int{1, 2, 3}
	ar3 := [3]int{1, 3, 2}

	fmt.Println(ar1 == ar2)
	fmt.Println(ar1 == ar3)

	fmt.Println(alg.BubbleSort([]int{5, 4, 6, 3, 7, 2, 8, 1, 9}))
}
