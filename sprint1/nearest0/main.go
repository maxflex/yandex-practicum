package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	blocks := strings.Fields(scanner.Text())
	solve(blocks)
}

func solve(blocks []string) {
	distances := make([]int, len(blocks))
	counter := -1
	for i, v := range blocks {
		if v == "0" {
			counter = 0
		}
		distances[i] = counter
		if counter != -1 {
			counter++
		}
	}

	result := make([]int, len(blocks))
	counter = -1
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] == "0" {
			counter = 0
			continue
		}
		switch counter {
		case -1:
			result[i] = distances[i]
		default:
			counter++
			if distances[i] == -1 || counter < distances[i] {
				result[i] = counter
			} else {
				result[i] = distances[i]
			}
		}
	}
	fmt.Print(result)
}
