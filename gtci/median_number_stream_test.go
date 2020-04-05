/*
Problem:
- Find the median of a number stream.

Example:
1. insert(1)
2. findMedian() -> 1
3. insert(6)
2. findMedian() -> 3.5
3. insert(2)
5. findMedian() -> 2
6. insert(5)
7. findMedian() -> 3.5
8. insert(3)
9. findMedian() -> 3

Approach:
- Divide the stream into 2 lists where one holds all numbers that are less
  than the median and vice versa.
- Since the median is either be the largest of the smaller list or the smallest
  one of the larger list, can use a min heap and max heap respectively.
- At every insertion, balance the numbers in both heap.

Cost:
- O(1) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMedianInStream(t *testing.T) {
	s := newMStream()
	s.insert(1)
	common.Equal(t, float64(1), s.findMedian())
	s.insert(6)
	common.Equal(t, float64(3.5), s.findMedian())
	s.insert(2)
	common.Equal(t, float64(2), s.findMedian())
	s.insert(5)
	common.Equal(t, float64(3.5), s.findMedian())
	s.insert(3)
	common.Equal(t, float64(3), s.findMedian())
	s.insert(4)
	common.Equal(t, float64(3.5), s.findMedian())
}

type mstream struct {
	minHeap *common.MinHeap
	maxHeap *common.MaxHeap
}

func newMStream() *mstream {
	return &mstream{
		minHeap: common.NewMinHeap(),
		maxHeap: common.NewMaxHeap(),
	}
}

func (s *mstream) insert(n int) {
	if s.maxHeap.Len() == 0 || s.maxHeap.Peek() >= n {
		s.maxHeap.Push(n)
	} else {
		s.minHeap.Push(n)
	}

	// reblancemaxHeap can have 1 more number than minHeap.
	if s.maxHeap.Len() > s.minHeap.Len()+1 {
		s.minHeap.Push(s.maxHeap.Pop())
	} else if s.maxHeap.Len() < s.minHeap.Len() {
		s.maxHeap.Push(s.minHeap.Pop())
	}
}

func (s *mstream) findMedian() float64 {
	if s.maxHeap.Len() == s.minHeap.Len() {
		return float64(s.maxHeap.Peek()+s.minHeap.Peek()) / 2.0
	}

	return float64(s.maxHeap.Peek())
}
