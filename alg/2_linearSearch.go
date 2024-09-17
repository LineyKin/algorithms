package alg

func LinearSearch(numbers []int, target int) int {
	for k, v := range numbers {
		if v == target {
			return k
		}
	}

	return -1
}
