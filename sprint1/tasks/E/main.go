package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var result string
	fmt.Scan()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > len(result) {
			result = word
		}
	}
	fmt.Println(result)
	fmt.Print(len(result))
}
