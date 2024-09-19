package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type recursion struct {
	n    int
	want int
}

func TestFactorial(t *testing.T) {
	testData := []recursion{
		{5, 120},
		{6, 720},
		{0, 1},
		{-4, 0},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.Factorial(v.n))
	}
}

func TestFibonacci(t *testing.T) {
	testData := []recursion{
		{0, 0},
		{2, 1},
		{7, 13},
		{-4, 0},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, alg.Fibonacci(v.n))
	}
}
