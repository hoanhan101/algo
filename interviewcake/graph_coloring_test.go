/*
Problem:
- Given an undirected graph, with maximum degree d, find a graph coloring
  using at most d + 1 colors. Assume that there is no node with a loop.

Example:
- For a graph with degree 3, we could use at most 4 colors.

Approach:
- Use a greedy approach to iterate over the graph and assign each node the
  first non-taken color that we found.

Solution:
- Iterate through the node in the graph and get the current node's neighbors
  colors to keep track of all colors that are already taken.
- Iterate through the list of colors and assign the first color in the given
  colors list that is not already taken.

Cost:
- O(m) time and O(d) space, where m is sum of all the nodes and edges, d is
  the number of all colors.
- Even though it seems like we have an outer and inner loop, we carefully walk
  through the graph one node at a time and stop checking for colors as soon as
  we found one that is non-taken.
- About the space complexity, it makes sense that we have to store all
  available colors in the worst case. Hence it takes up O(d) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

// GraphNode holds a graph node and its neighbors.
type GraphNode struct {
	label     string
	color     string
	neighbors []*GraphNode
}

func TestColorGraph(t *testing.T) {
	// define test cases.
	// case 1.
	a1 := &GraphNode{label: "a"}
	b1 := &GraphNode{label: "b"}
	c1 := &GraphNode{label: "c"}

	a1.neighbors = []*GraphNode{b1, c1}
	b1.neighbors = []*GraphNode{a1, c1}
	c1.neighbors = []*GraphNode{b1, a1}

	graph1 := []*GraphNode{a1, b1, c1}

	// case 2.
	a2 := &GraphNode{label: "a"}
	b2 := &GraphNode{label: "b"}
	c2 := &GraphNode{label: "c"}

	a2.neighbors = []*GraphNode{b2}
	b2.neighbors = []*GraphNode{a2, c2}
	c2.neighbors = []*GraphNode{b2}

	graph2 := []*GraphNode{a2, b2, c2}

	// case 3.
	a3 := &GraphNode{label: "a"}
	b3 := &GraphNode{label: "b"}
	c3 := &GraphNode{label: "c"}
	d3 := &GraphNode{label: "d"}

	a3.neighbors = []*GraphNode{b3}
	b3.neighbors = []*GraphNode{a3, c3, d3}
	c3.neighbors = []*GraphNode{b3}
	d3.neighbors = []*GraphNode{b3}

	graph3 := []*GraphNode{a3, b3, c3, d3}

	// case 4.
	a4 := &GraphNode{label: "a"}
	b4 := &GraphNode{label: "b"}
	c4 := &GraphNode{label: "c"}
	d4 := &GraphNode{label: "d"}

	a4.neighbors = []*GraphNode{b4, d4}
	b4.neighbors = []*GraphNode{a4, c4, d4}
	c4.neighbors = []*GraphNode{b4}
	d4.neighbors = []*GraphNode{b4, a4}

	graph4 := []*GraphNode{a4, b4, c4, d4}

	// case 5.
	a5 := &GraphNode{label: "a"}
	b5 := &GraphNode{label: "b"}
	c5 := &GraphNode{label: "c"}
	d5 := &GraphNode{label: "d"}

	a5.neighbors = []*GraphNode{b5, c5, d5}
	b5.neighbors = []*GraphNode{a5, c5, d5}
	c5.neighbors = []*GraphNode{a5, b5}
	d5.neighbors = []*GraphNode{a5, b5}

	graph5 := []*GraphNode{a5, b5, c5, d5}

	// case 6.
	a6 := &GraphNode{label: "a"}
	b6 := &GraphNode{label: "b"}
	c6 := &GraphNode{label: "c"}
	d6 := &GraphNode{label: "d"}

	a6.neighbors = []*GraphNode{b6, c6, d6}
	b6.neighbors = []*GraphNode{a6, c6, d6}
	c6.neighbors = []*GraphNode{a6, b6, d6}
	d6.neighbors = []*GraphNode{a6, b6, c6}

	graph6 := []*GraphNode{a6, b6, c6, d6}

	// define tests outputs.
	tests := []struct {
		in1      []*GraphNode
		in2      []string
		expected []string
	}{
		{graph1, []string{"black", "white", "yellow"}, []string{"black", "white", "yellow"}},
		{graph2, []string{"black", "white", "yellow"}, []string{"black", "white", "black"}},
		{graph3, []string{"black", "white", "yellow", "green"}, []string{"black", "white", "black", "black"}},
		{graph4, []string{"black", "white", "yellow", "green"}, []string{"black", "white", "black", "yellow"}},
		{graph5, []string{"black", "white", "yellow", "green"}, []string{"black", "white", "yellow", "yellow"}},
		{graph6, []string{"black", "white", "yellow", "green"}, []string{"black", "white", "yellow", "green"}},
	}

	for _, tt := range tests {
		colorGraph(tt.in1, tt.in2)

		result := []string{}
		for _, g := range tt.in1 {
			result = append(result, g.color)
		}

		common.Equal(t, tt.expected, result)
	}
}

func colorGraph(graph []*GraphNode, colors []string) {
	for _, node := range graph {
		// get the node's neighbors colors to keep track of all colors that are
		// already taken.
		taken := []string{}
		for _, neighbor := range node.neighbors {
			// only append the color that we haven't seen yet.
			if !common.ContainString(taken, neighbor.color) {
				taken = append(taken, neighbor.color)
			}
		}

		// assign the first color in the given colors list that is not already
		// taken.
		for _, color := range colors {
			if !common.ContainString(taken, color) {
				node.color = color
				break
			}
		}
	}
}
