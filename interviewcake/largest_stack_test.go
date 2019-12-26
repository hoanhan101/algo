/*
Problem:
- Implement a stack with a method getMax() that returns the largest element in
  the stack in O(1) time.

Approach:
- We use two stack implementation themselves: one holds all the items and the
  other holds all the maximum values after each push() and pop().
- That way, we could keep track of your maximum value up to date in constant
  time.

Cost:
- O(1) time, O(m) space where m is the number of operations performed on the
  stack.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMaxStack(t *testing.T) {
	s := newMaxStack()

	s.push(2)
	common.Equal(t, 2, s.getMax())

	s.push(6)
	common.Equal(t, 6, s.getMax())

	s.push(8)
	common.Equal(t, 8, s.getMax())

	s.push(4)
	common.Equal(t, 8, s.getMax())

	common.Equal(t, 4, s.pop())
	common.Equal(t, 8, s.getMax())
	common.Equal(t, 8, s.pop())
	common.Equal(t, 6, s.getMax())
	common.Equal(t, 6, s.pop())
	common.Equal(t, 2, s.getMax())
	common.Equal(t, 2, s.pop())
}

type maxStack struct {
	stack    *common.Stack
	maxStack *common.Stack
}

func newMaxStack() *maxStack {
	s := common.NewStack()
	ms := common.NewStack()

	return &maxStack{
		stack:    s,
		maxStack: ms,
	}
}

func (s *maxStack) push(i int) {
	// push the item on top of our stack.
	s.stack.Push(i)

	// only push the item to the maxStack if it greater or equal to the last
	// item in maxStack.
	if s.maxStack.Size() == 0 || i >= s.maxStack.Top().(int) {
		s.maxStack.Push(i)
	}
}

func (s *maxStack) pop() int {
	// pop the item of our stack.
	i := s.stack.Pop().(int)

	// if popped item is the largest item, also remove that from maxStack.
	if i == s.maxStack.Top().(int) {
		s.maxStack.Pop()
	}

	return i
}

func (s *maxStack) getMax() int {
	return s.maxStack.Top().(int)
}
