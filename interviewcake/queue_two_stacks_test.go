/*
Problem:
- Implement a queue with 2 stacks.

Approach:
- Use one stack for enqueuing item and the other to reverse the order them for
  dequeuing/popping item.

Cost:
- O(1) time, O(m) space m is the number of operations.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestQueueStack(t *testing.T) {
	s := newQueueStack()

	// return -1 if there is nothing in the stack.
	common.Equal(t, -1, s.dequeue())

	// define test cases.
	tests := []int{1, 2, 3, 4, 5, 6}

	// enqueue items to the queue one by one.
	for _, tt := range tests {
		s.enqueue(tt)
	}

	// dequeue items from the queue one by one and check against it.
	for _, tt := range tests {
		common.Equal(t, tt, s.dequeue())
	}

	// return -1 if there is nothing in the stack.
	common.Equal(t, -1, s.dequeue())
}

type queueStack struct {
	stack1 *common.Stack
	stack2 *common.Stack
}

func newQueueStack() *queueStack {
	s1 := common.NewStack()
	s2 := common.NewStack()

	return &queueStack{
		stack1: s1,
		stack2: s2,
	}
}

func (q *queueStack) enqueue(i int) {
	q.stack1.Push(i)
}

func (q *queueStack) dequeue() int {
	if q.stack2.Size() == 0 {
		// reverse the order of all items in other stack.
		for q.stack1.Size() > 0 {
			q.stack2.Push(q.stack1.Pop())
		}

		// return -1 if the queue is empty
		if q.stack2.Size() == 0 {
			return -1
		}
	}

	return q.stack2.Pop().(int)
}
