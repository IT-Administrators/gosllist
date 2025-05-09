package gosllist

import (
	"testing"
)

var sLL = List{}
var safeToFile = "./sllist.gob"

func Test_NewNode(t *testing.T) {
	testcases := []struct {
		name     string
		input    any
		expected *node
	}{
		{
			name:     "Test node creation (int)",
			input:    5,
			expected: &node{5, nil},
		},
		{
			name:     "Test node creation (string)",
			input:    "Test",
			expected: &node{"Test", nil},
		},
		{
			name:     "Test node creation (bool)",
			input:    true,
			expected: &node{true, nil},
		},
	}
	for _, tc := range testcases {
		n := NewNode(tc.input)
		if n.data != tc.expected.data && n.next != tc.expected.next {
			t.Errorf("got %v; expected %v", n, tc.expected)
		}
	}
}

func Test_ListAddNode(t *testing.T) {
	testcases := []struct {
		name     string
		input    any
		expected *node
	}{
		{
			name:     "Test node creation (int)",
			input:    5,
			expected: &node{5, nil},
		},
		{
			name:     "Test node creation (string)",
			input:    "Test",
			expected: &node{"Test", nil},
		},
		{
			name:     "Test node creation (bool)",
			input:    true,
			expected: &node{true, nil},
		},
		{
			name:     "Test node creation (string)",
			input:    "Test2",
			expected: &node{"Test2", nil},
		},
	}
	for _, tc := range testcases {
		sLL.AddNode(tc.input)
	}
	// Check if head matches first element in testcases. The head should not change while a node is appended.
	if testcases[0].expected.data != sLL.GetHead().data {
		t.Errorf("got head %v; expected head %v", sLL.GetHead().data, testcases[0].expected.data)
	}
	// Check if tail is last element of testcases. The tail must always be the last element of the list.
	if testcases[len(testcases)-1].expected.data != sLL.GetTail().data {
		t.Errorf("got tail %v; expected tail %v", sLL.GetTail().data, testcases[0].expected.data)
	}
}

func Test_ListRemoveNode(t *testing.T) {
	currentSize := sLL.Size()
	sLL.RemoveNode("Test")
	if sLL.Size() != currentSize-1 {
		t.Errorf("expected size: %v; got size %v", currentSize, sLL.Size())
	}
}

func Test_ListSize(t *testing.T) {
	t.Logf("current list size is %v.", sLL.Size())
}

func Test_ListTraverse(t *testing.T) {
	sLL.Traverse()
}
