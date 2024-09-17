package main

import (
	"algorithms/alg"
	"fmt"
)

func main() {
	numbers := []int{5, 9, 123, 457, 1654, 34654, 464, 2, 57, -123, 90}
	maxNumber := alg.MaxNumber(numbers)
	fmt.Println(maxNumber)
}
