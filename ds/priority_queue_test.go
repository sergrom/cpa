package ds

import (
	"container/heap"
	"testing"
)

func TestPriorityQueueUpdateAfterConstruction(t *testing.T) {
	for _, priorities := range [][]int{{7, 6, 5, 4, 3, 2, 1}, {1, 4, 2, 7, 3, 6, 5}} {
		items := make([]*PqItem, len(priorities))
		for i, p := range priorities {
			items[i] = &PqItem{Priority: p}
		}
		item := items[1]
		pq := NewPriorityQueue(items...)
		pq.Update(item, "updated", 0)
		previous := 8
		for pq.Len() > 0 {
			got := heap.Pop(pq).(*PqItem)
			if got.Priority > previous {
				t.Fatalf("priority %d follows %d", got.Priority, previous)
			}
			previous = got.Priority
		}
		if item.Value != "updated" {
			t.Fatal("value was not updated")
		}
	}
}

func TestNewPriorityQueueEmpty(t *testing.T) {
	pq := NewPriorityQueue()
	item := &PqItem{Value: "first", Priority: 1}
	heap.Push(pq, item)
	if heap.Pop(pq) != item || pq.Len() != 0 {
		t.Fatal("incorrect push/pop on initially empty queue")
	}
}
