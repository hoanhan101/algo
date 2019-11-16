package common

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

// NewTreeNode returns a new TreeNode.
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{nil, v, nil}
}
