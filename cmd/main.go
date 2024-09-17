package main

import (
	"fmt"
)

func main() {
	ar1 := [3]int{1, 2, 3}
	ar2 := [3]int{1, 2, 3}
	ar3 := [3]int{1, 3, 2}

	fmt.Println(ar1 == ar2)
	fmt.Println(ar1 == ar3)
}
