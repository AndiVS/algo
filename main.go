package main

import (
	"fmt"
	"github.com/AndiVS/algo/tree"
)

func main() {
	fmt.Print(tree.DfsSum(&tree.Node{
		Val: 1,
		R: &tree.Node{
			Val: 2,
			R: &tree.Node{
				Val: 3,
				R:   nil,
				L:   nil,
			},
			L: &tree.Node{
				Val: 3,
				R:   nil,
				L:   nil,
			},
		},
		L: nil,
	}))
}
