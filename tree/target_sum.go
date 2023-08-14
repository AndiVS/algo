package tree

func TargetSum(root *Node, curSum int, targetSum int) bool {
	if root == nil {
		return false
	}

	curSum += root.Val
	if root.L == nil && root.R == nil {
		if curSum == targetSum {
			return true
		}
	}

	return TargetSum(root.L, curSum, targetSum) || TargetSum(root.R, curSum, targetSum)
}
