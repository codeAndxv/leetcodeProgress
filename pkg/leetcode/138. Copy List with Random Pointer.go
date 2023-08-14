package leetcode

type NewNode struct {
	Val    int
	Next   *NewNode
	Random *NewNode
}

func copyRandomList(head *NewNode) *NewNode {
	if head == nil {
		return head
	}
	nodes := make([]*NewNode, 0)
	nodeIndexMap := make(map[*NewNode]int)
	index := 0
	point1 := head
	for point1 != nil {
		nodeIndexMap[point1] = index
		nodes = append(nodes, point1)
		index++
		point1 = point1.Next
	}
	newNodes := make([]*NewNode, len(nodes))
	for i, v := range nodes {
		newNodes[i] = &NewNode{
			Val:    v.Val,
			Next:   nil,
			Random: nil,
		}
	}
	for i, v := range nodes {
		if i != len(nodes)-1 {
			newNodes[i].Next = newNodes[i+1]
		}
		if v.Random != nil {
			temIndex := nodeIndexMap[v.Random]
			newNodes[i].Random = newNodes[temIndex]
		}
	}
	return newNodes[0]
}
