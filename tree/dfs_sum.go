package tree

func DfsSum(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Val + DfsSum(root.L) + DfsSum(root.R)
}
