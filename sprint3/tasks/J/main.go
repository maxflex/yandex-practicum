package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printItems(arr []int) string {
	var result string
	for _, v := range arr {
		result += fmt.Sprint(v) + " "
	}
	return result + "\n"
}

func main() {
	var n int
	fmt.Scan(&n)
	items := make([]int, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	for _, word := range words {
		v, _ := strconv.Atoi(word)
		items = append(items, v)
	}
	wasSorted := true
	swapped := false
	for i := 1; i < n; i++ {
		for j := 0; j < (n - i); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
				swapped = true
				wasSorted = false
			}
		}
		if swapped {
			writer.WriteString(printItems(items))
		} else {
			break
		}
		swapped = false
	}
	if wasSorted {
		writer.WriteString(printItems(items))
	}
	writer.Flush()
}
