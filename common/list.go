package common

// ListNode represents a node in a linked list.
type ListNode struct {
	Value int
	Next  *ListNode
}

// LinkedListToSlice converts a linked list into an array of integer.
func LinkedListToSlice(node *ListNode) []int {
	out := []int{}

	for n := node; n != nil; n = n.Next {
		out = append(out, n.Value)
	}

	return out
}
