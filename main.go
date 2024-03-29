package main

import (
	dl "dsa/practice/dynamicList"
	doubly "dsa/practice/linkedList"
)

//import "fmt"

func main() {

	LinkedList()
}

// Doubly LinkedList
func LinkedList() {

	l := doubly.DoublyLinkedList{Head: nil, Tail: nil, Size: 0}

	l.InsertHead(10)
	l.InsertHead(50)
	l.InsertHead(69)

	l.Remove_head()
	l.Add_tail(20)
	l.Remove_tail()
	l.Debug_iter()
}

// Dynamic List/Array
func dynamicList() {

	a := dl.DynamicList{Capacity: 5}

	a.Initialize()
	a.Add(10)
	a.Add(20)
	a.Add(30)
	a.Add(40)
	a.Add(50)
	a.Add(60)
	a.Add(70)
	a.Add(80)
	a.Add(90)

	a.Pop()

	a.Debug_iter()

}
