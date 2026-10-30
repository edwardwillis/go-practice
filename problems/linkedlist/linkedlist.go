// Package linkedlist contains linked-list interview problems and the shared
// node type / helpers used across them.
package linkedlist

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// FromSlice builds a linked list from vals and returns its head.
func FromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ToSlice flattens a linked list starting at head into a slice.
func ToSlice(head *Node) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}
