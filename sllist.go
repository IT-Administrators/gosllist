package gosllist

import "fmt"

// List is a container to keep track of head and tail of the list.
type List struct {
	head, tail *node
}

// Listnode. Must contain a reference to parent.
// Contains nil pointer reference when created.
type node struct {
	data any
	next *node
}

type ListWriter interface {
	AddNode(any)
	RemoveNode(any)
}

type ListReader interface {
	Traverse()
	Size() int
}

// Create new node with specified Data and reference to parent.
func NewNode(nodeValue any) *node {
	return &node{
		data: nodeValue,
		next: nil,
	}
}

// Get head of current list.
func (l *List) GetHead() *node {
	return l.head
}

// Get tail of current list.
func (l *List) GetTail() *node {
	return l.tail
}

// Get size of list.
func (l *List) Size() int {
	// Check if list is empty.
	if l.head == nil {
		return 0
	}

	i := 0
	for e := l.head; e != nil; e = e.next {
		i++
	}
	return i
}

// Add node to list.
//
// As a list can contain any value, it is possible to append the same value more than once.
func (l *List) AddNode(v any) {
	if l.tail == nil {
		l.head = &node{data: v}
		l.tail = l.head
	} else {
		l.tail.next = &node{data: v}
		l.tail = l.tail.next
	}
}

// Remove node.
func (l *List) RemoveNode(v any) {
	// Check if head equals value. Than set head to next node.
	// Tail stays as it is.
	if l.head.data == v {
		l.head = l.head.next
	}
	// Save head than traverse list.
	temp := l.head
	// Traverse as long there is a next element.
	for l.head.next != nil {
		if l.head.next.data == v {
			l.head.next = l.head.next.next
		}
		l.head = l.head.next
	}
	// Set head to head before traversal.
	l.head = temp
}

// Print current list to stdout.
func (l *List) Traverse() {
	var elems []any
	for e := l.head; e != nil; e = e.next {
		elems = append(elems, e.data)
	}
	fmt.Println(elems)
}
