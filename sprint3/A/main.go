package main

import "fmt"

func solve(n int, s string, open int, close int) {
	if len(s) == n*2 {
		fmt.Println(s)
		return
	}
	if open < n {
		solve(n, s+"(", open+1, close)
	}
	if close < open {
		solve(n, s+")", open, close+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	solve(n, "", 0, 0)
}
