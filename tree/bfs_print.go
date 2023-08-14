package tree

func BFS(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		ans = append(ans, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return ans
}
