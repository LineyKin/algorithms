package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type linearSearch struct {
	list   []int
	target int
	want   int
}

func TestLinearSearch(t *testing.T) {
	testData := []linearSearch{
		{[]int{325, 45745, 4432, 47554, 23, 1}, 45745, 1},
		{[]int{325, 45745, 4432, 47554, 23, 1}, 1, 5},
		{[]int{325, 45745, 4432, 47554, 23, 1}, 10, -1},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.LinearSearch(v.list, v.target))
	}
}
