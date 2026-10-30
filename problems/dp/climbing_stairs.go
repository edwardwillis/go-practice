// Package dp contains dynamic programming interview problems.
package dp

// ClimbStairs returns the number of distinct ways to climb n stairs taking
// 1 or 2 steps at a time. O(n) time, O(1) space.
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, cur := 1, 2
	for i := 3; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}
