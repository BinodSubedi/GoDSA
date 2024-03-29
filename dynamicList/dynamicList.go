package dynamiclist

import (
	"fmt"
)

type DynamicList struct {
	Capacity, length int
	val              []int
}

func (l *DynamicList) Initialize() {

	l.val = make([]int, l.Capacity)
	l.length = 0

}

func (l *DynamicList) Add(i int) {

	if l.length == l.Capacity {
		temp := l.val

		l.val = make([]int, l.length*2)

		for i := range len(temp) {

			l.val[i] = temp[i]

		}

	}

	l.val[l.length] = i
	l.length++
}

func (l *DynamicList) Pop() {
	l.length--
}

func (l DynamicList) Debug_iter() {

	for i := 0; i < l.length; i++ {

		fmt.Println(l.val[i])

	}

}
