package alg

// сортировка вставкой по возрастанию, сложность O(n^2)
func InsertionSort(numbers []int) []int {
	len := len(numbers)
	for i := 1; i < len; i++ {
		for j := i; j > -1; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}

	return numbers
}

// сортировка вставкой по убыванию, сложность O(n^2)
func InsertionSortDesc(numbers []int) []int {
	len := len(numbers)
	for i := 1; i < len; i++ {
		for j := i; j > -1; j-- {
			if numbers[j] > numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}

	return numbers
}
