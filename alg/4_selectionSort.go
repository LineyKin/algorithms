package alg

// сортировка выбором по возрастанию, сложность O(n^2)
func SelectionSort(numbers []int) []int {
	len := len(numbers)
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}

	return numbers
}

// сортировка выбором по убыванию, сложность O(n^2)
func SelectionSortDesc(numbers []int) []int {
	len := len(numbers)
	for i := 0; i < len; i++ {
		maxIndex := i
		for j := i + 1; j < len; j++ {
			if numbers[j] > numbers[maxIndex] {
				maxIndex = j
			}
		}

		numbers[i], numbers[maxIndex] = numbers[maxIndex], numbers[i]
	}

	return numbers
}
