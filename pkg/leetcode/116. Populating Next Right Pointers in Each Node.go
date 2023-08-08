package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	result := make([][]int, 0)
	queue := NodeQueue{}
	temQueue := NodeQueue{}
	queue.Enqueue(root)
	result = append(result, []int{root.Val})
	for queue.HasElement() || temQueue.HasElement() {
		if queue.HasElement() {
			item := queue.Dequeue()
			if item.Left != nil {
				temQueue.Enqueue(item.Left)
			}
			if item.Right != nil {
				temQueue.Enqueue(item.Right)
			}
			item.Next = queue.Front()
		} else {
			queue = temQueue
			temQueue = NodeQueue{}
		}
	}
	return root
}
