package main

func SolutionD(head *ListNode, data string) int {
	node := head
	index := 0
	for {
		if node.data == data {
			return index
		}
		if node.next == nil {
			break
		}
		index++
		node = node.next
	}
	return -1
}
