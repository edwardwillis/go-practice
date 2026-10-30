// Package trees contains binary tree interview problems and the shared
// node type used across them.
package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
