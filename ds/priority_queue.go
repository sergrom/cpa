package ds

import "container/heap"

type PqItem struct {
	index    int
	Value    string
	Priority int
}

type PriorityQueue []*PqItem

// NewPriorityQueue initializes a max-priority queue with the given items.
func NewPriorityQueue(items ...*PqItem) *PriorityQueue {
	pq := PriorityQueue(append([]*PqItem(nil), items...))
	for i, item := range pq {
		item.index = i
	}
	heap.Init(&pq)
	return &pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *PqItem, value string, priority int) {
	item.Value, item.Priority = value, priority
	heap.Fix(pq, item.index)
}

// Usage:
// pq := NewPriorityQueue(
// 	&PqItem{Value: "banana", Priority: 3},
// 	&PqItem{Value: "apple", Priority: 2},
// 	&PqItem{Value: "pear", Priority: 4},
// )

// item1 := heap.Pop(pq).(*PqItem) // pear (Priority=4)
// item2 := heap.Pop(pq).(*PqItem) // banana (Priority=3)
// item3 := heap.Pop(pq).(*PqItem) // apple (Priority=2)
