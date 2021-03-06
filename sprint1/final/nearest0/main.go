package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(blocks []string) []int {
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
			result[i] = counter
		} else if counter == -1 {
			result[i] = distances[i]
		} else {
			counter++
			if distances[i] == -1 || counter < distances[i] {
				result[i] = counter
			} else {
				result[i] = distances[i]
			}
		}
	}
	return result
}

func printResult(result []int) {
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range result {
		fmt.Fprintf(writer, "%v ", v)
	}
	writer.Flush()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := []byte{}
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	scanner.Scan()
	blocks := strings.Fields(scanner.Text())
	result := solve(blocks)
	printResult(result)
}
