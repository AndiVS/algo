package tree

func DfsSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + DfsSum(root.Left) + DfsSum(root.Right)
}
