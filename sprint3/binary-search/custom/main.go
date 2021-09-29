package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(arr []int, x int) int {
	mid := (len(arr) - 1) / 2
	index := mid
	// fmt.Println("initial index", index, "looking for", x)
	for arr[mid] != x {
		less := false
		if arr[mid] < x {
			arr = arr[mid+1:]
			index++
			// fmt.Println("index++", index)
		} else {
			arr = arr[:mid]
			less = true
			index--
			// fmt.Println("index--", index)
		}
		if len(arr) == 0 {
			return -1
		}
		mid = (len(arr) - 1) / 2
		if less {
			index -= (len(arr) - 1 - mid)
			// fmt.Println("index -= (len(arr) - 1 - mid)", index)
		} else {
			index += mid
			// fmt.Println("index += mid", index)
		}
	}
	return index
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	for _, v := range words {
		x, _ := strconv.Atoi(v)
		arr = append(arr, x)
	}
	scanner.Scan()
	price, _ := strconv.Atoi(scanner.Text())
	fmt.Println(arr, price)
	fmt.Println(search(arr, price))
}
