package alg

// пузырьковая сортировка по возрастанию, сложность О(n^2)
func BubbleSort(numbers []int) []int {
	len := len(numbers)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if numbers[j+1] < numbers[j] {
				tmp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = tmp
			}
		}
	}

	return numbers
}

// пузырьковая сортировка по убыванию, сложность О(n^2)
func BubbleSortDesc(numbers []int) []int {
	len := len(numbers)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if numbers[j+1] > numbers[j] {
				tmp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = tmp
			}
		}
	}

	return numbers
}
