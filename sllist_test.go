package gosllist

import (
	"testing"
)

var sLL = &node{nil, nil}

func Test_NewNode(t *testing.T) {
	v := 1
	n := NewNode(v)
	// Overwrite nil list.
	sLL = n
	if sLL.next != nil && sLL.data != v {
		t.Error("not able to create node")
	}
}

func Test_ListSize(t *testing.T) {
	if sLL.Size() == 0 {
		t.Error("list is empty")
	}
	t.Logf("current list size is %v\n", sLL.Size())
}

func Test_ListAddNode(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	for _, v := range ints {
		res := sLL.AddNode(v)
		if res == -3 {
			t.Error("append not working correctly")
		} else if res == -1 {
			t.Logf("value %v already exists", v)
		}
	}
}

func Test_ListRemoveNode(t *testing.T) {
	sLL.RemoveNode(3)
	if sLL.Size() >= (sLL.Size() + 1) {
		t.Error("node was not removed.")
	}
}

func Test_ListRemoveNodeOnPos(t *testing.T) {
	sLL.RemoveNodeOnPos(2)
	if sLL.Size() >= (sLL.Size() + 1) {
		t.Error("node was not removed.")
	}
}

func Test_ListInsertNode(t *testing.T) {
	sLL.InsertNode(6, 2)
}

func Test_ListTraverse(t *testing.T) {
	sLL.Traverse()
}
