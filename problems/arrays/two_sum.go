// Package arrays contains array/slice interview problems.
package arrays

// TwoSum returns the indices of the two numbers in nums that add up to target.
// Assumes exactly one solution exists. O(n) time, O(n) space.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
