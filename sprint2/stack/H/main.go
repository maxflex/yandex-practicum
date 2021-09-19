package main

import (
	"fmt"
)

func main() {
	var s Stack
	var seq string
	opening := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	closing := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	fmt.Scan(&seq)
	for _, i := range seq {
		if _, ok := opening[i]; ok {
			s.push(i)
		} else if _, ok := closing[i]; ok {
			popped, err := s.pop()
			if err != nil || popped != closing[i] {
				fmt.Print("False")
				return
			}
		}
	}
	if s.empty() {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
