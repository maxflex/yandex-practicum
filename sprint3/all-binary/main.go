package main

import (
	"fmt"
)

func allBinary(n int, s string) {
	if n == 0 {
		fmt.Println(s)
		return
	}
	allBinary(n-1, s+"0")
	allBinary(n-1, s+"1")
}

func main() {
	allBinary(3, "")
}
