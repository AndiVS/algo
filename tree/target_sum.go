package tree

func TargetSum(root *TreeNode, curSum int, targetSum int) bool {
	if root == nil {
		return false
	}

	curSum += root.Val
	if root.Left == nil && root.Right == nil {
		if curSum == targetSum {
			return true
		}
	}

	return TargetSum(root.Left, curSum, targetSum) || TargetSum(root.Right, curSum, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			return true
		}
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
