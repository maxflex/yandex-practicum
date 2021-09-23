package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	var lines [2][]string

	for i := range lines {
		line, _ := reader.ReadString('\n')
		lines[i] = strings.Fields(strings.TrimSpace(line))
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%s %s ", lines[0][i], lines[1][i])
	}

	writer.Flush()
}
