package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 2*1024*1024)
	scanner.Scan()
	scanner.Scan()
	arr := strings.Fields(scanner.Text())
	var cnt [3]int
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		cnt[i]++
	}
	for i, v := range cnt {
		for j := 0; j < v; j++ {
			fmt.Fprintf(writer, "%d ", i)
		}
	}
	writer.Flush()
}
