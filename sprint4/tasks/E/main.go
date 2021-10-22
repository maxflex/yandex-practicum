package main

import (
	"fmt"
)

func main() {
	var s string
	var max, localMax int
	fmt.Scan(&s)
	seen := make(map[byte]int)
	for i := 0; i < len(s); {
		char := s[i]
		if index, ok := seen[char]; ok {
			localMax = 0
			i = index + 1
			seen = make(map[byte]int)
		} else {
			localMax++
			if localMax > max {
				max = localMax
			}
			seen[char] = i
			i++
		}
	}
	fmt.Print(max)
}
