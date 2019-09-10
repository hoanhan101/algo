package common

// Stack represent a stack data structure.
type Stack struct {
	Items []int
}

// NewStack return a new Stack.
func NewStack() *Stack {
	items := []int{}
	return &Stack{Items: items}
}

// Push a new item onto the stack.
func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

// Pop remove and return the last item.
func (s *Stack) Pop() int {
	// return -1 if there is nothing in the stack.
	if len(s.Items) == 0 {
		return -1
	}

	// get the last item.
	i := len(s.Items) - 1
	top := s.Items[i]

	// remove the last item.
	s.Items = s.Items[:i]

	return top
}

// Peek return the last item without removing it.
func (s *Stack) Peek() int {
	// return -1 if there is nothing in the stack.
	if len(s.Items) == 0 {
		return -1
	}

	return s.Items[len(s.Items)-1]
}

// Size return the Stack's size.
func (s *Stack) Size() int {
	return len(s.Items)
}
