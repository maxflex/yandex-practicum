package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n byte, s string) {
	if len(s) == 1 {
		fmt.Println(s)
		return
	}
	digits := map[byte][]string{
		'1': {"a", "b", "c"},
		'2': {"d", "e", "f"},
	}
	for _, char := range digits[n] {
		solve(n, s+char)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := reader.ReadByte()
	fmt.Println(string(b))
	solve(b, "")
}
