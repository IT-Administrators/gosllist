package gosllist

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Listnode. Must contain a reference to parent.
// Contains nil pointer reference when created.
type node struct {
	Data any
	Next *node
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

// Create new node with specified Data and reference to parent.
func NewNode(nodeValue any) *node {
	return &node{
		Data: nodeValue,
		Next: nil,
	}
}

// Get list size.
func (n *node) Size() int {
	if n.Data == nil && n.Next == nil {
		// fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for n != nil {
		i++
		n = n.Next
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
		fmt.Printf("%v -> ", n.Data)
		n = n.Next
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
	// - If current node Data equals value, node is not appended as it already exists.
	// - If current node has no following element, append value as node to list.
	// Than advance to Next node.
	for n != nil {
		if n.Data == v {
			// fmt.Printf("node already exists %v\n", v)
			return -1
		}
		if n.Next == nil {
			n.Next = &node{v, nil}
			return -2
		}
		n = n.Next
	}
	return -3
}

// Remove node with specified value from list.
//
// The specified value to be removed has to match the Data property of the node.
func (n *node) RemoveNode(a any) {
	for n.Next != nil {
		if n.Next.Data == a {
			n.Next = n.Next.Next
		}
		n = n.Next
	}
}

// Remove node on specified position.
//
// Position count starts by 1.
func (n *node) RemoveNodeOnPos(pos int) {
	counter := 1
	for n != nil {
		if counter == pos {
			n.Next = n.Next.Next
		}
		n = n.Next
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
			// Save Next node.
			temp := n.Next
			// Append new node to current node.
			n.Next = &node{a, nil}
			// Append former Next node to new node.
			n.Next.Next = temp
		}
		// Advance.
		n = n.Next
		counter += 1
	}
}

// Save the current list to file.
//
// The list is saved to the specified file in binary format.
func (n *node) Save(file string) error {

	// Remove file if already exists.
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}

	// Create file if not exists.
	saveTo, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("cannot create file %v with error %v", file, err)
	}
	defer saveTo.Close()

	// Create new encoder and encode.
	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(&n)
	if err != nil {
		return fmt.Errorf("cannot save to file %v with error %v", file, err)
	}
	return nil
}

// Loads a list from a file.
//
// The file must have been created by the Save() function.
func (n *node) Load(file string) error {
	// Open file.
	loadFrom, err := os.Open(file)

	if err != nil {
		return fmt.Errorf("empty list!, Error: %v", err)
	}
	defer loadFrom.Close()

	// Create new decoder and decode.
	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&n)

	return nil
}
