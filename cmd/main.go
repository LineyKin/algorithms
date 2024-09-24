package main

import (
	"algorithms/alg"
	"fmt"
)

// 5+10+11+7+5+1+5+10+15
func main() {
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

	t2 := alg.Tree{{Val: 15, Children: nil}}

	fmt.Println(tree)
	fmt.Println(t2.RecursiveTreeSum())

}
