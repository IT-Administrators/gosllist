package gosllist

import (
	"fmt"
)

// Listnode. Must contain a reference to parent.
// Contains nil pointer reference when created.
type node struct {
	data any
	next *node
}

type ListWriter interface {
	AddNode(any) int
	RemoveNode(any)
	RemoveNodeOnPos(int)
	InsertNode(any, int)
}

type ListReader interface {
	Traverse()
	Size() int
}

// Create new node with specified data and reference to parent.
func NewNode(nodeValue any) *node {
	return &node{
		data: nodeValue,
		next: nil,
	}
}

// Get list size.
func (n *node) Size() int {
	if n.data == nil && n.next == nil {
		// fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for n != nil {
		i++
		n = n.next
	}
	return i
}

// Print list to stdout.
//
// The list printed from head (first node) to tail (last node).
// It can only be traversed in this direction.
func (n *node) Traverse() {
	if n == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for n != nil {
		fmt.Printf("%v -> ", n.data)
		n = n.next
	}
	fmt.Println()
}

// Add node to list. All nodes are appended at tail.
//
// Any type can be appended, the node type is build inside the function.
// So appending AddNode("Test") will append the node &{"Test",nil}.
func (n *node) AddNode(v any) int {
	// Check if current list is empty.
	// If empty append node.
	if n == nil {
		t := &node{v, nil}
		n = t
		return 0
	}

	// Iterate through list to check following cases:
	// - If current node data equals value, node is not appended as it already exists.
	// - If current node has no following element, append value as node to list.
	// Than advance to next node.
	for n != nil {
		if n.data == v {
			// fmt.Printf("node already exists %v\n", v)
			return -1
		}
		if n.next == nil {
			n.next = &node{v, nil}
			return -2
		}
		n = n.next
	}
	return -3
}

// Remove node with specified value from list.
//
// The specified value to be removed has to match the data property of the node.
func (n *node) RemoveNode(a any) {
	for n.next != nil {
		if n.next.data == a {
			n.next = n.next.next
		}
		n = n.next
	}
}

// Remove node on specified position.
//
// Position count starts by 1.
func (n *node) RemoveNodeOnPos(pos int) {
	counter := 1
	for n != nil {
		if counter == pos {
			n.next = n.next.next
		}
		n = n.next
		counter += 1
	}
}

// Insert node after specified position.
//
// The position is counted from 1.
func (n *node) InsertNode(a any, pos int) {
	// Start counting by 1.
	if pos <= 0 {
		pos = 1
	}
	counter := 1
	for n != nil {
		if counter == pos {
			// Save next node.
			temp := n.next
			// Append new node to current node.
			n.next = &node{a, nil}
			// Append former next node to new node.
			n.next.next = temp
		}
		// Advance.
		n = n.next
		counter += 1
	}
}
