package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var q, m int
	fmt.Scan(&q, &m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	s := scanner.Text()
	// метод Гонера: ((s1*q + s2)*q + s3) % M
	var h int
	for i := 0; i < len(s); i++ {
		h = (h*q + int(s[i])) % m
	}
	fmt.Print(h)
}
