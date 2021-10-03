package main

import (
	"bufio"
	"fmt"
	"os"
)

var digits = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func solve(s string, r string) {
	if len(s) == 0 {
		fmt.Print(r + " ")
		return
	}
	d := digits[s[0]]
	for j := 0; j < len(d); j++ {
		solve(s[1:], r+d[j])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	solve(scanner.Text(), "")
}
