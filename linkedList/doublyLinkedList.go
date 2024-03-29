package linkedlist

import "fmt"

type Node struct {
	prev, next *Node
	val        int
}

type DoublyLinkedList struct {
	Head, Tail *Node
	Size       int
}

/*
func (l *DoublyLinkedList) initiate() {

	l.head, l.tail, l.size = nil, nil, 0

}
*/

func (l *DoublyLinkedList) InsertHead(a int) {

	//create a new node

	new_node := Node{prev: nil, next: nil, val: a}

	if l.Head == nil {

		l.Head, l.Tail = &new_node, &new_node

		l.Size++

		return

	}

	new_node.next = l.Head
	l.Head.prev = &new_node

	l.Head = &new_node

	l.Size++

}

func (l *DoublyLinkedList) Add_tail(i int) {

	new_node := Node{prev: nil, next: nil, val: i}

	if l.Head == nil {

		l.Head, l.Tail = &new_node, &new_node

		l.Size++

		return
	}

	l.Tail.next = &new_node

	new_node.prev = l.Tail

	l.Tail = &new_node

	l.Size++
}

func (l *DoublyLinkedList) Remove_tail() {

	if l.Size == 0 {
		return
	} else if l.Size == 1 {

		l.Head, l.Tail = nil, nil
		l.Size--
		return

	}

	var temp *Node = l.Tail.prev

	temp.next = nil

	l.Tail = nil

	l.Tail = temp

	l.Size--

}

func (l *DoublyLinkedList) Remove_head() {

	if l.Head == nil {
		return
	} else if l.Size == 1 {
		l.Head, l.Tail = nil, nil
		l.Size--
		return
	}

	temp := l.Head.next
	temp.prev = nil

	l.Head = nil

	l.Head = temp
	l.Size--

}

func (l DoublyLinkedList) Head_val() int {

	return l.Head.val

}

func (l DoublyLinkedList) Debug_iter() {

	var now = l.Head
	for range l.Size {

		fmt.Println(now.val)
		now = now.next
	}
}
