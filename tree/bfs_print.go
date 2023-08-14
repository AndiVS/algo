package tree

func BFS(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return []int{}
	}

	queue := Queue[*TreeNode]{}
	queue.Push(root)
	for queue.Len() > 0 {
		node := queue.Pop()

		ans = append(ans, node.Val)

		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
	}

	return ans
}

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Len() int {
	return len(q.queue)
}

func (q *Queue[T]) Push(elem T) {
	q.queue = append(q.queue, elem)
}

func (q *Queue[T]) Pop() T {
	defer func() {
		q.queue = q.queue[1:]
	}()
	return q.queue[0]
}
