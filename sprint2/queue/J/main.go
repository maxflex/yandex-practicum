package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var l LinkedList
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		fmt.Println(line)
		switch line[0] {
		case "size":
			fmt.Fprintln(writer, l.size)
		case "put":
			l.Enqueue(line[1])
		case "get":
			if val, err := l.Dequeue(); err != nil {
				fmt.Fprintln(writer, "error")
			} else {
				fmt.Fprintln(writer, val)
			}
		}
	}
	writer.Flush()
}
