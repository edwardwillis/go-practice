// Package strings contains string manipulation interview problems.
package strings

// IsAnagram reports whether t is an anagram of s. O(n) time, O(1) space
// (bounded alphabet), assumes ASCII lowercase letters.
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
