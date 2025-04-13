# gosllist

This go library implements a singly linked list.

## Table of contents

1. [Introduction](#introduction)
1. [Getting started](#getting-started)
    1. [Prerequisites](#prerequisites)
    1. [Installation](#installation)
1. [How to use](#how-to-use)
1. [License](/LICENSE)

## Introduction

I created this repository to refresh my knowledge on lists, in this case singly linked lists and see how they can be implemented in golang.

This library might be extended in the future. Currently it uses no concurrency to manipulate the list.

### What are singly linked lists? 

The root element of a singly linked list is a ***node***. A node consists of a value and pointer to the memory address of the next element in the list. On creation of a node the pointer to the next element is a nil pointer. This pointer changes when the node is inserted to a list or a node is appended to the current node.

```Go
// Listnode. Must contain a reference to parent.
// Contains nil pointer reference when created.
type node struct {
	data any
	next *node
}
```

Visually reprensentation of a list:

            head ->      element ->         tail
    [value| ptr] -> [value| ptr] -> [value| nil]

The advantage of a single linked list is that data is stored sequentially with pointer references which makes access to the underlying data very efficient for small lists.

A disadvantage is that a list can only be traversed from head to tail which makes accessing nodes of really large lists slower leading to an O(n) time complexity. The same applies to inserting and deleting elements in huge lists.

# Getting started

### Prerequisites

- Golang installed
- Operatingsystem: Linux or Windows, not tested on mac
- IDE like VS Code, if you want to contribute or change the code

### Installation

The recommended way to use this module is using the go cli.

    go get github.com/IT-Administrators/gosllist

## How to use

After importing the module you can use it the following way.

First of all you need to create an empty list. As the **node** struct is currently unexported
the ```NewNode(any)``` function can be used to create a node. This function is meant to be used to create a new list. All other functions don't need a __*node__ pointer as parameter. The list can be created with **any** starting value. Every new node created by the ```NewNode(any)``` function as a **nil** pointer to the next element as it doesn't exist on creation.

```Go
// Create an empty list. This list will carry all values that will be appended.
sLL := gosllist.NewNode(nil)
```

To append nodes to the list the ```AddNode(any)``` function can be used. It is not necessary to create a node first with the ```NewNode(any)``` function as every value specified in ```AddNode(any)``` is converted to a node internally.

```Go
// Append int value to list.
sLL.AppendNode(1)
```

Every list created, supports the **Traverse()** function which prints the current list to stdout. 

```Go
// Print list to stdout.
sLL.Traverse()
// Output:
// 1 ->
```

### License

[MIT](./LICENSE)