package main

import "errors"

type ListNode struct {
	val  string
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size uint
}

func (l *LinkedList) Init(val string) {
	node := ListNode{val, nil}
	l.head = &node
	l.tail = &node
	l.size = 1
}

func (l *LinkedList) Enqueue(val string) {
	if l.IsEmpty() {
		l.Init(val)
	} else {
		node := ListNode{val, nil}
		prevTail := l.tail
		l.tail = &node
		prevTail.next = &node
		l.size++
	}
}

func (l *LinkedList) Dequeue() (string, error) {
	if l.IsEmpty() {
		return "", errors.New("empty list")
	}
	val := l.head.val
	l.head = l.head.next
	l.size--
	return val, nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l LinkedList) String() string {
	var result string
	node := l.head
	for node != nil {
		result += node.val
		node = node.next
	}
	return result
}
