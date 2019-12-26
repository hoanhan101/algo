/*
Problem:
- Design a stack that supports push, pop, top, and retrieving the minimum element
  in constant time.

Approach:
- We use two stack implementation themselves: one holds all the items and the
  other holds all the minimum values after each push() and pop().

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMinStack(t *testing.T) {
	s := newMinStack()

	s.push(4)
	common.Equal(t, 4, s.getMin())

	s.push(8)
	common.Equal(t, 4, s.getMin())

	s.push(6)
	common.Equal(t, 4, s.getMin())

	s.push(2)
	common.Equal(t, 2, s.getMin())

	common.Equal(t, 2, s.pop())
	common.Equal(t, 4, s.getMin())
	common.Equal(t, 6, s.pop())
	common.Equal(t, 4, s.getMin())
	common.Equal(t, 8, s.pop())
	common.Equal(t, 4, s.getMin())
	common.Equal(t, 4, s.pop())
}

type minStack struct {
	stack    *common.Stack
	minStack *common.Stack
}

func newMinStack() *minStack {
	s := common.NewStack()
	ms := common.NewStack()

	return &minStack{
		stack:    s,
		minStack: ms,
	}
}

func (s *minStack) push(i int) {
	// push the item on top of our stack.
	s.stack.Push(i)

	// only push the item to the minStack if it smaller  or equal to the last
	// item in minStack.
	if s.minStack.Size() == 0 || i <= s.minStack.Top().(int) {
		s.minStack.Push(i)
	}
}

func (s *minStack) pop() int {
	// pop the item of our stack.
	i := s.stack.Pop().(int)

	// if popped item is the smallest item, also remove that from minStack.
	if i == s.minStack.Top().(int) {
		s.minStack.Pop()
	}

	return i
}

func (s *minStack) getMin() int {
	return s.minStack.Top().(int)
}
