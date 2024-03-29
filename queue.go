package main

import (
	. "dsa/practice/linkedList"
)

type Heap struct {
	value *DoublyLinkedList
}

func (h *Heap) enqueue(val int) {

	h.value.InsertHead(val)

}

func (h *Heap) dequeue() {

	h.value.Remove_tail()

}

func (h Heap) length() int {

	return h.value.Size

}

func (h Heap) next() int {

	return h.value.Tail_val()

}
