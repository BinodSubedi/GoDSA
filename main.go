package main

import (
	dl "dsa/practice/dynamicList"
	doubly "dsa/practice/linkedList"
	. "fmt"
)

//import "fmt"

func main() {

	stack()

}

// Stack
func stack() {
	stack := Stack{value: &doubly.DoublyLinkedList{Head: nil, Tail: nil, Size: 0}}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	Println("Head::", stack.peek())

	Println("Length::", stack.length())

	stack.pop()

	Println("Head::", stack.peek())
	Println("Length::", stack.length())
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
