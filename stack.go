package main

import (
	. "dsa/practice/linkedList"
)

type Stack struct {
	value *DoublyLinkedList
}

func (s *Stack) push(val int) {

	s.value.InsertHead(val)

}

func (s *Stack) pop() {
	s.value.Remove_head()
}

func (s Stack) length() int {
	return s.value.Size
}

func (s Stack) peek() int {

	return s.value.Head_val()

}
