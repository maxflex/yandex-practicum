package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func printResult(head *ListNode) {
	node := head
	limit := 10
	i := 0
	for {
		fmt.Println(node)
		if node.next == nil || i >= limit {
			break
		}
		node = node.next
		i++
	}
}

func main() {
	head := ListNode{"node0", nil, nil}
	node1 := ListNode{"node1", nil, &head}
	head.next = &node1
	node2 := ListNode{"node2", nil, &node1}
	node1.next = &node2
	node3 := ListNode{"node3", nil, &node2}
	node2.next = &node3

	result := Solution(&head)
	printResult(result)
}
