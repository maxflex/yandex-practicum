package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	m := make(map[string]bool)
	for scanner.Scan() {
		key := scanner.Text()
		if _, ok := m[key]; !ok {
			fmt.Println(key)
		}
		m[key] = true
	}
}
