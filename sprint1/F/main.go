package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := []rune(scanner.Text())
	i := 0
	j := len(s) - 1
	for i < j {
		for !unicode.IsLetter(s[i]) {
			i++
		}
		for !unicode.IsLetter(s[j]) {
			j--
		}
		if unicode.ToLower(s[i]) != unicode.ToLower(s[j]) {
			fmt.Print("False")
			return
		}
		i++
		j--
	}
	fmt.Print("True")
}
