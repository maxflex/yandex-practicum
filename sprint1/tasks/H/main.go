package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2, result string
	var carry int
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	for len(s2) != len(s1) {
		s2 = "0" + s2
	}
	for i := len(s1) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(s1[i]))
		y, _ := strconv.Atoi(string(s2[i]))
		sum := x + y + carry
		carry = sum / 2
		result += fmt.Sprint(sum % 2)
	}
	if carry > 0 {
		result += "1"
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(string(result[i]))
	}
}
