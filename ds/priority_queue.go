package ds

import "container/heap"

type PqItem struct {
	index    int
	value    string
	priority int
}

type PriorityQueue []*PqItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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
	item.value, item.priority = value, priority
	heap.Fix(pq, item.index)
}

// Ussage:
// pq := PriorityQueue{
// 	&PqItem{index: 0, value: "banana", priority: 3},
// 	&PqItem{index: 1, value: "apple", priority: 2},
// 	&PqItem{index: 2, value: "pear", priority: 4},
// }
// heap.Init(&pq)

// item1 := heap.Pop(&pq).(*PqItem) // pear (priority=4)
// item2 := heap.Pop(&pq).(*PqItem) // banana (priority=3)
// item3 := heap.Pop(&pq).(*PqItem) // apple (priority=2)
