package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	s1 := scanner.Bytes()
	scanner.Scan()
	s2 := scanner.Bytes()

	var i, j int
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	if i == len(s1) {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
