package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m uint16
	fmt.Scan(&n)
	fmt.Scan(&m)
	matrix := make([][]string, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Fields(line))
	}
	writer := bufio.NewWriter(os.Stdout)
	var i, j uint16
	for j = 0; j < m; j++ {
		for i = 0; i < n; i++ {
			fmt.Fprintf(writer, "%v ", matrix[i][j])
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
