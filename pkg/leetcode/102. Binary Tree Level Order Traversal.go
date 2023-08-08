package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := TreeNodeQueue{}
	temQueue := TreeNodeQueue{}
	queue.Enqueue(root)
	result = append(result, []int{root.Val})
	temSlice := make([]int, 0)
	for queue.HasElement() || temQueue.HasElement() {
		if queue.HasElement() {
			item := queue.Dequeue()
			if item.Left != nil {
				temQueue.Enqueue(item.Left)
				temSlice = append(temSlice, item.Left.Val)
			}
			if item.Right != nil {
				temQueue.Enqueue(item.Right)
				temSlice = append(temSlice, item.Right.Val)
			}
		} else {
			queue = temQueue
			temQueue = TreeNodeQueue{}
			result = append(result, temSlice)
			temSlice = make([]int, 0)
		}
	}
	return result
}
