package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	scanner.Scan()
	arr := strings.Fields(scanner.Text())
	var max int
	for i := 0; i < len(arr); i++ {
		var counter int
		for j := i; j < len(arr); j++ {
			if arr[j] == "0" {
				counter--
			} else {
				counter++
			}
			if counter == 0 {
				streak := j - i + 1
				if streak > max {
					max = streak
				}
			}
		}
	}
	fmt.Print(max)
}
