package main

import "fmt"

func SolutionB(head *ListNode) {
	node := head
	for {
		fmt.Println(node.data)
		if node.next == nil {
			break
		}
		node = node.next
	}
}
