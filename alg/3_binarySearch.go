package alg

// бинарный поиск, сложность O(log2(n))
func BinarySearch(numbers []int, target int) int {
	start := 0
	len := len(numbers)
	end := len
	var middle int
	for start <= end {
		middle = (start + end) / 2

		// для случаев, когда target больше любого элемента numbers
		if middle >= len {
			return -1
		}

		if numbers[middle] == target {
			return middle
		}
		if target < numbers[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return -1
}

// бинарный поиск, сложность O(log2(n))
// реализация с помощью рекурсии
func RecursiveBinarySearch(numbers []int, target, start, end int) int {

	if start > end {
		return -1
	}

	middle := (start + end) / 2

	// для случаев, когда target больше любого элемента numbers
	if middle >= len(numbers) {
		return -1
	}

	if target == numbers[middle] {
		return middle
	}

	if target < numbers[middle] {
		return RecursiveBinarySearch(numbers, target, start, middle-1)
	} else {
		return RecursiveBinarySearch(numbers, target, middle+1, end)
	}
}
