package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if (abs(a)%2 == 0 && abs(b)%2 == 0 && abs(c)%2 == 0) || (abs(a)%2 == 1 && abs(b)%2 == 1 && abs(c)%2 == 1) {
		fmt.Print("WIN")
	} else {
		fmt.Print("FAIL")
	}
}
