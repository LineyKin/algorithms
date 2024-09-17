package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type maxNumber struct {
	list []int
	want int
}

func TestMaxNumber(t *testing.T) {
	testData := []maxNumber{
		{[]int{12, 43, 67, 89, 234}, 234},
		{[]int{-12, -43, -67, -89, -234, -1}, -1},
		{[]int{}, 0},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.MaxNumber(v.list))
	}
}
