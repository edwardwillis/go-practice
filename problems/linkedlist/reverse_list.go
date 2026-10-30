package linkedlist

// Reverse reverses a singly linked list iteratively and returns the new head.
// O(n) time, O(1) space.
func Reverse(head *Node) *Node {
	var prev *Node
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
