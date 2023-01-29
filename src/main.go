package main

import (
	"errors"
	"fmt"
)

// A single node in a data structure
type node struct {
	data interface{}
	next *node
}

// A singlely linked list
type LinkedList struct {
	head *node
	tail *node
	Size int
}

// Will add data to the end of the singlely linked list
// Args:
//		data (interface{}): the information to store in the list
func (list *LinkedList) Append(data ...interface{}) {
	for i := range data {
		item := &node{data[i], nil}

		if list.head == nil {
			list.head = item
		}
		if list.tail != nil {
			list.tail.next = item
		}

		list.tail = item; list.Size++
	}
}

// Will delete data from the list at the given index
// ARGS:
//		itemIndex (int): The index of the data to delete from the list
func (list *LinkedList) Delete(itemIndex int) {
	if itemIndex < 0 || itemIndex >= list.Size {
		panic("Index out of range for LinkedList")
	}

	currentNode := list.head
	for i := 0; i < itemIndex - 1; i++ {
		currentNode = currentNode.next
	}
	
	currentNode.next = currentNode.next.next
}

// Will remove the first instance of a specified piece of data in the list
// ARGS:
//		data (interface{}): The data to remove
func (list *LinkedList) Remove(data interface{}) error {
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			return nil
		}
	}

	return errors.New("LinkedList: Data not found in instance")
}

// Will return the data stored at the node found at a given index
// ARGS:
//		itemIndex (int): The index of the data to retrieve
func (list *LinkedList) Index(itemIndex int) interface{} {
	if itemIndex < 0 || itemIndex >= list.Size {
		panic("Index out of range for LinkedList")
	}

	currentNode := list.head
	for i := 0; i < itemIndex; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

// Will display the information in the list in the terminal in a simple format
func (list *LinkedList) Display() {
	defer fmt.Println()
	for currentNode := list.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("%v --> ", currentNode.data)
	}
}

// Controls the flow of the program
func main() {
	var myList LinkedList
	myList.Append(1, "Bunny", nil, LinkedList{},)
	myList.Remove(nil)
	myList.Delete(2)
	myList.Display()
}
