package main

import "fmt"

func solve(n uint8) uint {
	if n < 2 {
		return 1
	}
	return solve(n-1) + solve(n-2)
}

func main() {
	var n uint8
	fmt.Scan(&n)
	fmt.Print(solve(n))
}
