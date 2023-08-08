package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type TreeNodeQueue struct {
	items []*TreeNode
}

func (q *TreeNodeQueue) Enqueue(item *TreeNode) {
	q.items = append(q.items, item)
}

func (q *TreeNodeQueue) Dequeue() *TreeNode {
	if len(q.items) == 0 {
		return nil // 队列为空，返回特定值表示错误
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *TreeNodeQueue) HasElement() bool {
	return len(q.items) > 0
}

type NodeQueue struct {
	items []*Node
}

func (q *NodeQueue) Enqueue(item *Node) {
	q.items = append(q.items, item)
}

func (q *NodeQueue) Dequeue() *Node {
	if len(q.items) == 0 {
		return nil // 队列为空，返回特定值表示错误
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *NodeQueue) Front() *Node {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *NodeQueue) HasElement() bool {
	return len(q.items) > 0
}
