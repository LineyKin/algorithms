package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type bubbleSort struct {
	list []int
	want []int
}

func TestBubbleSort(t *testing.T) {
	testData := []bubbleSort{
		{[]int{5, 4, 6, 3, 7, 2, 8, 1, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{5, 4, 6, 3, 7, 2, 8, 1, 9, -10, -20}, []int{-20, -10, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{}, []int{}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.BubbleSort(v.list))
	}
}

func TestBubbleSortDesc(t *testing.T) {
	testData := []selectionSort{
		{[]int{5, 4, 6, 3, 7, 2, 8, 1, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{5, 4, 6, 3, 7, 2, 8, 1, 9, -10, -20}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, -10, -20}},
		{[]int{}, []int{}},
		{[]int{2, 1, 0}, []int{2, 1, 0}},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.BubbleSortDesc(v.list))
	}
}
