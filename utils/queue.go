package utils

import (
	"fmt"
	"strconv"
)

type Node struct {
	val  int
	next *Node
}

type IntQueue struct {
	head *Node
	tail *Node
}

func NewQueue() *IntQueue {
	return &IntQueue{
		head: nil,
		tail: nil,
	}
}

func (q *IntQueue) Pop() int {
	if q.head == nil {
		return -1
	}
	ret := q.head.val
	q.head = q.head.next
	if q.Empty() {
		q.tail = nil
	}
	return ret
}

func (q *IntQueue) Push(i int) {
	n := &Node{
		val:  i,
		next: nil,
	}
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
}

func (q *IntQueue) Empty() bool {
	return q.head == nil
}

func (q *IntQueue) PrintFull() {
	if q.Empty() {
		fmt.Println("Queue empty")
	}
	ret := "Head: " + strconv.Itoa(q.head.val)
	next := q.head.next
	for next != nil {
		ret += ", next: "
		ret += strconv.Itoa(next.val)
		next = next.next
	}
	fmt.Println(ret)
}

func (q *IntQueue) PrintShort() {
	if q.Empty() {
		fmt.Println("Queue empty")
	}
	ret := "Queue: " + strconv.Itoa(q.head.val)
	next := q.head.next
	for next != nil {
		ret += " "
		ret += strconv.Itoa(next.val)
		next = next.next
	}
	fmt.Println(ret)
}
