package alg

// быстрая сортировка по возрастанию, сложность O(n*log2(n))
func QuikSort(numbers []int) []int {
	len := len(numbers)

	if len <= 1 {
		return numbers
	}

	// опорный элемент и его индекс (просто число из середины слайса)
	pivotIndex := len / 2
	pivot := numbers[pivotIndex]
	less := make([]int, 0, pivotIndex+1)
	greater := make([]int, 0, pivotIndex+1)

	for i := 0; i < len; i++ {
		if i == pivotIndex {
			continue
		}

		if numbers[i] < pivot {
			less = append(less, numbers[i])
		} else {
			greater = append(greater, numbers[i])
		}
	}

	result := append(QuikSort(less), pivot)
	result = append(result, QuikSort(greater)...)

	return result
}

// быстрая сортировка по убыванию, сложность O(n*log2(n))
func QuikSortDesc(numbers []int) []int {
	len := len(numbers)

	if len <= 1 {
		return numbers
	}

	// опорный элемент и его индекс (просто число из середины слайса)
	pivotIndex := len / 2
	pivot := numbers[pivotIndex]
	less := make([]int, 0, pivotIndex+1)
	greater := make([]int, 0, pivotIndex+1)

	for i := 0; i < len; i++ {
		if i == pivotIndex {
			continue
		}

		if numbers[i] < pivot {
			less = append(less, numbers[i])
		} else {
			greater = append(greater, numbers[i])
		}
	}

	result := append(QuikSortDesc(greater), pivot)
	result = append(result, QuikSortDesc(less)...)

	return result
}
