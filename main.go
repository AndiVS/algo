package main

import (
	"fmt"
	"github.com/AndiVS/algo/tree"
)

func main() {
	fmt.Print(tree.TargetSum(&tree.TreeNode{
		Val: 1,
		Right: &tree.TreeNode{
			Val: 2,
			Right: &tree.TreeNode{
				Val:   3,
				Right: nil,
				Left:  nil,
			},
			Left: &tree.TreeNode{
				Val:   3,
				Right: nil,
				Left:  nil,
			},
		},
		Left: nil,
	},
		0, 6))
}
