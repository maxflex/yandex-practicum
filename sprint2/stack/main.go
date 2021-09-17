package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s Stack
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		switch line[0] {
		case "get_max":
			max, err := s.getMax()
			if err == nil {
				fmt.Fprintln(writer, max)
			} else {
				fmt.Fprintln(writer, "None")
			}
		case "push":
			i, _ := strconv.Atoi(line[1])
			s.push(i)
		case "pop":
			_, err := s.pop()
			if err != nil {
				fmt.Fprintln(writer, "error")
			}
		}
	}
	writer.Flush()
}
