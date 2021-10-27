package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

func getMax(node *Node) int {
	if node == nil {
		return -1
	}
	max := node.value
	leftMax := getMax(node.left)
	rightMax := getMax(node.right)
	if leftMax > max {
		max = leftMax
	}
	if rightMax > max {
		max = rightMax
	}
	return max
}

func main() {
	root := Node{25, nil, nil}
	root.left = &Node{15, nil, nil}
	root.right = &Node{27, nil, nil}
	root.left.left = &Node{30, nil, nil}
	max := getMax(&root)
	fmt.Println(max)
}
