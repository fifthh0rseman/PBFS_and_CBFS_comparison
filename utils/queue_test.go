package utils

import (
	"testing"
)

func TestIntQueue_Empty_NotEmpty(t *testing.T) {
	a := &Node{
		val:  1,
		next: nil,
	}
	q := IntQueue{
		head: a,
		tail: a,
	}
	if q.Empty() {
		t.Errorf("Method q.Empty() says that queue is empty, though shouldn't be")
	}
}

func TestIntQueue_Empty_Empty(t *testing.T) {
	q := IntQueue{
		head: nil,
		tail: nil,
	}
	if !q.Empty() {
		t.Errorf("Method q.Empty() says that queue is not empty, though should be")
	}
}

func TestIntQueue_Pop(t *testing.T) {
	b := &Node{
		val:  1,
		next: nil,
	}
	a := &Node{
		val:  2,
		next: b,
	}
	q := &IntQueue{
		head: a,
		tail: b,
	}
	act := q.Pop()
	if a.val != act {
		t.Errorf("Values mismatch: popped %d instead of %d", act, a.val)
	}
	act = q.Pop()
	if b.val != act {
		t.Errorf("Values mismatch: popped %d instead of %d", act, b.val)
	}
	act = q.Pop()
	if -1 != act {
		t.Errorf("Values mismatch: popped %d instead of %d", act, -1)
	}
}

func TestIntQueue_Push(t *testing.T) {
	q := IntQueue{
		head: nil,
		tail: nil,
	}
	q.Push(1)
	if q.head.val != 1 {
		t.Errorf("Head of queue: %d instead of %d", q.head.val, 1)
	}
	q.Push(2)
	if q.head.next.val != 2 {
		t.Errorf("Next to head: %d instead of %d", q.head.val, 2)
	}
	q.Push(3)
	if q.head.next.next.val != 3 {
		t.Errorf("Next to next to head: %d instead of %d", q.head.val, 3)
	}
}
