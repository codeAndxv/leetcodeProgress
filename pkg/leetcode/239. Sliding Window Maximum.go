package leetcode

/*func MaxSlidingWindow(nums []int, k int) []int {
	highL := make([]int, len(nums))
	highL[0] = 0
	maxList := make([]int, len(nums)-k+1)
	for i := 1; i < len(nums); i++ {
		highL[i] = i
		tem := i - 1
		for tem >= i-k+1 {
			if nums[tem] > nums[i] {
				highL[i] = tem
				break
			}
			if tem == highL[tem] {
				break
			}
			tem = highL[tem]
		}
	}
	for i := k - 1; i < len(nums); i++ {
		maxList[i-k+1] = nums[highL[i]]
	}
	return maxList
}*/

/*func maxSlidingWindow(nums []int, k int) []int {
	maxList := make([]int, len(nums)-k+1)
	maxHeap := NewMaxHeap()
	for i := 0; i < len(nums); i++ {
		if i > k-1 {
			maxHeap.Pop()
			maxHeap.Push(nums[i])
			maxList[i-k+1] = maxHeap.
		}else {
			maxHeap.Push(nums[i])
		}
	}
	return maxList
}

// 最大堆类型
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 使用Greater函数
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 队列类型
type MaxHeapQueue struct {
	queue   []int
	maxHeap MaxHeap
}

func NewMaxHeapQueue() *MaxHeapQueue {
	return &MaxHeapQueue{
		queue:   []int{},
		maxHeap: MaxHeap{},
	}
}

func (q *MaxHeapQueue) Enqueue(val int) {
	q.queue = append(q.queue, val)
	q.maxHeap.Push(val)
}

func (q *MaxHeapQueue) Dequeue() (int, bool) {
	if len(q.queue) == 0 {
		return 0, false
	}

	val := q.queue[0]
	q.queue = q.queue[1:]
	q.maxHeap.Pop()
	return val, true
}

func (q *MaxHeapQueue) GetMax() (int, bool) {
	if len(q.maxHeap) == 0 {
		return 0, false
	}
	return q.maxHeap[0], true
}
*/
