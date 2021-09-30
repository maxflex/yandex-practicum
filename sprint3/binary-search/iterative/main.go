package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(arr []int, x int) int {
	i := 0
	j := len(arr) - 1
	mid := (i + j) / 2
	for arr[mid] != x {
		if x > arr[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
		if i > j {
			return -1
		}
		mid = (i + j) / 2
	}
	return mid
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
