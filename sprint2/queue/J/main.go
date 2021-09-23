package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ListNode struct {
	val  string
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size uint
}

func (l *LinkedList) Enqueue(val string) {
	node := ListNode{val, nil}
	if l.IsEmpty() {
		l.head = &node
		l.tail = &node
	} else {
		prevTail := l.tail
		l.tail = &node
		prevTail.next = &node
	}
	l.size++
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

func main() {
	var l LinkedList
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		switch line[0] {
		case "size":
			fmt.Fprintln(writer, l.size)
		case "put":
			l.Enqueue(line[1])
		case "get":
			if val, err := l.Dequeue(); err != nil {
				fmt.Fprintln(writer, "error")
			} else {
				fmt.Fprintln(writer, val)
			}
		}
	}
	writer.Flush()
}
