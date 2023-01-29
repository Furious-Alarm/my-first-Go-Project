package main

import (
	"fmt"
)

// A single node in a data structure
type Node struct {
	data interface{}
	next *Node
}

// A singlely linked list
type LinkedList struct {
	head *Node
	tail *Node
	Size uint64
}

// [METHOD: LinkedList]
// Will add data to the end of the singlely linked list
// Args:
//
//	data (interface{}): the information to store in the list
func (list *LinkedList) Append(data ...interface{}) {
	for i := range data {
		item := &Node{data[i], nil}

		if list.head == nil {
			list.head = item
		}
		if list.tail != nil {
			list.tail.next = item
		}

		list.tail = item; list.Size++
	}
}

func (list *LinkedList) Index(item_index int) interface{} {
	currentNode := list.head
	for i := 0; i < item_index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (list *LinkedList) Display() {
	defer fmt.Println()
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("%v --> ", currentNode.data)
	}
}

func main() {
	var myList LinkedList
	myList.Append(1, "Bunny", nil,)
	myList.Display()
}
