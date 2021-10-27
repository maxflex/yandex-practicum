package main

func Solution(root *Node) int {
	if root == nil {
		return -1
	}
	max := root.value
	leftMax := Solution(root.left)
	rightMax := Solution(root.right)
	if leftMax > max {
		max = leftMax
	}
	if rightMax > max {
		max = rightMax
	}
	return max
}
