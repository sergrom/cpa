package alg

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head
	for curr != nil {
		prev, curr.Next = curr.Next, prev
		prev, curr = curr, prev
	}
	return prev
}
