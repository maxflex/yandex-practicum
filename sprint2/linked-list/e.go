package main

func Solution(head *ListNode) *ListNode {
	node := head
	var nextNode *ListNode
	for {
		nextNode = node.next
		node.next, node.prev = node.prev, node.next
		if nextNode == nil {
			return node
		}
		node = nextNode
	}
}
