package trees

// MaxDepth returns the length of the longest root-to-leaf path. O(n) time,
// O(h) space for the recursion stack (h = tree height).
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
