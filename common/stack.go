package common

type Stack struct {
	Items []interface{}
}

func NewStack() *Stack {
	items := []interface{}{}
	return &Stack{Items: items}
}

// Push a new item onto the stack.
func (s *Stack) Push(i interface{}) {
	s.Items = append(s.Items, i)
}

// Pop remove and return the last item.
func (s *Stack) Pop() interface{} {
	if len(s.Items) == 0 {
		return nil
	}

	// get the last item.
	i := len(s.Items) - 1
	top := s.Items[i]

	// remove the last item.
	s.Items = s.Items[:i]

	return top
}

// Peek return the last item without removing it.
func (s *Stack) Peek() interface{} {
	if len(s.Items) == 0 {
		return nil
	}

	return s.Items[len(s.Items)-1]
}

func (s *Stack) Size() int {
	return len(s.Items)
}
