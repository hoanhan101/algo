package common

// ListNode represents a node in a linked list.
type ListNode struct {
	Value int
	Next  *ListNode
}

// NewListNode returns a new list node.
func NewListNode(v int) *ListNode {
	return &ListNode{v, nil}
}

// AddNext adds a next node to the end of list.
func (l *ListNode) AddNext(v int) {
	for n := l; n != nil; n = n.Next {
		if n.Next == nil {
			new := NewListNode(v)
			n.Next = new
			break
		}
	}
}

// LinkedListToSlice converts a linked list into an array of integer.
func LinkedListToSlice(node *ListNode) []int {
	out := []int{}

	for n := node; n != nil; n = n.Next {
		out = append(out, n.Value)
	}

	return out
}
