package main

func SolutionC(head *ListNode, index int) *ListNode {
	if index == 0 {
		return head.next
	}
	var node, prev *ListNode
	currentIndex := 0
	node = head
	for {
		currentIndex++
		prev = node
		node = node.next
		if currentIndex == index {
			prev.next = node.next
			break
		}
	}
	return head
}
