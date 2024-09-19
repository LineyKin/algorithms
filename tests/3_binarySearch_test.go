package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type binarySearch struct {
	list   []int
	target int
	want   int
}

func TestBinarySearch(t *testing.T) {
	testData := []binarySearch{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 17, 16},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 25, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, 17, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, -10, -1},
		{[]int{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, 3, -1},
		{[]int{7, 81, 92, 103, 114, 125, 136, 147, 158, 169, 270}, 92, 2},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.BinarySearch(v.list, v.target))
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	testData := []binarySearch{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 17, 16},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 25, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, 17, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, -10, -1},
		{[]int{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 27}, 3, -1},
		{[]int{7, 81, 92, 103, 114, 125, 136, 147, 158, 169, 270}, 92, 2},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.RecursiveBinarySearch(v.list, v.target, 0, len(v.list)))
	}
}
