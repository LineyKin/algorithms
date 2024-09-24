package test

import (
	"algorithms/alg"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tree struct {
	tree alg.Tree
	want int
}

func TestTreeSum(t *testing.T) {
	tree := alg.Tree{
		{
			Val: 5,
			Children: []alg.TreeNode{
				{
					Val: 10,
					Children: []alg.TreeNode{
						{Val: 11, Children: nil},
					},
				},
				{
					Val: 7,
					Children: []alg.TreeNode{
						{
							Val: 5, Children: []alg.TreeNode{
								{Val: 1, Children: nil},
							},
						},
					},
				},
			},
		},
		{
			Val: 5,
			Children: []alg.TreeNode{
				{Val: 10, Children: nil},
				{Val: 15, Children: nil},
			},
		},
	}

	testData := []Tree{
		{tree, 69},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, v.tree.TreeSum())
		assert.Equal(t, v.want, v.tree.RecursiveTreeSum())
	}
}

func TestRecursiveTreeSum(t *testing.T) {
	tree := alg.Tree{
		{
			Val: 5,
			Children: []alg.TreeNode{
				{
					Val: 10,
					Children: []alg.TreeNode{
						{Val: 11, Children: nil},
					},
				},
				{
					Val: 7,
					Children: []alg.TreeNode{
						{
							Val: 5, Children: []alg.TreeNode{
								{Val: 1, Children: nil},
							},
						},
					},
				},
			},
		},
		{
			Val: 5,
			Children: []alg.TreeNode{
				{Val: 10, Children: nil},
				{Val: 15, Children: nil},
			},
		},
	}

	testData := []Tree{
		{tree, 69},
	}

	for _, v := range testData {
		assert.Equal(t, v.want, v.tree.RecursiveTreeSum())
	}
}
